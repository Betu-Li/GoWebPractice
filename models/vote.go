package models

import (
	"GinExample2/dao"
	"time"
)

type Vote struct {
	ID       int `json:"id"`
	UserID   int `json:"userid"`
	PlayerID int `json:"player_id"`
	AddTime  int `json:"add_time"`
}

func (Vote) TableName() string {
	return "vote"
}

// GetVoteInfo 查询用户是否已经投过票
func GetVoteInfo(userId, playerId int) (vote *Vote, err error) {
	vote = new(Vote)
	err = dao.DB.Debug().Where("user_id=? and player_id=?", userId, playerId).First(vote).Error
	if err != nil {
		return vote, err
	}
	return vote, nil
}

// AddVote 添加投票
func AddVote(userId, playerId int) (int, error) {
	vote := Vote{
		UserID:   userId,
		PlayerID: playerId,
		AddTime:  int(time.Now().Unix()),
	}
	if err := dao.DB.Create(&vote).Error; err != nil {
		return 0, err
	}
	return vote.ID, nil
}
