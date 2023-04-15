package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vanhocvp/junctionx-hackathon/transfer-demo/setting"
	"log"
	"net/http"
)

type GetListRecipientRequest struct {
	SenderID   string  `json:"sender_id"`
	ReceiverID *string `json:"receiver_id"`
}

func GetListRecipient(c *gin.Context) {
	var form GetListRecipientRequest

	if err := c.ShouldBind(&form); err != nil {
		log.Printf("[error] GetListRecipientRequest | Bad request %v", err)
		c.JSON(http.StatusOK, gin.H{
			"status": setting.AppSetting.StatusError,
			"msg":    "Bad request",
		})
		return
	}

	log.Printf("[info] GetListRecipient | form: %v ", form)
	if form.ReceiverID == nil {

	}
}
