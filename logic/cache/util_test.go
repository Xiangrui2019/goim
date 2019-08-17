package cache

import (
	"fmt"
	"testing"
)

type A struct {
	A int
}

func TestHset(t *testing.T) {
	fmt.Println(hset("1", "1", A{A: 1}))
}

func TestHget(t *testing.T) {
	var a A
	fmt.Println(hget("1", "1", &a))
	fmt.Println(a)
}

func TestHdel(t *testing.T) {
	fmt.Println(hdel("1", "1"))
}
