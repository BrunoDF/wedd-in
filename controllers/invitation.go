package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"
	"wedd-in/models"

	"github.com/astaxie/beego"
)

// Operations about invitation
type InvitationController struct {
	beego.Controller
}

// @Title Get
// @Description find invitation by invitationID
// @Param	invitationID		path 	string	true		"the invitationID you want to get"
// @Success 200 {invitation} models.Invitation
// @Failure 404 html
// @router /:invitationID [get]
func (i *InvitationController) Get() {
	invitationID := i.Ctx.Input.Param(":invitationID")
	if invitationID != "" {
		invitation, err := models.GetInvitation(invitationID)
		if err != nil {
			if invitation == nil {
				i.Abort("404")
			}
		} else {
			i.Data["json"] = invitation
		}
	}
	i.ServeJSON()
}

func GenerateCode(i *models.Invitation) {
	data := sha1.Sum([]byte(i.Name + time.Now().Format(time.RFC850)))
	encoded := hex.EncodeToString(data[:4])
	i.Code = strings.ToUpper(encoded[:4])
}

// @Title Create
// @Description create invitation
// @Param	body		body 	models.Invitation	true		"The invitation content"
// @Success 200 {string} models.Invitation.Id
// @Failure 500 body is empty
// @router / [post]
func (i *InvitationController) Post() {
	var in models.Invitation
	json.Unmarshal(i.Ctx.Input.RequestBody, &in)
	GenerateCode(&in)
	var invitationID uint
	var err bool
	if invitationID, err = models.CreateInvitation(in); err {
		i.Abort("500")
	}

	i.Data["json"] = map[string]uint{"InvitationID": invitationID}
	i.ServeJSON()
}

// invitation := Invitation{Name: "Usuario 1", MaxAdults: 2, MaxChildren: 0, Code: "F13S"}
