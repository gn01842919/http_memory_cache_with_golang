package cache

import "log"

// Cache interface
type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

// New for Cache
func New(typ string) Cache {
	var c Cache // TODO: Why not pointer?? Does 'return' acts like std::move in C++?
	if typ == "inmemory" {
		c = newInMemoryCache()
	}
	if c == nil {
		panic("unknown cache tyype" + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
