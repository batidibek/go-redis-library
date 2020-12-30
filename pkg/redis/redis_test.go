package redis

import (
	"testing"
	"time"
)

const (
	testKey   = "foo"
	testValue = "bar"
)

func TestRedis(t *testing.T) {

	c := NewRedisAdapter("localhost:6379", 0)

	if err := c.Set(testKey, testValue, 10*time.Second); err != nil {
		t.Errorf("set fail: expected nil, got %v", err)
	} else {
		t.Log("set success")
	}

	if res, _ := c.Get(testKey); res != testValue {
		t.Errorf("get fail, wrong value: expected %s, got %s", testValue, res)
	} else {
		t.Log("get success")
	}

	if _, err := c.Get("bar"); err == nil {
		t.Errorf("get fail: expected an error, got %v", err)
	} else {
		t.Log("get success (invalid key)")
	}


	if res, _ := c.Contains(testKey); res != 1 {
		t.Errorf("contains failed: the key %s should be exist", testKey)
	} else {
		t.Log("contains success: key exists")
	}


	if err := c.HSet("bar", "field1", "value1", "field2", "value2"); err != nil {
		t.Errorf("Hset fail: expected nil, got %v", err)
	} else {
		t.Log("Hset success")
	}

	if values, _ := c.HGetAll("bar"); values["field1"] == "" || values["field2"] == ""  {
		t.Errorf("HGetAll failed, values: %s", values)
	} else {
		t.Log("HGetAll success")
	}


	if err := c.Delete(testKey); err != nil {
		t.Errorf("delete failed: expected nil, got %v", err)
	} else {
		t.Log("delete success")
	}


	if res, _ := c.Contains(testKey); res == 1  {
		t.Errorf("contains failed: the key %s should not be exist", testKey)
	} else {
		t.Log("contains success: key does not exist exist")
	}


}
