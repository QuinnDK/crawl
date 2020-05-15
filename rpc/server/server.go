package main

import (
	"crawl/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	rpc.Register(rpcdemon.DemoService{})
	listen, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	for {
		connect, err := listen.Accept()

		if err != nil {
			log.Printf("accper error :%v", err)
			continue
		}
		go jsonrpc.ServeConn(connect)
	}
}
