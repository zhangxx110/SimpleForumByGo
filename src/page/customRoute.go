package page

import (
	"utils/logger"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"websocket"
)

var tag_customRoute string = "customRoute"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true }, //不检查数据头文件里的Origin字段
}

func ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	logger.Println(tag_customRoute, "ServeWebSocket")
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		logger.Println(tag_customRoute, err)
		return
	}
	doWebSocket(w, r, conn)
}
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                        //解析参数，默认是不会解析的
	logger.Println(tag_customRoute, r.Form) //这些信息是输出到服务器端的打印信息
	logger.Println(tag_customRoute, "path:"+r.URL.Path)
	logger.Println(tag_customRoute, "Scheme:"+r.URL.Scheme)
	logger.Println(tag_customRoute, "Proto:"+r.Proto)
	doHttp(w, r)
}

/*
websocket接收发送消息
*/
func doWebSocket(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
	for {
		var inputData []byte
		msgType, inputData, err := conn.ReadMessage()
		if err != nil {
			logger.Println(tag_customRoute, err)
			return
		}
		logger.Println(tag_customRoute, "receive data:"+string(inputData))
		err = conn.WriteMessage(msgType, inputData)
		if err != nil {
			logger.Println(tag_customRoute, err)
			return
		}
	}
}

/*
http页面路由
*/
func doHttp(w http.ResponseWriter, r *http.Request) {
	var path string = r.URL.Path
	if path == "/" {
		Login(w, r)
		return
	} else if path == "/login" {
		Login(w, r)
	} else if path == "/register" {
		 Register(w, r)
	} else if path == "/upload" {
		Upload(w, r)
	} else if path == "/ajax" {
		OnAjax(w, r)
	} else if strings.HasPrefix(path, "/static/") || strings.HasPrefix(path, "static/") {
		file := "template/static" + r.URL.Path[len("/static"):]
		f, err := os.Open(file)
		defer f.Close()
		if err != nil && os.IsNotExist(err) {
			logger.Println(tag_customRoute, err)
		}
		http.ServeFile(w, r, file)
		return
	} else {
		http.NotFound(w, r)
	}
}

/*
处理http /OnAjax 请求
*/
func OnAjax(w http.ResponseWriter, r *http.Request) {
	logger.Println(tag_customRoute, "OnAjax")
	r.ParseForm()
	//set cookie
	expiration := time.Now()
	expiration = expiration.Add(1 * time.Minute)
	cookie := http.Cookie{
		Name:    "login",
		Value:   r.Form.Get("username") + r.Form.Get("password") + r.Form.Get("token"),
		Expires: expiration}
	http.SetCookie(w, &cookie)
	io.WriteString(w, "恭喜你，登陆成功！")
}
