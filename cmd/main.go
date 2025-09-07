package main

import (
	"fmt"

	"github.com/HoangNguyen-CA/consistent-hashing/internal/hash"
	"github.com/HoangNguyen-CA/consistent-hashing/internal/server"
)

func main() {
	hr := hash.NewHashRing()
	s1 := server.NewServer([]byte("s1"))
	hr.AddServer(s1)

	fmt.Println(hr.GetServer([]byte("s0")))
}
