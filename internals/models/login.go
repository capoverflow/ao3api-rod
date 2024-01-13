package models

import "github.com/go-rod/rod/lib/proto"

type Login struct {
	Scheme      string
	Addr        string
	Login       string
	Username    string
	Password    string
	RememberMe  bool
	CookiesPath string
	Cookies     []proto.NetworkCookieParam
}
