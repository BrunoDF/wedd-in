package models

import "github.com/jinzhu/gorm"

var (
	Confirmations map[string]*Confirmation
)

type Confirmation struct {
	gorm.Model
	Confirmed    bool
	Adults       uint       `gorm:"not null"`
	Children     uint       `gorm:"not null"`
	Invitation   Invitation `gorm:"ForeignKey:InvitationID"`
	InvitationID uint
}

func init() {

}
