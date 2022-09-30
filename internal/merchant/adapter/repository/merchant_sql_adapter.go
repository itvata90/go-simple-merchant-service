package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	domain "merchant-service/internal/merchant/entity"
)

func NewMerchantSQLAdapter(db *sql.DB) *MerchantSQLAdapter {
	return &MerchantSQLAdapter{DB: db}
}

type MerchantSQLAdapter struct {
	DB *sql.DB
}

func (r *MerchantSQLAdapter) GetMerchants(ctx context.Context, pageSize int, pageIdx int) ([]domain.Merchant, int64, error) {
	var total int64
	stmt, err := r.DB.Prepare(`
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
from merchants limit ? offset ?`)
	if err != nil {
		return nil, total, err
	}
	rows, err := stmt.QueryContext(ctx, pageSize, pageIdx*pageSize)
	if err != nil {
		return nil, total, err
	}
	defer rows.Close()
	var merchants []domain.Merchant
	for rows.Next() {
		var merchant domain.Merchant
		err := rows.Scan(
			&merchant.Code,
			&merchant.ContactName,
			&merchant.Province,
			&merchant.District,
			&merchant.Street,
			&merchant.ContactEmail,
			&merchant.ContactPhoneNo,
			&merchant.OwnerId,
			&merchant.TaxId,
			&merchant.Status,
			&merchant.CreatedAt,
			&merchant.UpdatedAt,
		)
		if err != nil {
			return merchants, total, err
		}
		merchants = append(merchants, merchant)
	}

	stmt, err = r.DB.Prepare(`
		select 
			count(*) as total
		from merchants`)

	if err != nil {
		return nil, total, err
	}

	if err := stmt.QueryRow().Scan(&total); err != nil {
		return nil, total, err
	}

	return merchants, total, nil
}

func (r *MerchantSQLAdapter) GetMerchantByCode(ctx context.Context, code string) (*domain.Merchant, error) {
	var merchant domain.Merchant
	stmt, err := r.DB.Prepare(`
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
		from merchants where code = ? limit 1`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(code).Scan(
		&merchant.Code,
		&merchant.ContactName,
		&merchant.Province,
		&merchant.District,
		&merchant.Street,
		&merchant.ContactEmail,
		&merchant.ContactPhoneNo,
		&merchant.OwnerId,
		&merchant.TaxId,
		&merchant.Status,
		&merchant.CreatedAt,
		&merchant.UpdatedAt,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return nil, nil
		}
	}
	return &merchant, nil
}

func (r *MerchantSQLAdapter) CreateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	tx := GetTx(r, ctx)

	if notValid, err := inValidField(tx, "contact_phone_no", merchant.ContactPhoneNo, ""); err != nil {
		return -1, err
	} else if notValid {
		return -1, fmt.Errorf("duplicate phone %s", merchant.ContactPhoneNo)
	}
	if notValid, err := inValidField(tx, "contact_email", merchant.ContactEmail, ""); err != nil {
		return -1, err
	} else if notValid {
		return -1, fmt.Errorf("duplicate email %s", merchant.ContactEmail)
	}

	query := `
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

	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(
		merchant.Code,
		merchant.ContactName,
		merchant.Province,
		merchant.District,
		merchant.Street,
		merchant.ContactEmail,
		merchant.ContactPhoneNo,
		merchant.OwnerId,
		merchant.TaxId,
		merchant.Status,
		merchant.CreatedAt,
	)

	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (r *MerchantSQLAdapter) UpdateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	tx := GetTx(r, ctx)

	if notValid, err := inValidField(tx, "contact_phone_no", merchant.ContactPhoneNo, merchant.Code); err != nil {
		return -1, err
	} else if notValid {
		return 0, fmt.Errorf("duplicate phone %s", merchant.ContactPhoneNo)
	}
	if notValid, err := inValidField(tx, "contact_email", merchant.ContactEmail, merchant.Code); err != nil {
		return -1, err
	} else if notValid {
		return 0, fmt.Errorf("duplicate email %s", merchant.ContactEmail)
	}
	stmt, err := tx.Prepare(`
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
		where code = ?`)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(
		merchant.ContactName,
		merchant.Province,
		merchant.District,
		merchant.Street,
		merchant.ContactEmail,
		merchant.ContactPhoneNo,
		merchant.OwnerId,
		merchant.TaxId,
		merchant.Status,
		merchant.UpdatedAt,
		merchant.Code,
	)

	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (r *MerchantSQLAdapter) DeleteMerchant(ctx context.Context, code string) (int64, error) {
	query := "delete from merchants where code = ?"

	tx := GetTx(r, ctx)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(code)

	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func GetTx(r *MerchantSQLAdapter, ctx context.Context) *sql.Tx {
	tx, err := r.DB.Begin()

	if err != nil {
		return nil
	}

	return tx

}

func inValidField(obj interface{}, field string, value string, excludeValue string) (bool, error) {
	var notValid bool
	var stmt *sql.Stmt
	var err error
	query := fmt.Sprintf(`select if(count(*), 'true', 'false') as no_valid from merchants where %s = ? and not code = ?`, field)
	if db, ok := obj.(*sql.DB); ok {
		stmt, err = db.Prepare(query)
		if err != nil {
			return false, err
		}
	} else if tx, ok := obj.(*sql.Tx); ok {
		stmt, err = tx.Prepare(query)
		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("Unknow db handler type")
	}
	if err = stmt.QueryRow(value, excludeValue).Scan(&notValid); err != nil {
		return false, err
	}
	return notValid, nil
}
