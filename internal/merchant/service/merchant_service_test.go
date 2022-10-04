package service

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"merchant-service/internal/merchant/adapter/repository"
	domain "merchant-service/internal/merchant/entity"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
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

var m1 = domain.Merchant{
	Code:           uuid.New().String(),
	ContactName:    "TestContactContactName",
	Province:       "Provine1",
	District:       "District1",
	Street:         "Street1",
	ContactEmail:   "ContactEmail1@gmail.com",
	ContactPhoneNo: "0123456789",
	OwnerId:        uuid.New().String(),
	TaxId:          "123455555",
	Status:         "Active",
	CreatedAt:      &t,
	UpdatedAt:      &t,
}

var m2 = domain.Merchant{
	Code:           uuid.New().String(),
	ContactName:    "TestContactContactName2",
	Province:       "Provine2",
	District:       "District2",
	Street:         "Street2",
	ContactEmail:   "ContactEmail2@gmail.com",
	ContactPhoneNo: "0123456781",
	OwnerId:        uuid.New().String(),
	TaxId:          "12345978",
	Status:         "Active",
	CreatedAt:      &t,
	UpdatedAt:      &t,
}

var dbField []string = []string{
	"code",
	"contact_name",
	"province",
	"district",
	"street",
	"contact_email",
	"contact_phone_no",
	"owner_id",
	"tax_id",
	"status",
	"created_at",
	"updated_at",
}

var selectQuery = `
select
	code,
	contact_name,
	province,
	district,
	street,
	contact_email,
	contact_phone_no,
	owner_id,
	tax_id,
	status,
	created_at,
	updated_at
from merchants where code = ? limit 1`

var selectAllQuery = `
select
	code,
	contact_name,
	province,
	district,
	street,
	contact_email,
	contact_phone_no,
	owner_id,
tax_id,
	status,
	created_at,
	updated_at
from merchants limit ? offset ?`

var countQuery = `
select
	count(*) as total
from merchants`

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type AnyId struct{}

func (a AnyId) Match(v driver.Value) bool {
	s, ok := v.(string)
	if ok {
		_, err := uuid.Parse(s)
		if err != nil {
			ok = false
		}
	}
	return ok
}

func TestGetMerchants(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewMerchantSQLAdapter(db)
	service := &merchantService{repository: repo}
	defer func() {
		repo.DB.Close()
	}()
	rows := sqlmock.NewRows(dbField).
		AddRow(
			m1.Code,
			m1.ContactName,
			m1.Province,
			m1.District,
			m1.Street,
			m1.ContactEmail,
			m1.ContactPhoneNo,
			m1.OwnerId,
			m1.TaxId,
			m1.Status,
			m1.CreatedAt,
			m1.UpdatedAt,
		).AddRow(
		m2.Code,
		m2.ContactName,
		m2.Province,
		m2.District,
		m2.Street,
		m2.ContactEmail,
		m2.ContactPhoneNo,
		m2.OwnerId,
		m2.TaxId,
		m2.Status,
		m2.CreatedAt,
		m2.UpdatedAt,
	)
	counts := sqlmock.NewRows([]string{"total"}).AddRow(2)
	mock.ExpectPrepare(selectAllQuery).ExpectQuery().WithArgs(5, 0).WillReturnRows(rows)
	mock.ExpectPrepare(countQuery).ExpectQuery().WillReturnRows(counts)
	merchants, total, err := service.repository.GetMerchants(context.TODO(), 5, 0)
	assert.Equal(t, int64(2), total)
	assert.NotNil(t, merchants)
	assert.NoError(t, err)
}

func TestGetMerchantByCode(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewMerchantSQLAdapter(db)
	service := &merchantService{repository: repo}
	defer func() {
		repo.DB.Close()
	}()
	rows := sqlmock.NewRows(dbField).
		AddRow(
			m1.Code,
			m1.ContactName,
			m1.Province,
			m1.District,
			m1.Street,
			m1.ContactEmail,
			m1.ContactPhoneNo,
			m1.OwnerId,
			m1.TaxId,
			m1.Status,
			m1.CreatedAt,
			m1.UpdatedAt,
		)
	mock.ExpectPrepare(selectQuery).ExpectQuery().WithArgs(m1.Code).WillReturnRows(rows)
	merchant, err := service.GetMerchantByCode(context.TODO(), m1.Code)
	assert.NotNil(t, merchant)
	assert.NoError(t, err)
}

func TestGetMerchantByCodeErr(t *testing.T) {
	m1.Code = "" // creating new don't need a guid
	db, mock := NewMock()
	repo := repository.NewMerchantSQLAdapter(db)
	service := &merchantService{repository: repo}
	defer func() {
		repo.DB.Close()
	}()
	rows := sqlmock.NewRows(dbField)
	randomId := uuid.New().String()
	mock.ExpectPrepare(selectQuery).ExpectQuery().WithArgs(randomId).WillReturnRows(rows)
	merchant, err := service.GetMerchantByCode(context.TODO(), randomId)
	assert.Empty(t, merchant, fmt.Sprintf("test code: %+v", randomId))
	assert.Empty(t, err, fmt.Sprintf("error: %v", err))
}

func validQuery(field string) string {
	return fmt.Sprintf(`select if(count(*), 'true', 'false') as no_valid from merchants where %s = ? and not code = ?`, field)
}

var insertQuery = `
insert
	into merchants (
		code,
		contact_name,
		province,
		district,
		street,
		contact_email,
		contact_phone_no,
		owner_id,
		tax_id,		
		status,
		created_at)
values (?,?,?,?,?,?,?,?,?,?,?)`

var updateQuery = `
update merchants
set
	contact_name = ?,
	province = ?,
	district = ?,
	street = ?,
	contact_email = ?,
	contact_phone_no = ?,
	owner_id = ?,
	tax_id = ?,
	status = ?,
	updated_at = ?
where code = ?`

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewMerchantSQLAdapter(db)
	service := &merchantService{repository: repo}
	defer func() {
		repo.DB.Close()
	}()
	mock.ExpectBegin()
	row1 := sqlmock.NewRows([]string{"not_valid"}).AddRow("false")
	mock.ExpectPrepare(validQuery("contact_phone_no")).ExpectQuery().WithArgs(m1.ContactPhoneNo, m1.Code).WillReturnRows(row1)
	row2 := sqlmock.NewRows([]string{"not_valid"}).AddRow("false")
	mock.ExpectPrepare(validQuery("contact_email")).ExpectQuery().WithArgs(m1.ContactEmail, m1.Code).WillReturnRows(row2)
	mock.ExpectPrepare(updateQuery).ExpectExec().WithArgs(
		m1.ContactName,
		m1.Province,
		m1.District,
		m1.Street,
		m1.ContactEmail,
		m1.ContactPhoneNo,
		m1.OwnerId,
		m1.TaxId,
		m1.Status,
		AnyTime{},
		m1.Code,
	).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	rowsAffected, err := service.UpdateMerchant(context.TODO(), &m1)
	assert.Equal(t, int64(1), rowsAffected)
	assert.NoError(t, err)
}

var deleteQuery = `delete from merchants where code = ?`

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	repo := repository.NewMerchantSQLAdapter(db)
	service := &merchantService{repository: repo}
	defer func() {
		repo.DB.Close()
	}()
	mock.ExpectBegin()
	mock.ExpectPrepare(deleteQuery).ExpectExec().WithArgs(m1.Code).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	rowsAffected, err := service.DeleteMerchant(context.TODO(), m1.Code)
	assert.Equal(t, int64(1), rowsAffected)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	m1.Code = "" // creating new don't need a guid
	db, mock := NewMock()
	repo := repository.NewMerchantSQLAdapter(db)
	service := &merchantService{repository: repo}
	defer func() {
		repo.DB.Close()
	}()
	mock.ExpectBegin()
	row1 := sqlmock.NewRows([]string{"not_valid"}).AddRow("false")
	mock.ExpectPrepare(validQuery("contact_phone_no")).ExpectQuery().WithArgs(m1.ContactPhoneNo, m1.Code).WillReturnRows(row1)
	row2 := sqlmock.NewRows([]string{"not_valid"}).AddRow("false")
	mock.ExpectPrepare(validQuery("contact_email")).ExpectQuery().WithArgs(m1.ContactEmail, m1.Code).WillReturnRows(row2)
	mock.ExpectPrepare(insertQuery).
		ExpectExec().
		WithArgs(
			AnyId{},
			m1.ContactName,
			m1.Province,
			m1.District,
			m1.Street,
			m1.ContactEmail,
			m1.ContactPhoneNo,
			m1.OwnerId,
			m1.TaxId,
			m1.Status,
			AnyTime{},
		).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	rowsAffected, err := service.CreateMerchant(context.TODO(), &m1)
	assert.Equal(t, int64(1), rowsAffected)
	assert.NoError(t, err)
}
