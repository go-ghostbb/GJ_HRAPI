package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type Menu struct {
	gorm.Model
	Type      enum.MenuType `gorm:"size:20;comment:類型" json:"type"`
	Show      enum.MenuShow `gorm:"size:20;not null;comment:顯示於..." json:"show"`
	ParentID  uint          `gorm:"default:0;not null;comment:父節點" json:"parentId"`
	Path      string        `gorm:"not null;comment:路由" json:"path"`
	Name      string        `gorm:"not null;comment:路由名稱" json:"name"`
	Component string        `gorm:"not null;comment:組件路徑" json:"component"`
	Redirect  string        `json:"redirect"`
	Sort      int           `gorm:"default:0;not null;comment:排序標記" json:"sort"`
	Status    bool          `gorm:"default:true;comment:是否啟用" json:"status"`
	Meta      `gorm:"embedded;comment:附加屬性" json:"meta"`

	Children []*Menu `gorm:"-" json:"children"`

	Roles []*Role `gorm:"many2many:role_menu;" json:"roles"`
}

type Meta struct {
	// 路由title  一般必填
	Title string `json:"title"`
	// 動態路由可打開Tab頁數
	DynamicLevel int `json:"dynamicLevel"`
	// 動態路由的實際Path, 即去除路由的動態部分;
	RealPath string `json:"realPath"`
	// 是否忽略KeepAlive緩存
	IgnoreKeepAlive bool `json:"ignoreKeepAlive"`
	// 是否固定標簽
	Affix bool `json:"affix"`
	// 圖標，也是菜單圖標
	Icon string `json:"icon"`
	// 內嵌iframe的地址
	FrameSrc string `json:"frameSrc"`
	// 指定該路由切換的動畫名
	TransitionName string `json:"transitionName"`
	// 隱藏該路由在面包屑上面的顯示
	HideBreadcrumb bool `json:"hideBreadcrumb"`
	// 如果該路由會攜帶參數，且需要在tab頁上面顯示。則需要設置為true
	CarryParam bool `json:"carryParam"`
	// 隱藏所有子菜單
	HideChildrenInMenu bool `json:"hideChildrenInMenu"`
	// 當前激活的菜單。用於配置詳情頁時左側激活的菜單路徑
	CurrentActiveMenu string `json:"currentActiveMenu"`
	// 當前路由不再標簽頁顯示
	HideTab bool `json:"hideTab"`
	// 當前路由不再菜單顯示
	HideMenu bool `json:"hideMenu"`
	// 忽略路由。用於在ROUTE_MAPPING以及BACK權限模式下，生成對應的菜單而忽略路由。2.5.3以上版本有效
	IgnoreRoute bool `json:"ignoreRoute"`
	// 是否在子級菜單的完整path中忽略本級path。2.5.3以上版本有效
	HidePathForChildren bool `json:"hidePathForChildren"`
}
