package humantouch

import (
	"errors"
	"time"

	cache "github.com/patrickmn/go-cache"
)

var (
	// ErrKeyCollide error for a key that collide
	ErrKeyCollide = errors.New("ERR_KEY_COLLIDE")
	// ErrNoKey error for key that does not exists
	ErrNoKey = errors.New("ERR_NO_KEY")
)

type storeClient struct {
	store *cache.Cache
}

func newStoreClient() (*storeClient, error) {
	c := &storeClient{
		store: cache.New(cache.NoExpiration, 0*time.Second),
	}

	return c, nil
}

func (c *storeClient) add(key string) error {
	_, found := c.store.Get(key)
	if found {
		return ErrKeyCollide
	}
	if err := c.store.Add(key, "", cache.DefaultExpiration); err != nil {
		return err
	}
	return nil
}

func (c *storeClient) del(key string) {
	c.store.Delete(key)
}

func (c *storeClient) exists(key string) bool {
	_, found := c.store.Get(key)
	return found
}
