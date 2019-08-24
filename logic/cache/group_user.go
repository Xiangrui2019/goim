package cache

import (
	"errors"
	"goim/logic/db"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
	"strconv"
	"time"

	"github.com/json-iterator/go"
)

const (
	GroupUserKey = "group_user:"
	GroupUserExp = 2 * time.Hour
)

var ErrResult = errors.New("error redis result")

type groupUserCache struct{}

var GroupUserCache = new(groupUserCache)

func (*groupUserCache) Key(appId, groupId int64) string {
	return GroupUserKey + strconv.FormatInt(appId, 10) + ":" + strconv.FormatInt(groupId, 10)
}

// SetAll 保存群组所有用户的信息
func (c *groupUserCache) SetAll(ctx *imctx.Context, appId, groupId int64, userInfos []model.GroupUser) error {
	users := make(map[string]interface{}, len(userInfos)+1)
	for _, userInfo := range userInfos {
		bytes, err := jsoniter.Marshal(userInfo)
		if err != nil {
			logger.Sugar.Error(err)
			return err
		}

		users[strconv.FormatInt(userInfo.UserId, 10)] = bytes
	}

	key := c.Key(appId, groupId)
	err := db.RedisCli.HMSet(key, users).Err()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = db.RedisCli.Expire(key, GroupUserExp).Err()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

// GetAll 获取群组的所有用户，如果缓存里面没有，len(GroupUser == 0)
func (c *groupUserCache) GetAll(ctx *imctx.Context, appId, groupId int64) ([]model.GroupUser, error) {
	result, err := getHashAll(c.Key(appId, groupId))
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	users := make([]model.GroupUser, 0, len(result))

	for _, v := range result {
		var user model.GroupUser
		err = jsoniter.Unmarshal(v, &user)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Get 获取用户在群组中的信息
func (c *groupUserCache) Get(ctx *imctx.Context, appId, groupId, userId int64) (*model.GroupUser, error) {
	var groupUser model.GroupUser
	err := getHash(c.Key(appId, groupId), strconv.FormatInt(userId, 10), &groupUser)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	return &groupUser, nil
}

// Set 设置用户在群组中的信息
func (c *groupUserCache) Set(ctx *imctx.Context, appId, groupId int64, user model.GroupUser) error {
	err := setHash(c.Key(appId, groupId), strconv.FormatInt(user.UserId, 10), &user)
	if err != nil && err != ErrKeyNotExists {
		logger.Sugar.Error(err)
		return err
	}

	return nil
}

func (c *groupUserCache) Del(ctx *imctx.Context, appId, groupId, userId int64) error {
	err := delHash(c.Key(appId, groupId), strconv.FormatInt(userId, 10))
	if err != nil && err != ErrKeyNotExists {
		logger.Sugar.Error(err)
		return err
	}

	return nil
}
