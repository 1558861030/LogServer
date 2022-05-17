package models

import (
	"gorm.io/gorm"
)

//
//type LogType struct {
//	gorm.Model `json:"-"`
//	//项目名称
//	ProjectName string `gorm:"index"`
//	//应用名称
//	App string `gorm:"index"`
//	//函数名称
//	FuncName string
//	//类型
//	Type string `gorm:"index"`
//	//用户ID
//	UUID string
//	//用户cookie
//	UserCookie string
//	//用户session
//	UserSession string
//	//客户端类型
//	UserOS string
//	//浏览器类型
//	UserBrowserType string
//	//请求响应时间
//	ResponseTime string
//	//请求类型
//	ResponseType string
//	//状态码
//	State string
//	//请求路由
//	ResponseRouter string
//	//错误信息
//	ErrInfo string
//	//用户IP
//	UserIP string
//	//请求主机名称
//	Host string
//	//其他信息
//	OtherInfo []otherInfo `gorm:"foreignkey:LogId"`
//}

type LogType struct {
	gorm.Model `json:"-"`
	//项目名称
	ProjectName string `gorm:"index"`
	//应用名称
	App string `gorm:"index"`
	//函数名称
	FuncName string
	//类型
	Type string `gorm:"index"`
	//请求信息
	RequstInfo requst `gorm:"foreignkey:LogId"`
	//信息
	Info string
	//其他信息
	OtherInfo []otherInfo `gorm:"foreignkey:LogId"`
}
type otherInfo struct {
	gorm.Model `json:"-"`
	//对应log日志的id
	LogId uint `json:"-" gorm:"index"`
	Key   string
	Value string
}

type requst struct {
	gorm.Model `json:"-"`
	LogId      uint `json:"-" gorm:"index"`
	//用户IP
	UserIP string
	//用户cookie
	UserCookie string
	//用户session
	UserSession string
	//客户端类型
	UserOS string
	//浏览器类型
	UserBrowserType string
	//请求响应时间
	ResponseTime string
	//请求主机名称
	Host string
	//请求类型
	ResponseType string
	//状态码
	State string
	//请求路由
	ResponseRouter string
	//返回信息
	ReturnInfo string
	//头部信息
	HeadInfo []HeadInfo `gorm:"foreignkey:RequstID"`
}
type HeadInfo struct {
	gorm.Model `json:"-"`
	//对应log日志的id
	RequstID uint `json:"-" gorm:"index"`
	Key      string
	Value    string
}
