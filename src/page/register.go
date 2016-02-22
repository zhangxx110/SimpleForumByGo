package page

import (
	"dao"
	"fmt"
	"html/template"
	"net/http"
	"utils/logger"
	"time"
)

var tag_regis = "register"

func Register(w http.ResponseWriter, r *http.Request) {
	logger.Println(tag_regis, "Register")
	err := r.ParseForm()
	if err != nil {
		logger.Println(tag_regis, err)
	} //解析参数，默认是不会解析的
	logger.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		cookie, err := r.Cookie("login")
		var t *template.Template
		if err == nil && len(cookie.Value) > 0 {
			logger.Println(tag_lg, "cookie:"+cookie.Value)
			t, err = template.ParseFiles("template" + "/forum.html")
		} else {
			t, err = template.ParseFiles("template" + "/register.htm")
		}
		if err != nil {
			logger.Println(tag_lg, err)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseMultipartForm(10)
		logger.Println("form:", r.Form)
		logger.Println("username:", r.FormValue("username"))
		logger.Println("password:", r.FormValue("password"))
		name := r.FormValue("username")
		pwd := r.FormValue("password")
		isExist := dao.CheckUserName(name)
		if !isExist {
			state := dao.AddUser(name, pwd)
			if state == false {
				logger.Println(tag_regis, "addUser failed")
				fmt.Fprintf(w, "注册失败，请重新注册")
			} else {
				expiration := time.Now()
				expiration = expiration.Add(10 * time.Minute)
				//		stamp := utils.EncMd5(name + password)
				cookie := http.Cookie{
					Name:    "login",
					Value:   name,
					Expires: expiration}
				http.SetCookie(w, &cookie)
				fmt.Fprintf(w, "hello "+r.Form.Get("username"))
			}
		} else {
			fmt.Fprintf(w, name+" has existed")
		}
	}
}
