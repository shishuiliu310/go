package common

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	//session对象	login-authentication-key密钥
	store = sessions.NewCookieStore([]byte("login-authentication-key"))
)

func Session(r *http.Request, params string) interface{} {
	//获取session
	session, _ := store.Get(r, "session-user")
	res := session.Values[params]
	return res
}

//url截取
func Substrurl(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start错误")
	}

	if end < 0 || end > length {
		panic("end错误")
	}

	return string(rs[start+1 : end])
}

//字符串截取
func SubString(str string, length int) string {
	res := []rune(str)
	if len(res) <= length {
		return string(res[:len(res)])
	} else {
		return string(res[:length]) + "..."
	}

}

//将byte转化为string
func ByteToString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

//数字加1
func NumAdd(i int) int {
	return i + 1
}
