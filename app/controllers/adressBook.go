package controllers

import (
	"encoding/json"
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
	"strconv"
	"strings"
)

var Address = new(AddressBookController)

type AddressBookController struct {
	BaseController
}

// View a list of address maps
func (ctl *AddressBookController) List() {
	ack := dto.AbGetAck{}
	ack.Tags = []string{}
	// Inquire tags
	tag_colors := dto.AbTag_colors{}
	tags := services.Tags.FindTags(ctl.loginUserInfo.Id)
	for _, item := range tags {
		ack.Tags = append(ack.Tags, item.Tag)

		if item.Color != "" {
			tag_colors[item.Tag], _ = strconv.ParseInt(item.Color, 10, 64)
		}
	}
	jdata_tag_colors, _ := json.Marshal(tag_colors)
	ack.Tag_colors = string(jdata_tag_colors)

	// Query peers
	ack.Peers = []dto.AbGetPeer{}
	peerDbs := services.Peers.FindPeers(ctl.loginUserInfo.Id)
	for _, item := range peerDbs {
		ack.Peers = append(ack.Peers, dto.AbGetPeer{
			Id:       item.ClientId,
			Username: item.Username,
			Hostname: item.Hostname,
			Alias:    item.Alias,
			Platform: item.Platform,
			Tags:     strings.Split(item.Tags, ","),
		})
	}

	// Query the list of all logged-in accounts
	tokens := services.Token.FindTokens(ctl.loginUserInfo.Id)
	for _, item := range *tokens {
		ist := false
		for _, bookItem := range ack.Peers {
			if bookItem.Id == item.ClientId {
				ist = true
				break
			}
		}
		if !ist {
			ack.Peers = append(ack.Peers, dto.AbGetPeer{
				Id:       item.ClientId,
				Username: "----",
				Hostname: item.ClientId,
				Alias:    "id:" + item.ClientId,
				Platform: "os",
				Tags:     strings.Split("", ","),
			})
		}
	}

	jdata, _ := json.Marshal(ack)

	ctl.JSON(beegoHelper.H{
		//"error":     false,
		"data": string(jdata),
		//"update_at": time.Now().Format("2006-01-02 15:04:05"),
	})
}

// Update the address spectrum
func (ctl *AddressBookController) Update() {
	req := dto.AbUpdateReq{}

	if err := ctl.BindJSON(&req); err != nil {
		ctl.JSON(beegoHelper.H{
			"error": "The request parameter is abnormal",
		})
		return
	}

	// Parse the data
	reqSub := &dto.AbUpdateSub{}
	err := json.Unmarshal([]byte(req.Data), reqSub)
	if err != nil {
		ctl.JSON(beegoHelper.H{
			"error": "The request data is abnormal",
		})
	}

	// Delete tags in batches
	services.Tags.DeleteAll(ctl.loginUserInfo.Id)
	// Delete peers in batches
	services.Peers.DeleteAll(ctl.loginUserInfo.Id)

	// Start inserting tags in bulk
	if !services.Tags.BatchAdd(ctl.loginUserInfo.Id, reqSub.Tags, reqSub.Tag_colors) {
		ctl.JSON(beegoHelper.H{
			"error": "Failed to import tags",
		})
	}
	// Start inserting peers in bulk
	if !services.Peers.BatchAdd(ctl.loginUserInfo.Id, reqSub.Peers) {
		ctl.JSON(beegoHelper.H{
			"error": "Failed to import the address book",
		})
	}

	ctl.JSON(beegoHelper.H{
		"data": "success",
	})

}
