package repository

import (
	"context"
	"database/sql"
	domain "merchant-service/internal/merchant/entity"
	"time"
)

func NewMerchantSQLAdapter(db *sql.DB) *MerchantSQLAdapter {
	return &MerchantSQLAdapter{DB: db}
}

type MerchantSQLAdapter struct {
	DB *sql.DB
}

func (r *MerchantSQLAdapter) GetMerchants(ctx context.Context, pageSize int, pageIdx int) ([]domain.Merchant, error) {
	stmt, err := r.DB.Prepare(`
		select 
			code,
			name,
			province,
			country,
			address,
			email,
			phone,
			status,
			created_at,
			updated_at
		from merchants limit ? offset ?`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, pageSize, pageIdx*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var merchants []domain.Merchant
	for rows.Next() {
		var merchant domain.Merchant
		err := rows.Scan(
			&merchant.Code,
			&merchant.Name,
			&merchant.Province,
			&merchant.Country,
			&merchant.Address,
			&merchant.Email,
			&merchant.Phone,
			&merchant.Status,
			&merchant.CreatedAt,
			&merchant.UpdatedAt,
		)
		if err != nil {
			return merchants, err
		}
		merchants = append(merchants, merchant)
	}
	if err != nil {
		return merchants, err
	}
	return merchants, nil
}

func (r *MerchantSQLAdapter) GetMerchantByCode(ctx context.Context, code string) (*domain.Merchant, error) {
	var merchant domain.Merchant
	stmt, err := r.DB.Prepare(`
		select 
			code,
			name,
			province,
			country,
			address,
			email,
			phone,
			status,
			created_at,
			updated_at
		from merchants where code = ? limit 1`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(code).Scan(
		&merchant.Code,
		&merchant.Name,
		&merchant.Province,
		&merchant.Country,
		&merchant.Address,
		&merchant.Email,
		&merchant.Phone,
		&merchant.Status,
		&merchant.CreatedAt,
		&merchant.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &merchant, nil
}

func (r *MerchantSQLAdapter) CreateMerchant(ctx context.Context, merchant *domain.Merchant) (int64, error) {
	merchant.CreatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	query := `
		insert 
			into merchants (
				code,
				name,
				province,
				country,
				address,
				email,
				phone,
				status,
				created_at) 
		values (?,?,?,?,?,?,?,?,?)`

	tx := GetTx(r, ctx)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(
		ctx,
		merchant.Code,
		merchant.Name,
		merchant.Province,
		merchant.Country,
		merchant.Address,
		merchant.Email,
		merchant.Phone,
		merchant.Status,
		merchant.CreatedAt)

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
	merchant.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	query := `
		update merchants 
		set 
			name = ?,
			province = ?,
			district = ?,
			street = ?,
			email = ?,
			phone = ?,
			status = ?,
			updated_at = ?
		where code = ?`

	tx := GetTx(r, ctx)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(
		ctx,
		merchant.Name,
		merchant.Province,
		merchant.Country,
		merchant.Address,
		merchant.Email,
		merchant.Phone,
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

	res, err := stmt.ExecContext(ctx, code)

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
