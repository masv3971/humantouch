package humantouch

import "testing"

func TestAdd(t *testing.T) {
	c, err := newStoreClient()
	if err != nil {
		t.Error(err)
	}

	testKey := "testKey"

	if err := c.add(testKey); err != nil {
		t.Error(err)
	}
	if !c.exists(testKey) {
		t.Error("key not found!")
	}
}
