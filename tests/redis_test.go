package tests

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/xuyu/goredis"
)

func TestCommandStorageRedis(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	c, err := goredis.Dial(&goredis.DialConfig{
		Network: "tcp",
		Address: s.Addr(),
	})
	if err != nil {
		panic(err)
	}

	//Test to SET
	err = c.Set("foo", "bar", 0, 0, false, false)
	if err != nil {
		panic(err)
	}

	if got, err := s.Get("foo"); err != nil || got != "bar" {
		t.Error("`foo` has the wrong value")
	}
	//---------------------------------------------------------------
	// s.CheckGet(t, "foo", "bar")
	// ADD
	a, err := c.ExecuteCommand(
		"HMSET", "ads:1",
		"id", 1,
		"brand", "Subaru",
		"model", "Forester",
		"color", "Blue",
		"price", 30000,
	)
	if err != nil {
		panic(err)
	}

	if ans, err := a.StatusValue(); ans != "OK" || err != nil {
		t.Error("error to add.")
	}
	// Get , GetAll
	reply, err := c.HGetAll("ads:1")
	if err != nil {
		panic(err)
	}

	if reply["id"] != "1" || reply["brand"] != "Subaru" || reply["model"] != "Forester" || reply["color"] != "Blue" || reply["price"] != "30000" {
		t.Error("Error get (wrong data)")
	}

}
