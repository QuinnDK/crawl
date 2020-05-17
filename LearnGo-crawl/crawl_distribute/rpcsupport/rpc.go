package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {

	rpc.Register(service)
	listen, err := net.Listen("tcp", host)

	if err != nil {
		return err
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

func NewClient(host string) (*rpc.Client, error) {

	connect, err := net.Dial("tcp", host)

	if err != nil {
		panic(err)
	}
	return jsonrpc.NewClient(connect), nil
}
