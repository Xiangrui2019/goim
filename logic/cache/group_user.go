package cache

import (
	"goim/logic/db"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
	"strconv"
	"time"

	"github.com/json-iterator/go"
)

const (
	groupUserKey  = "group_user:"
	groupUserExp  = 2 * time.Hour
	groupUserSign = "-1" // 群组成员标记，如果存在，表示群组的hash缓存没有过期
)

var ErrKeyExp = "err key expire"

type groupUserCache struct{}

var GroupUserCache = new(groupUserCache)

func (*groupUserCache) Key(appId, groupId int64) string {
	return groupUserKey + strconv.FormatInt(appId, 10) + strconv.FormatInt(groupId, 10)
}

// 保存用户信息
func (c *groupUserCache) MSet(ctx *imctx.Context, appId, groupId int64, userInfos []model.GroupUser) error {
	users := make(map[string]interface{}, len(userInfos)+1)
	for _, userInfo := range userInfos {
		bytes, err := jsoniter.Marshal(userInfo)
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}

		users[strconv.FormatInt(userInfo.UserId, 10)] = bytes
	}
	users["-1"] = []byte{}

	key := c.Key(appId, groupId)
	err := db.RedisCli.HMSet(key, users).Err()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = db.RedisCli.Expire(key, groupUserExp).Err()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (c *groupUserCache) Get(ctx *imctx.Context, appId, groupId, userId int64) (*model.GroupUser, error) {
	values, err := db.RedisCli.HMGet(c.Key(appId, groupId), groupUserSign, strconv.FormatInt(userId, 10)).Result()
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	if values[0] == nil {
		return nil, nil
	}

	if values[1] == nil {
		return nil, nil
	}

	var groupUser = new(model.GroupUser)
	err = jsoniter.Unmarshal(values[2].([]byte), groupUser)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

}
