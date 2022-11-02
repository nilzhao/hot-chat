// 联系人
package model

type ContactType int

const (
	CONTACT_TYPE_FRIEND ContactType = iota
	CONTACT_TYPE_COMMUNITY
)

type Contact struct {
	BaseModel
	OwnerId    int64       `json:"ownerId"`
	TargetId   int64       `json:"targetId"`
	Type       ContactType `json:"type"`
	Memo       string      `json:"memo"`
	TargetUser User        `json:"targetUser,omitempty" gorm:"foreignKey:TargetId"`
}
