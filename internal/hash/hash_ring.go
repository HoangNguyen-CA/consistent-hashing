package hash

import (
	"fmt"
	"math/big"

	"github.com/HoangNguyen-CA/consistent-hashing/internal/server"
	"github.com/igrmk/treemap/v2"
)

type HashRing struct {
	servers  *treemap.TreeMap[*big.Int, *server.Server]
	replicas uint
}

func (hr *HashRing) AddServer(server *server.Server) {
	h := Hash(server.Id)
	hr.servers.Set(h, server)
}

func (hr *HashRing) GetServer(id []byte) *server.Server {
	h := Hash(id)

	it := hr.servers.Iterator()
	for it.Valid() {
		fmt.Println(h, it.Key())
		if h.Cmp(it.Key()) <= 0 {
			return it.Value()
		}
		it.Next()
	}

	return hr.servers.Iterator().Value()
}

func (hr *HashRing) PrintAllServers() {
	it := hr.servers.Iterator()
	for it.Valid() {
		fmt.Printf("Server %s stored at: %v\n", it.Value().Id, it.Key())
		it.Next()
	}
}

func NewHashRing() *HashRing {
	tr := treemap.NewWithKeyCompare[*big.Int, *server.Server](func(a, b *big.Int) bool {
		return a.Cmp(b) < 0
	})
	return &HashRing{servers: tr}
}
