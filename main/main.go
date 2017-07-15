package main

import (
	"common"
	"controller/Ad"
	"controller/Match"
	//"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

// 初始化一个cookie存储对象
// something-very-secret应该是一个你自己的密匙，只要不被别人知道就行
var store = sessions.NewCookieStore([]byte("login-authentication-key"))

func rout(w http.ResponseWriter, r *http.Request) {

	//　获取一个session对象，session-name是session的名字
	session, _ := store.Get(r, "session-user")
	//　设置session生存时间
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	// 在session中存储值
	session.Values["uid"] = "70"
	// 保存更改
	session.Save(r, w)

	//定义URL规则 https://www.idoukou.com:port/controller/action?id=1&name=2
	path := r.RequestURI //获取当前的URL
	//由于是从根目录获取当前请求的URL所以也会请求/favicon.ico。因此，在此屏蔽掉
	if path == "/favicon.ico" {
		return
	}
	Findex := strings.Index(path, "?")          //?第一次出现的位置
	router := common.Substrurl(path, 0, Findex) //当前请求的控制器名
	switch router {

	case "sq/match":
		Match.Index(w, r)
	case "sq/list":
		Match.List(w, r)
	case "sq/praise":
		Match.Praise(w, r)
	}

}

func main() {

	http.HandleFunc("/", rout) //设置访问的路由
	http.Handle("/js/", http.FileServer(http.Dir("../static")))
	http.Handle("/css/", http.FileServer(http.Dir("../static")))
	http.Handle("/images/", http.FileServer(http.Dir("../static")))
	err := http.ListenAndServe(":1800", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
