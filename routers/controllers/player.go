package controllers

import (
	"GinExample2/cache"
	"GinExample2/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type PlayerController struct{}

// GetPlayers 获取选手列表（按id从小到大排序）
func (pc PlayerController) GetPlayers(c *gin.Context) {
	aid, _ := strconv.Atoi(c.DefaultPostForm("aid", "0"))

	playerList, err := models.GetPlayers(aid, " id asc ")
	if err != nil {
		ReturnError(c, 4004, "获取选手信息失败")
		return
	}

	ReturnSuccess(c, 200, "success", playerList, int64(len(playerList)))
}

// GetPlayer 获取选手信息
func (pc PlayerController) GetPlayer(c *gin.Context) {

}

// PlayerSortByScore 根据票数排序选手(排行榜)
func (pc PlayerController) PlayerSortByScore(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	redisKey := "ranking:" + aidStr
	// 从redis中获取排行榜
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		fmt.Println("从redis中到获取排行榜")
		var playerList []*models.Player
		for _, value := range rs {
			// 从redis中获取选手id
			id, _ := strconv.Atoi(value)
			player, _ := models.GetPlayerInfo(id)
			if player.ID > 0 {
				playerList = append(playerList, player)
			}
		}
		ReturnSuccess(c, 200, "success", playerList, 1)
		return
	}

	// 从数据库中获取选手信息
	fmt.Println("从数据库中获取选手信息")
	rsDB, err := models.GetPlayers(aid, " score desc ")
	if err == nil {
		for _, value := range rsDB {
			// 将选手信息存入redis
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.ID, value.Score)).Err()
		}
		// 设置过期时间
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour).Err()
		ReturnSuccess(c, 200, "success", rsDB, 1)
		return
	}
	ReturnError(c, 4004, "获取选手信息失败")
	return
}
