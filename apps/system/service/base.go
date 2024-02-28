package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"gorm.io/gorm"
	"hrapi/apps/system/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/jwtx"
	"hrapi/internal/utils/jwtx/claims"
	"hrapi/internal/utils/password"
	"hrapi/internal/utils/token"
)

func Base(ctx context.Context) IBase {
	return &base{ctx: ctx}
}

type (
	IBase interface {
		Migrate() error
		Login(in model.LoginReq) (out model.LoginRes, err error)
		Logout(employeeID uint, accessToken string) error
	}

	base struct {
		ctx context.Context
	}
)

var (
	ErrUserFrozen        = gberror.New("this user has been frozen")
	ErrPasswordIncorrect = gberror.New("password incorrect")
	ErrUserNotFound      = gberror.New("user not found")
)

func (b *base) Migrate() error {
	g.Log().Info(b.ctx, "db auto migrate")
	if err := g.DB().AutoMigrate(types.Basic()...); err != nil {
		g.Log().Error(b.ctx, "db auto migrate error:", err.Error())
		return err
	}
	return nil
}

func (b *base) Login(in model.LoginReq) (out model.LoginRes, err error) {
	var (
		qLoginInfo = query.LoginInformation
		loginInfo  *types.LoginInformation
	)

	loginInfo, err = qLoginInfo.WithContext(dbcache.WithCtx(b.ctx)).
		Where(qLoginInfo.Account.Value(in.Account)).
		Preload(qLoginInfo.Employee, qLoginInfo.Employee.Roles).First()
	if err != nil {
		if gberror.Is(err, gorm.ErrRecordNotFound) {
			return model.LoginRes{}, ErrUserNotFound
		}
		return
	}
	// 檢查帳號是否啟用
	if !loginInfo.Status {
		return model.LoginRes{}, ErrUserFrozen
	}
	// 檢查密碼正確性
	if !password.Check(loginInfo.Password, in.Password) {
		return model.LoginRes{}, ErrPasswordIncorrect
	}

	// Create token
	roles := make([]string, 0)
	for _, role := range loginInfo.Employee.Roles {
		roles = append(roles, role.Code)
	}
	accessToken, refreshToken, err := jwtx.Service.CreateToken(claims.BaseClaims{
		EmployeeID: loginInfo.Employee.ID,
		RealName:   loginInfo.Employee.RealName,
		Account:    loginInfo.Account,
		Roles:      roles,
	})
	if err != nil {
		return model.LoginRes{}, err
	}

	// Put the refreshToken into redis
	if err = token.SetRefreshToken(g.Redis(), gbconv.String(loginInfo.EmployeeID), refreshToken, jwtx.Service.GetRefreshTokenEp()); err != nil {
		return model.LoginRes{}, err
	}

	return model.LoginRes{Token: accessToken}, nil
}

func (b *base) Logout(employeeID uint, accessToken string) error {
	// Delete the refreshToken from redis
	if err := token.DelRefreshToken(g.Redis(), gbconv.String(employeeID)); err != nil {
		return err
	}
	// Blacklist the accessToken to avoid using the same of the accessToken for verification
	if err := token.SetBlack(g.Redis(), accessToken); err != nil {
		return err
	}
	return nil
}
