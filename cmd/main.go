package main

import (
	"fmt"
	"os"

	"githu.com/samuelmunyoki/ProxyDNS/handlers"
	"githu.com/samuelmunyoki/ProxyDNS/utils"

	_ "github.com/joho/godotenv/autoload"
	"github.com/miekg/dns"
)

func main(){

	Clog := utils.NewConsoleLogger()
	Flog := utils.NewFileLogger(os.Getenv("LOGFILE"))
	Flog.Info("Starting ProxyDNS ... ")
	Clog.Info("Starting ProxyDNS ... ")

	forwarder := &handlers.Forwarder{ForwardAddr: "8.8.8.8:53"}

	// Create handlers with the embedded Forwarder
	udpHandler := &handlers.HandleUDPReq{Forwarder: forwarder}
	tcpHandler := &handlers.HandleTCPReq{Forwarder: forwarder}

	
	
	go func (){
		tcpserver := &dns.Server{

			Addr:    ":53",
			Net:     "tcp",
			Handler: dns.HandlerFunc(tcpHandler.ServeDNS),
		}
		

		fmt.Println("TCP DNS server listening on", tcpserver.Addr)
		err := tcpserver.ListenAndServe()
		if err != nil {
			Clog.Error("Error starting DNS server:", err)
			Flog.Error("Error starting DNS server:", err)
		}

	}()

	udpserver := &dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: dns.HandlerFunc(udpHandler.ServeDNS),
	}
	

	fmt.Println("UDP DNS server listening on", udpserver.Addr)
	err := udpserver.ListenAndServe()
	if err != nil {
		Clog.Error("Error starting DNS server:", err)
		Flog.Error("Error starting DNS server:", err)
	}
	
}