package main

import (
	"fmt"
	"os"

	"github.com/samuelmunyoki/ProxyDNS/firebase"
	"github.com/samuelmunyoki/ProxyDNS/handlers"
	"github.com/samuelmunyoki/ProxyDNS/utils"

	_ "github.com/joho/godotenv/autoload"
	"github.com/miekg/dns"
)

func main(){

	utils.Clog = utils.NewConsoleLogger()
	utils.Flog = utils.NewFileLogger(os.Getenv("LOGFILE"))

	err := firebase.InitFirestore()
	if err != nil{
		utils.Clog.Error(err)
	}

	utils.Flog.Info("Starting ProxyDNS ... ")
	utils.Clog.Info("Starting ProxyDNS ... ")



	forwarder := handlers.NewForwarder()

	// Create handlers with the embedded Forwarder
	udpHandler := &handlers.HandleUDPReq{Forwarder: forwarder}
	
	tcpHandler := &handlers.HandleTCPReq{Forwarder: forwarder}

	udpserver := &dns.Server{
		Addr:    ":53",
		Net:     "udp",
		Handler: dns.HandlerFunc(udpHandler.ServeDNS),
	}

	tcpserver := &dns.Server{
			Addr:    ":53",
			Net:     "tcp",
			Handler: dns.HandlerFunc(tcpHandler.ServeDNS),
		}

	go func (tcpserver *dns.Server){
		
		fmt.Println("TCP DNS server listening on", tcpserver.Addr)

		err := tcpserver.ListenAndServe()
		if err != nil {
			utils.Clog.Error("Error starting DNS server:", err)
			utils.Flog.Error("Error starting DNS server:", err)
		}

	}(tcpserver)

	fmt.Println("UDP DNS server listening on", udpserver.Addr)

	err = udpserver.ListenAndServe()
	if err != nil {
		utils.Clog.Error("Error starting DNS server:", err)
		utils.Flog.Error("Error starting DNS server:", err)
	}
    
}