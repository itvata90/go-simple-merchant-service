package domain

import (
	"time"
)

type TeamMember struct {
	Id             string     `json:"id,omitempty"`
	Username       string     `json:"username,omitempty"`
	Password       string     `json:"password,omitempty"`
	FirstName      string     `json:"firstName,omitempty"`
	LastName       string     `json:"lastName,omitempty"`
	BirthDate      string     `json:"birthDate,omitempty"`
	Nationality    string     `json:"nationality,omitempty"`
	ContactEmail   string     `json:"contactEmail,omitempty"`
	ContactPhoneNo string     `json:"contactPhoneNo,omitempty"`
	Province       string     `json:"province,omitempty"`
	District       string     `json:"district,omitempty"`
	Street         string     `json:"street,omitempty"`
	MerchantCode   string     `json:"merchantCode,omitempty"`
	Role           string     `json:"role,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
}

type GetTeamMembersResponse struct {
	Data   []TeamMember `json:"data"`
	Paging Pagination   `json:"pagination"`
}

type Pagination struct {
	Total     int64 `json:"total"`
	PageIndex int   `json:"pageIndex"`
	PageSize  int   `json:"pageSize"`
}
