package page

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"utils"
)

var tag_upload string = "upload"

// 处理/upload 逻辑
func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	mycookie, err := r.Cookie("login")
	if err != nil || mycookie == nil {
		var prompt string = "not login or login timeout"
		fmt.Println(tag_upload, prompt)
		w.Write([]byte(prompt))
		return
	}
	fmt.Println(tag_upload, "mycookie Name:"+mycookie.Name+" value:"+mycookie.Value+" Expires:"+mycookie.Expires.String())
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		Separator := utils.GetSeparator()
		t, err := template.ParseFiles("template" + Separator + "testupload.html")
		if err != nil {
			fmt.Println(tag_upload, err)
			return
		}
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		var f *os.File
		var subpath string = "data"
		f, err = utils.CreatFile(subpath, handler.Filename)
		if err != nil {
			fmt.Println(tag_upload, err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
