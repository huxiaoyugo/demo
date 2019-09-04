package main

import (
	"demo/osutil"
	"demo/osutil/pid"
	"fmt"
	"github.com/mkideal/log"
	"io"
	"net"
	"net/http"
	"time"
)

// 39.107.111.1
var Port = 8811
var pidFile = "tcp.pid"

func main() {

	// 检查 pid
	err := pid.New(pidFile)
	log.If(err != nil).Fatal("pid.New: %v", err)
	defer func() {
		pid.Remove(pidFile)
		log.Info("remove pid file")
		log.Uninit(nil)
	}()

	go listenHttp()
	osutil.ListenQuitAndDump()
	log.Info("exit")
}

func listenHttp() {
	svr := &http.Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", Dispatch)

	svr = &http.Server{
		Addr:           fmt.Sprintf(":%d", Port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go svr.ListenAndServe()
	log.Info("web port: %d", Port)
}

func Dispatch(w http.ResponseWriter, r *http.Request) {
	c, err := w.Write([]byte("hello world"))
	if err != nil {
		log.Error("%v", err)
		return
	}
	log.Info("response write %d byte", c)
}


func listenTCP() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", Port)) //监听TCP，8080端口
	log.If(err != nil).Fatal("net.Listen err: %v", err)
	log.Info("start listen: %v", Port)
	defer l.Close()
	for {
		conn, err := l.Accept() //接受连接，当上面Dial的时候，这就会接受连接
		log.If(err != nil).Fatal("net.Accept err: %v", err)

		fmt.Println("new accept")
		go func(c net.Conn) { //使用goroutine，并发
			//time.Sleep(time.Second*10)
			for {
				buf := make([]byte, 10) //接受字段
				_, err = c.Read(buf)    //读取内容

				if err == io.EOF {
					break
				}
				if err != nil {
					log.Error("err: %v", err)
					break
				}
				log.Info(string(buf))
			}
			c.Close()
		}(conn)
	}
}
