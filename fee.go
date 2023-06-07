package models

import "gorm.io/datatypes"

var EmptyFee = Fee{}

type Fee struct {
	FeeData datatypes.JSONMap `gorm:"jsonb" json:"feeData"`
}

func (*Fee) TableName() string {
	return "product_price_fees"
}
