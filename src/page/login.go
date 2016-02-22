package page

import (
	"utils/logger"
	"html/template"
	"net/http"
	"time"
	"dao"
	"fmt"
)

var tag_lg string = "login"

func Login(w http.ResponseWriter, r *http.Request) {
	logger.Println(tag_lg, "login")
	r.ParseForm()                    //解析参数，默认是不会解析的
	logger.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		cookie, err := r.Cookie("login")
		var t *template.Template
		if err == nil && len(cookie.Value) > 0 {
			logger.Println(tag_lg, "cookie:"+cookie.Value)
			t, err = template.ParseFiles("template" + "/forum.html")
		} else {
			t, err = template.ParseFiles("template" + "/login.htm")
		}
		if err != nil {
			logger.Println(tag_lg, err)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseMultipartForm(10)
		name := r.FormValue("username")
		pwd := r.FormValue("password")
		state := dao.CheckUserPwd(name, pwd)
		if state == true {
			//set cookie
			expiration := time.Now()
			expiration = expiration.Add(15 * time.Minute)
			cookie := http.Cookie{
				Name:    "login",
				Value:   name,
				Expires: expiration}
			http.SetCookie(w, &cookie)
			fmt.Fprintf(w, "登陆成功")
		} else {
			fmt.Fprintf(w, "登陆失败")
		}
	}
}
