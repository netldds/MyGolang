package main

import (
	address "MyGolang/proto3"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"time"
)

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
func main() {
	fmt.Println(ByteCountIEC(10240000))
}
func Start() *grpc.ClientConn {
	var opts []grpc.DialOption
	//opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("abc", opts...)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return conn
}
func Post(client address.AddressServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var pps []*address.Person
	pps = append(pps, &address.Person{
		Name:   time.Now().String(),
		Id:     0,
		Email:  "",
		Phones: nil,
		Time:   ptypes.TimestampNow(),
	})
	cs, err := client.PostAddrBook(ctx, &address.AddressBook{
		People: pps,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(*cs)
	return nil
}
