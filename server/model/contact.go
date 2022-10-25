// 联系人
package model

type ContactType int

const (
	ContactTypeFriend ContactType = iota
	ContactTypeCommunity
)

type Contact struct {
	BaseModel
	OwnerId  int64       `json:"ownerId"`
	TargetId int64       `json:"targetId"`
	Type     ContactType `json:"type"`
	Memo     string      `json:"memo"`
}
