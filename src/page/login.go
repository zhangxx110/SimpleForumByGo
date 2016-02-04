package page

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"
	"utils"
)

var tag_lg string = "login"

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(tag_lg, "login")
	r.ParseForm()                    //解析参数，默认是不会解析的
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		Separator := utils.GetSeparator()
		t, err := template.ParseFiles("template" + Separator + "login.html")
		if err != nil {
			fmt.Println(tag_lg, err)
			return
		}
		fmt.Println(tag_lg, "token："+token)
		t.Execute(w, token)
	} else if r.Method == "POST" {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("token:", r.Form.Get("token"))
		fmt.Println("username:", r.Form.Get("username"))
		fmt.Println("password:", r.Form.Get("password"))

		//set cookie
		expiration := time.Now()
		expiration = expiration.Add(1 * time.Minute)
		cookie := http.Cookie{
			Name:    "login",
			Value:   r.Form.Get("username") + r.Form.Get("password") + r.Form.Get("token"),
			Expires: expiration}
		http.SetCookie(w, &cookie)
		fmt.Fprintf(w, r.Form.Get("username"))
	}
}
