package models

import "github.com/jinzhu/gorm"

var (
	Invitations map[string]*Invitation
)

type Invitation struct {
	gorm.Model
	Name        string `gorm:"size:255;not null"`
	MaxAdults   uint   `gorm:"not null"`
	MaxChildren uint   `gorm:"not null"`
	Code        string `gorm:"size:4;not null"`
}

func CreateInvitation(invitation Invitation) (InvitationID uint, err bool) {
	if exists := DB.NewRecord(invitation); !exists {
		return 0, true
	}

	DB.Create(&invitation)
	return invitation.ID, false
}

func GetInvitation(InvitationID string) (invitation *Invitation, err error) {
	in := &Invitation{}
	if err := DB.Find(&in, InvitationID).Error; err != nil {
		return nil, err
	}

	return in, nil
}
