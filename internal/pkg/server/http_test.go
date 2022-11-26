package server

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/alightzy/groupcache/internal/pkg/cache"
)

func TestServer(t *testing.T) {
	db := map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}

	cache.NewGroup(
		"scores",
		2<<10,
		cache.GetterFunc(
			func(key string) ([]byte, error) {
				log.Println("[SlowDB] search key", key)
				if v, ok := db[key]; ok {
					return []byte(v), nil
				}
				return nil, fmt.Errorf("%s not exist", key)
			},
		),
	)

	// curl http://localhost:9999/_groupcache/scores/Tom
	addr := "localhost:9999"
	peers := NewHTTPPool(addr)
	log.Println("groupcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
