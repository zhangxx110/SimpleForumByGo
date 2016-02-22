package main

import (
	"utils/logger"
	"log"
	"net/http"
	"page"
	"utils/dbutil"
)

var tag_hs string = "http_Server"

type customMux struct {
}
type customWSMux struct {
}

func (p *customMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.ServeHTTP(w, r)
}
func (p *customWSMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page.ServeWebSocket(w, r)
}

/*
创建一个webSocket服务端（:8888）和一个http服务端（:9090）
*/
func main() {
	initApplication()
	releaseApplication()
	logger.Println(tag_hs, "main over")
}

/*
启动http服务器，http://localhost:9090
*/
func startHttp() {
	logger.Println("main", "start http，listen on 9090 port")
	err := http.ListenAndServe(":9090", &customMux{})
	if err != nil {
		log.Println("main", err)
		log.Fatal("server error")
	}
}

/*
启动websocket服务器ws://localhost:8888
*/
func startWs() {
	logger.Println("main", "start websocket port 8888")
	err := http.ListenAndServe(":8888", &customWSMux{})
	if err != nil {
		logger.Println("main", err)
	}
}
func initApplication() {
	startDb()
	go startWs()
	startHttp()
}
func releaseApplication() {
	dbutil.Relase()
}

/**
连接数据库
*/
func startDb() {
	err := dbutil.Init("mysql", "root:root@tcp(127.0.0.1:3306)/forum?charset=utf8&parseTime=true")
	if err != nil {
		log.Println(err)
	}

}
