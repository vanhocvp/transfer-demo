package models

type ViettelMoney struct {
	ID          int
	DriverID    string
	PhoneNumber string
	ViettelPay  float64
	Money       float64
}

func GetBalance(PhoneNumber string) (*ViettelMoney, error) {
	viettelMoney := ViettelMoney{}
	err := db.Model(ViettelMoney{}).Where("phone_number").First(&viettelMoney).Error
	if err != nil {
		return nil, err
	}

	return &viettelMoney, nil
}
