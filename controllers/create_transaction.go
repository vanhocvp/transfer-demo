package controllers

type CreateTransactionDraftRequest struct {
	UserID             *string `json:"user_id" binding:"required"`
	ReceiverID         *string `json:"receiver_id"`
	AccountNumber      *string `json:"account_number"`
	CardNumber         *string `json:"card_number"`
	PhoneNumber        *string `json:"phone_number"`
	BankName           *string `json:"bank_name"`
	ReceiverName       *string `json:"receiver_name"`
	PaymentType        *string `json:"payment_type"`
	PaymentSource      *int    `json:"payment_source"`
	PaymentDestination *int    `json:"payment_destination"`
}

//func CreateTransactionDraft(c *gin.Context) {
//	var form CreateTransactionDraftRequest
//
//	if err := c.ShouldBind(&form); err != nil {
//		log.Printf("[error] CreateTransactionDraft | Bad request %v", err)
//		c.JSON(http.StatusOK, gin.H{
//			"status": setting.AppSetting.StatusError,
//			"msg":    "Bad request",
//		})
//		return
//	}
//	log.Printf("[info] CreateTransactionDraft | form: %v", form)
//
//	err := models.Create
//
//}
