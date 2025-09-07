package main

import (
	"fmt"

	"github.com/HoangNguyen-CA/consistent-hashing/internal/hash"
	"github.com/HoangNguyen-CA/consistent-hashing/internal/server"
)

func main() {
	hr := hash.NewHashRing(3)
	s1 := server.NewServer([]byte("s1"))
	s2 := server.NewServer([]byte("s2"))
	s3 := server.NewServer([]byte("s3"))
	hr.AddServer(s1)
	hr.AddServer(s2)
	hr.AddServer(s3)

	hr.PrintAllServers()

	for i := 0; i < 100; i++ {
		fmt.Printf("%s\n", hr.GetServer([]byte(fmt.Sprint(i))).Id)
	}
}
