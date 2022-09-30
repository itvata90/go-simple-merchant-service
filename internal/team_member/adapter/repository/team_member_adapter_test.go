package repository

import (
	"database/sql"
	"log"

	"context"
	domain "merchant-service/internal/team_member/entity"
	"testing"

	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

var t time.Time = time.Now()

var t1 = domain.TeamMember{
	Id:             "00000000-0000-0000-0000-00000000ID01",
	Username:       "User1",
	Password:       "password1",
	FirstName:      "Name1",
	LastName:       "LastName1",
	BirthDate:      "01/01/2020",
	Nationality:    "Nationality1",
	ContactEmail:   "email1@gmail.com",
	ContactPhoneNo: "0123456789",
	Province:       "Province1",
	District:       "District1",
	Street:         "Stree1",
	MerchantCode:   "MerchantCode1",
	Role:           "Role1",
	CreatedAt:      &t,
	UpdatedAt:      &t,
}

var dbField []string = []string{
	"id",
	"username",
	"first_name",
	"last_name",
	"birth_date",
	"nationality",
	"contact_email",
	"contact_phone_no",
	"province",
	"district",
	"street",
	"merchant_code",
	"role",
	"created_at",
	"updated_at",
}

var loadByIdQuery = `
select
	id,
	username,
	first_name,
	last_name,
	birth_date,
	nationality ,
	contact_email,
	contact_phone_no,
	province,
	district,
	street,
	merchant_code,
	role,
	created_at,
	updated_at
from team_members where id = 00000000-0000-0000-0000-00000000ID01 limit 1`

func TestLoadById(t *testing.T) {
	db, mock := NewMock()
	repo := &TeamMemberSQLAdapter{db}
	defer func() {
		repo.DB.Close()
	}()
	rows := sqlmock.NewRows(dbField).
		AddRow(
			t1.Id,
			t1.Username,
			t1.FirstName,
			t1.LastName,
			t1.BirthDate,
			t1.Nationality,
			t1.ContactEmail,
			t1.ContactPhoneNo,
			t1.Province,
			t1.District,
			t1.Street,
			t1.MerchantCode,
			t1.Role,
			t1.CreatedAt,
			t1.UpdatedAt,
		)
	mock.ExpectPrepare(loadByIdQuery).ExpectQuery().WithArgs(t1.Id).WillReturnRows(rows)
	merchant, err := repo.GetTeamMemberById(context.TODO(), t1.Id)
	assert.NotNil(t, merchant)
	assert.NoError(t, err)
}
