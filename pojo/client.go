package pojo

import (
	"errors"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	RoomId      string             `json:"roomId"`
	UserId      string             `json:"userId"`
	UserName    string             `json:"userName"`
	UserProfile string             `json:"userProfile"`
	UserCer     bool               `json:"userCer"`
	UserConn    *websocket.Conn    `json:"-"`
	Hub         map[string]*Client `json:"-"`
	Score       int                `json:"score"`
}

// Certificate 客服端认证
func (c *Client) Certificate(certificationInfo *CertificationInfo) (bool, error) {
	if certificationInfo.UserId != "" && certificationInfo.UserName != "" && certificationInfo.UserProfile != "" {
		c.UserId = certificationInfo.UserId
		c.UserName = certificationInfo.UserName
		c.UserProfile = certificationInfo.UserProfile
		c.UserCer = true
		return true, nil
	} else {
		c.UserCer = false
		err := errors.New("用户认证失败")
		return false, err
	}
}

// CreateRoom 创建房间 -- 用户连接存在并且用户认证通过
func (c *Client) CreateRoom() {
	if c.UserConn != nil && c.UserCer {
		uid := uuid.NewV4()
		c.RoomId = uid.String()
	}
}

// JoinHub 将连接加入中心
func (c *Client) JoinHub() {

}

// DeleteFromHub 将连接从中心里删除
func (c *Client) DeleteFromHub() {

}