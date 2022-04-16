package jwt

type Session struct {
	UserName  string `form:"userName" json:"userName" xml:"userName"`
	Ip        string `form:"ip" json:"ip" xml:"ip"`
	LoginType string `form:"loginType" json:"loginType" xml:"loginType"`
}
