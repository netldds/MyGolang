package main

import (
	address "MyGolang/proto3"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
	"time"
)

type AddressBookS struct {
}

func (a *AddressBookS) PostAddrBook(ctx context.Context, in *address.AddressBook) (*address.CoreResp, error) {
	fmt.Println(in.People)
	var addr address.CoreResp
	addr.Ok = true
	addr.Err = fmt.Sprintf("server timestamp is :%v", time.Now())
	fmt.Println("client timestamp is :%v", in.People[0].Time)
	return &addr, nil
}
func (a *AddressBookS) PostAddrBookStreamS(addr *address.AddressBook, stream address.AddressService_PostAddrBookStreamSServer) error {
	stream.Send(&address.CoreResp{
		Err: "server stream",
		Ok:  true,
	})
	return nil
}
func (a *AddressBookS) PostAddrBookStreamC(stream address.AddressService_PostAddrBookStreamCServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			break
		}
		if len(in.People) > 0 {
			name := in.People[0].Name
			fmt.Println(name)
		}
	}
	ctx := stream.Context()
	fmt.Println(ctx)
	stream.SendAndClose(&address.CoreResp{
		Err: "client...stream",
		Ok:  true,
	})
	return nil
}
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 2080))
	if err != nil {
		fmt.Println(err)
	}
	grpcServer := grpc.NewServer()
	address.RegisterAddressServiceServer(grpcServer, &AddressBookS{})
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}
}
