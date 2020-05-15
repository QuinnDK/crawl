package main

import (
	rpcdemon "crawl/rpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	connect, err := net.Dial("tcp", ":1234")

	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(connect)

	var result float64

	err = client.Call("DemoService.Div", rpcdemon.Args{10, 3}, &result)

	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", result)

}
