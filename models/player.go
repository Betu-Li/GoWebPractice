package models

import (
	"GinExample2/dao"
	"github.com/jinzhu/gorm"
)

type Player struct {
	ID          int    `json:"id"`
	Aid         int    `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nickname"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int    `json:"score"`
	AddTime     int    `json:"add_time"`
	UpdateTime  int    `json:"update_time"`
}

// GetPlayerInfo 根据id获取参赛者
func GetPlayerInfo(id int) (player *Player, err error) {
	player = new(Player)
	err = dao.DB.Debug().Where("id=?", id).First(player).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetPlayers 获取所有参赛者(带排序)
func GetPlayers(aid int, sort string) (playerList []*Player, err error) {
	err = dao.DB.Debug().Where("aid=?", aid).Order(sort).Find(&playerList).Error
	if err != nil {
		return nil, err
	}
	return
}

// UpdatePlayerScore 更新参赛者分数
func UpdatePlayerScore(id int) {
	var player Player
	err := dao.DB.Debug().Where("id=?", id).First(&player).UpdateColumn("score", gorm.Expr("score + ?", 1)).Error
	if err != nil {
		return
	}
}
