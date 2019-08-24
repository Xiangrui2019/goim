package dao

import (
	"database/sql"
	"goim/logic/model"
	"goim/public/imctx"
	"goim/public/logger"
)

type friendDao struct{}

var FriendDao = new(friendDao)

// Get 获取一个朋友关系
func (*friendDao) Get(ctx *imctx.Context, appId, userId, friendId int64) (*model.Friend, error) {
	var friend = model.Friend{
		AppId:    appId,
		UserId:   userId,
		FriendId: friendId,
	}
	row := ctx.Session.QueryRow(`select id,label,extra,create_time,update_time 
		from friend where app_id = ? and user_id = ? and friend_id = ?`,
		appId, userId, friendId)
	err := row.Scan(&friend.Id, &friend.Label, &friend.Extra, &friend.CreateTime, &friend.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		logger.Sugar.Error(err)
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &friend, nil
}

// Add 插入一条朋友关系
func (*friendDao) Add(ctx *imctx.Context, appId, userId int64, friendId int64, label string, extra string) error {
	_, err := ctx.Session.Exec("insert ignore into friend(app_id,user_id,friend_id,label,extra) values(?,?,?,?,?)",
		appId, userId, friendId, label, extra)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return err
}

// Delete 删除一条朋友关系
func (*friendDao) Delete(ctx *imctx.Context, appId, userId, friendId int64) error {
	_, err := ctx.Session.Exec("delete from friend where app_id = ? and user_id = ? and friend_id = ? ",
		appId, userId, friendId)
	if err != nil {
		logger.Sugar.Error(err)
	}
	return err
}

// ListFriends 获取用户的朋友列表
func (*friendDao) ListUserFriend(ctx *imctx.Context, appId, userId int64) ([]model.UserFriend, error) {
	rows, err := ctx.Session.Query(`select f.label,f.extra,u.user_id,u.nickname,u.sex,u.avatar_url,u.extra,u.create_time,u.update_time 
		from friend f left join user u on f.friend_id = u.user_id where f.app_id = ? and f.user_id = ?`, appId, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	friends := make([]model.UserFriend, 0, 5)
	for rows.Next() {
		var friend model.UserFriend
		err := rows.Scan(&friend.Label, &friend.FriendExtra, &friend.UserId, &friend.Nickname, &friend.Sex, &friend.AvatarUrl,
			&friend.UserExtra, &friend.CreateTime, &friend.UpdateTime)
		if err != nil {
			logger.Sugar.Error(err)
			return nil, err
		}
		friends = append(friends, friend)
	}
	return friends, nil
}
