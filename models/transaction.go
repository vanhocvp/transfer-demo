package models

type Transaction struct {
	ID                 int `gorm:"primaryKey;index"`
	SenderUserID       string
	ReceiverUserID     string
	PaymentSource      int
	PaymentDestination int
	PaymentType        string
	AccountNumber      string
	PhoneNumber        string
	CardNumber         string
	BankName           string
	ReceiverName       string
	Amount             int
	Content            int
	Otp                string
	VoiceBio           bool
	Status             int
	CreateAt           int64 `gorm:"autoCreateTime:milli;index;not null"`
	UpdateAt           int64 `gorm:"autoUpdateTime:milli;index"`
}

func GetTransactionByUserID(userID string, accountNumber *string, cardNumber *string, phoneNumber *string, receiverName *string) (*Transaction, error) {
	transaction := Transaction{}
	query := db.Model(Transaction{}).Where("sender_user_id = ?", userID)
	if accountNumber != nil && *accountNumber != "" {
		query = query.Where("account_number = ?", *accountNumber)
	}
	if cardNumber != nil && *cardNumber != "" {
		query = query.Where("card_number = ?", *cardNumber)
	}
	if phoneNumber != nil && *phoneNumber != "" {
		query = query.Where("phone_number = ?", *phoneNumber)
	}
	if receiverName != nil && *receiverName != "" {
		query = query.Where("receiver_name = ?", *receiverName)
	}
	err := query.First(&transaction).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// func CreateTransaction()
