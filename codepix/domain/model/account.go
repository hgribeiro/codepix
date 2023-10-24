package model


type Account struct {
	Base `valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank *Bank `valid:"-"`
	Numer string `json:"number" valid:"notnull"`



	func (account *Account) isValid() error {
		_, err := govalidator.ValidateStruct(bank)
			if err!= nil {
				return err
			}
			
		return nil
	}

	
func NewAccount(bank *Bank, number string,) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
}