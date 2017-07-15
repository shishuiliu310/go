package Match

import (
	"common"
	"fmt"
	"html/template"
	"model"
	//"io/ioutil"
	"encoding/json"
	"net/http"
)

type HtmlContent struct {
	User       []map[string]string
	FirstUser  []map[string]string
	IsLogin    string
	CountUsers int
}

var (
	htmlcontent HtmlContent
)

func Index(w http.ResponseWriter, r *http.Request) {
	//获取登录用户的uid
	uid := common.Session(r, "uid")
	if uid != nil && uid != "" {
		htmlcontent.IsLogin = "1"
	} else {
		htmlcontent.IsLogin = "2"
	}
	//第一种注册函数并加载模板	可注册多个
	//t := template.New("index.html")                             //创建模板
	//t.Funcs(template.FuncMap{"substring": common.SubString})    //向模板中注入函数
	//bytes, _ := ioutil.ReadFile("../template/match/index.html") //读文件
	//template.Must(t.Parse(string(bytes)))                       //将字符串读作模板
	//第二种注册函数并加载模板	得给模板起个名字才行
	//	t := template.New("index.html").Funcs(template.FuncMap{"substring": common.SubString})
	//	t, _ = t.ParseFiles("../template/match/index.html")
	//第三种注册函数并加载模板	得给模板起个名字才行	可注册多个
	tempfunc := make(template.FuncMap)
	tempfunc["substring"] = common.SubString
	tempfunc["numadd"] = common.NumAdd
	t := template.New("index.html")
	t = t.Funcs(tempfunc)
	t, _ = t.ParseFiles("../template/match/index.html")
	// 将一个文件读作模板	无注册函数情况下使用
	//t, _ := template.ParseFiles("../template/match/index.html")
	firstSql := `你的sql语句`
	countSql := `你的sql语句`
	firstUser := model.DataToSlice(firstSql)
	countUser := model.Count(countSql)
	htmlcontent.FirstUser = firstUser
	htmlcontent.CountUsers = countUser
	t.Execute(w, htmlcontent)
}
func List(w http.ResponseWriter, r *http.Request) {
	tempfunc := make(template.FuncMap)
	tempfunc["substring"] = common.SubString
	tempfunc["numadd"] = common.NumAdd
	t := template.New("list.html")
	t = t.Funcs(tempfunc)
	t, _ = t.ParseFiles("../template/match/list.html")
	//获取page
	page := "0"
	p := r.PostFormValue("page")
	if p != "" {
		page = p
	}
	sql := `你的sql语句` + page + ",1"
	user := model.DataToSlice(sql)
	htmlcontent.User = user
	t.Execute(w, htmlcontent)
}
func Praise(w http.ResponseWriter, r *http.Request) {
	//获取歌曲id
	res := make(map[string]bool)
	musicId := r.PostFormValue("musicId")
	sql := "你的sql语句" + musicId
	row := model.AddPraiseNum(sql)
	if row {
		res["state"] = true
	} else {
		res["state"] = false
	}
	lang, _ := json.Marshal(res)
	str := string(lang)
	fmt.Fprintf(w, string(str))

}
