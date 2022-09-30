package domain

import (
	"time"
)

type Merchant struct {
	Code           string     `json:"code"`
	ContactName    string     `json:"contactName"`
	Province       string     `json:"province,omitempty"`
	District       string     `json:"district,omitempty"`
	Street         string     `json:"street,omitempty"`
	ContactEmail   string     `json:"contactEmail,omitempty"`
	ContactPhoneNo string     `json:"contactPhoneNo,omitempty"`
	OwnerId        string     `json:"ownerId,omitempty"`
	TaxId          string     `json:"taxId,omitempty"`
	Status         string     `json:"status,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	CreatedBy      string     `json:"created_by,omitempty"`
	UpdatedBy      string     `json:"updated_by,omitempty"`
}

type GetMerchantsResponse struct {
	Data   []Merchant `json:"data"`
	Paging Pagination `json:"pagination"`
}

type Pagination struct {
	Total     int64 `json:"total"`
	PageIndex int   `json:"pageIndex"`
	PageSize  int   `json:"pageSize"`
}
