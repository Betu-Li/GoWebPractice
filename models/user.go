package models

import (
	"GinExample2/dao"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int    `json:"add_time"`
	UpdateTime int    `json:"update_time"`
}

// CreateUser 新增用户
func CreateUser(user *User) (err error) {
	user.AddTime = int(time.Now().Unix())
	user.UpdateTime = int(time.Now().Unix())
	err = dao.DB.Create(user).Error
	return
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (user *User, err error) {
	user = new(User)
	err = dao.DB.Debug().Where("username=?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetUserInfo 根据id获取用户信息
func GetUserInfo(id int) (user *User, err error) {
	user = new(User)
	err = dao.DB.Debug().Where("id=?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return
}
