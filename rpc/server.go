package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//RPC for msisdn.go

func main(){

	msisdn := new(MSISDN)
	err := rpc.Register(msisdn)
	if err != nil {
		log.Fatalf("Format of service msisdn isn't correct. %s", err)
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}
	
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}