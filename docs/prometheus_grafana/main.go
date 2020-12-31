package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/mem"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	//初始化日志服务
	logger := log.New(os.Stdout, "[Memory]", log.Lshortfile|log.Ldate|log.Ltime)

	//初始一个http handler
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/hello", Hello)

	//初始化一个容器
	diskPercent := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "memeory_percent",
		Help: "memeory use percent",
	},
		[]string{"percent"},
	)
	prometheus.MustRegister(diskPercent)

	// 启动web服务，监听11024端口
	go func() {
		logger.Println("ListenAndServe at:11024")
		err := http.ListenAndServe(":11024", nil)
		if err != nil {
			logger.Fatal("ListenAndServe: ", err)
		}
	}()

	//收集内存使用的百分比
	for {
		logger.Println("start collect memory used percent!")
		v, err := mem.VirtualMemory()
		if err != nil {
			logger.Printf("get memeory use percent error:%s", err)
		}
		usedPercent := v.UsedPercent
		logger.Println("get memeory use percent:", usedPercent)
		diskPercent.WithLabelValues("usedMemory").Set(usedPercent)
		time.Sleep(time.Second * 2)
	}
}
