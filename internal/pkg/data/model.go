package data

type Ad struct {
	ID    uint   `json:"id,omitempty"`
	Brand string `json:"brand,omitempty"`
	Model string `json:"model,omitempty"`
	Color string `json:"color,omitempty"`
	Price int    `json:"price,omitempty"`
}

//Geter`s
func (a *Ad) GetID() uint {
	return a.ID
}

func (a *Ad) SetID(id uint) {
	a.ID = id
}

func (a *Ad) GetBrand() string {
	return a.Brand
}

func (a *Ad) GetModel() string {
	return a.Model
}

func (a *Ad) GetColor() string {
	return a.Color
}

func (a *Ad) GetPrice() int {
	return a.Price
}

type Account struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token"`
}

func (a *Account) GetID() int {
	return a.ID
}

func (a *Account) GetUserName() string {
	return a.Username
}

func (a *Account) GetPassword() string {
	return a.Password
}

func (a *Account) GetToken() string {
	return a.Token
}

func (a *Account) SetToken(token string) {
	a.Token = token
}
