package cache

import (
	"goim/logic/db"
	"goim/public/imctx"
	"testing"
)

var ctx = imctx.NewContext(db.Factoty.GetSession())

func TestGroupUserCache_Get(t *testing.T) {
	GroupUserCache.Get(ctx, 0, 0, 0)
}
