package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/sunrisedo/conf"
	"github.com/sunrisedo/daemon/controllers"
)

var (
	cfg = conf.NewConfig("init.conf")
)

func init() {
	// 初始化配置文件
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("init data start...")

	log.Println("init data finish.")
}

func main() {
	log.Println("Daemon start.")
	switch len(os.Args) {
	case 1:
		//注册路由
		for addr, controller := range RouteMap {
			http.HandleFunc(addr, controller)
		}
		// http.Handle("/webroot/", http.FileServer(http.Dir("webroot")))
		if err := http.ListenAndServe(cfg.Read("server", "port"), nil); err != nil {
			log.Println("Daemon server error:", err)
		}
	default:
		if err := controllers.NewClient(cfg.Read("server", "port")).Listen(); err != nil {
			log.Println("Daemon execute error:", err)
		}
	}

	// // Set up channel on which to send signal notifications.
	// // We must use a buffered channel or risk missing the signal
	// // if we're not ready to receive when the signal is sent.
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	// // Waiting for interrupt by system signal
	// killSignal := <-interrupt
	// log.Println("Daemon exit:", killSignal)
}
