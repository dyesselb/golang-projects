package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"tracker/internal"
	"tracker/internal/config"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	handler := internal.NewHandler()
	log.Println("register routes")
	startServer(handler.Mux)
}

func startServer(router *http.ServeMux){
	log.Println("Start the application...")
	conf, err := config.LoadConfig("config.json")
	if err != nil{
		log.Fatal(err)
	}
	listner, err := net.Listen(conf.Listen.Protocol, fmt.Sprintf("%s:%s", conf.Listen.BindIp, conf.Listen.Port))
	if err != nil{
		log.Fatal(err)
	}
	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Printf("Server is listening port %s:%s\n", conf.Listen.BindIp, conf.Listen.Port)
	log.Fatal(server.Serve(listner))
}
