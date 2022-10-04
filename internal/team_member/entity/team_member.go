package domain

import (
	"time"
)

type TeamMember struct {
	Id             string     `json:"id" gorm:"column:id;primary_key"`
	Username       string     `json:"username" gorm:"column:username"`
	Password       string     `json:"password" gorm:"column:password"`
	FirstName      string     `json:"firstName" gorm:"column:first_name"`
	LastName       string     `json:"lastName" gorm:"column:last_name"`
	BirthDate      *time.Time `json:"birthDate,omitempty" gorm:"column:birth_date"`
	Nationality    string     `json:"nationality,omitempty" gorm:"column:nationality"`
	ContactEmail   string     `json:"contactEmail" gorm:"column:contact_email"`
	ContactPhoneNo string     `json:"contactPhoneNo" gorm:"column:contact_phone_no"`
	Province       string     `json:"province,omitempty" gorm:"column:province"`
	District       string     `json:"district,omitempty" gorm:"column:district"`
	Street         string     `json:"street,omitempty" gorm:"column:street"`
	MerchantCode   string     `json:"merchantCode" gorm:"column:merchant_code"`
	Role           string     `json:"role,omitempty" gorm:"column:role"`
	CreatedAt      *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
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
