package email

import (
	gbaes "ghostbb.io/gb/crypto/gb_aes"
	gbbase64 "ghostbb.io/gb/encoding/gb_base64"
	"ghostbb.io/gb/frame/g"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbconv "ghostbb.io/gb/util/gb_conv"
	gbutil "ghostbb.io/gb/util/gb_util"
	"gopkg.in/gomail.v2"
	"time"
)

func init() {
	var (
		configMap g.Map
		err       error
	)

	if configMap, err = g.Config().Data(gbctx.New()); err != nil {
		panic(err)
	}

	if len(configMap) > 0 {
		if _, v := gbutil.MapPossibleItemByKey(configMap, "Email"); v != nil {
			if err = gbconv.Struct(v, Service); err != nil {
				panic(err)
			}
		}
	}

	// password decode
	if Service.Password != "" {
		var (
			key, _      = gbbase64.Decode([]byte("dd79wy1M9ZCVGTcHMJB2bSlmuScjSaOcnj7hxv02aWc="))
			iv, _       = gbbase64.Decode([]byte("TDp4nh5Bc1+9jloLgU3nMA=="))
			password, _ = gbbase64.Decode([]byte(Service.Password))
		)
		result, err := gbaes.Decrypt(password, key, iv)
		if err != nil {
			panic(err)
		}
		Service.Password = string(result)
	}
}

var (
	Service = new(SmtpSender)
)

// SmtpSender
// A smtp email client
type SmtpSender struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	FromName string `json:"fromName"`
	FromMail string `json:"fromMail"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *SmtpSender) Send(to []string, cc []string, bcc []string, subject string, body string, file ...string) error {
	msg := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	msg.SetHeader("From", msg.FormatAddress(s.FromMail, s.FromName))
	msg.SetHeader("To", to...)
	msg.SetHeader("Cc", cc...)   // Carbon Copy
	msg.SetHeader("Bcc", bcc...) // blind carbon copy
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html;charset=utf-8", body)

	for _, v := range file {
		msg.Attach(v)
	}

	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	return d.DialAndSend(msg)
}

func (s *SmtpSender) SendTo(to []string, subject string, body string, file ...string) error {
	var err error
	// try 3 times, if send error
	for i := 0; i < 3; i++ {
		err = s.Send(to, nil, nil, subject, body, file...)
		if err != nil {
			time.Sleep(time.Millisecond * 500)
			continue
		}
		err = nil
		break
	}
	return err
}
