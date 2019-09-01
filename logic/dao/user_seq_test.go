package dao

import (
	"fmt"
	"testing"
)

func TestUesrSeqDao_Add(t *testing.T) {
	fmt.Println(UserSeqDao.Add(ctx, 1, 1, 1))
}

func TestUesrSeqDao_Increase(t *testing.T) {
	fmt.Println(UserSeqDao.Increase(ctx, 1, 1))
}

func TestUesrSeqDao_Get(t *testing.T) {
	fmt.Println(UserSeqDao.Get(ctx, 1, 1))
}
