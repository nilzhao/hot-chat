// 群聊组
package model

type Community struct {
	BaseModel
	Name    string `json:"name"`
	OwnerId string `json:"ownerId"`
	Avatar  string `json:"avatar"`
}
