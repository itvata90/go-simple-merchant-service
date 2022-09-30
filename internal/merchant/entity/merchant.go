package domain

import (
	"database/sql"
)

type MerchantStatus string

const (
	MerchantActive   MerchantStatus = "Active"
	MerchantInActive MerchantStatus = "Inactive"
)

type Merchant struct {
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	Country   string         `json:"country"`
	Province  string         `json:"province"`
	Address   string         `json:"address,omitempty"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Status    MerchantStatus `json:"status"`
	CreatedAt sql.NullTime   `json:"create_at"`
	CreatedBy string         `json:"created_by,omitempty"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	UpdatedBy string         `json:"updated_by,omitempty"`
}
