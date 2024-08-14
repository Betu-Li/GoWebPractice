package controllers

import (
	"GinExample2/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VoteController struct{}

// AddVote 投票
func (vc VoteController) AddVote(c *gin.Context) {
	// 获取用户ID 选手ID
	userIdStr := c.DefaultPostForm("userid", "0")
	playerIdStr := c.DefaultPostForm("playerid", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)

	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}

	player, _ := models.GetPlayerInfo(playerId)
	if player.ID == 0 {
		ReturnError(c, 4001, "选手不存在")
		return
	}

	// 判断用户是否已经投过票
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.ID != 0 {
		ReturnError(c, 4001, "您已经投过票了")
		return
	}

	// 添加投票
	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		// 更新选手票数
		models.UpdatePlayerScore(playerId)
		ReturnSuccess(c, 200, "投票成功", rs, 1)
		return
	}
	ReturnError(c, 4001, "投票失败，请联系管理员")
}
