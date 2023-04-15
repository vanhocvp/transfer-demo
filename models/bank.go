package models

type BankAccount struct {
	ID            int
	AccountNumber string
	CardNumber    string
	BankName      string
	CustomerName  string
	Balance       float64
}

func GetAccountNumberInfo(accountNumber string, bankName string) (*BankAccount, error) {
	account := BankAccount{}
	err := db.Model(BankAccount{}).Where("account_number = ? AND bank_name = ?", accountNumber, bankName).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil

}
