package dao

import (
	"fmt"
	"testing"
)

func TestFriendDao_Get(t *testing.T) {
	fmt.Println(FriendDao.Get(ctx, 1, 1, 2))
}

func TestFriendDao_Add(t *testing.T) {
	fmt.Println(FriendDao.Add(ctx, 1, 1, 1, "lable", "extra"))
}

func TestFriendDao_Delete(t *testing.T) {
	fmt.Println(FriendDao.Delete(ctx, 1, 1, 1))
}

func TestFriendDao_ListUserFriend(t *testing.T) {
	fmt.Println(FriendDao.ListUserFriend(ctx, 1, 1))
}
