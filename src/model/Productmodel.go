package model

import (
	"time"

	httperors "github.com/myrachanto/custom-http-error"
	"gorm.io/gorm"
)

var ExpiresAt = time.Now().Add(time.Minute * 100000).Unix()

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	BuyPrice    float64 `json:"buy_price"`
	SellPrice   float64 `json:"sell_price"`
	Quantity    int64   `json:"quantity"`
	Picture     string  `json:"picture"`
	Available   bool    `json:"available"`
	Usercode    string  `json:"usercode"`
	Productcode string  `json:"code"`
	gorm.Model
}

//Validate ..
func (product Product) Validate() httperors.HttpErr {
	if product.Name == "" {
		return httperors.NewNotFoundError("Invalid Product Name")
	}
	if product.Description == "" {
		return httperors.NewNotFoundError("Invalid Description")
	}
	if product.Picture == "" {
		return httperors.NewNotFoundError("Invalid Picture")
	}
	return nil
}
