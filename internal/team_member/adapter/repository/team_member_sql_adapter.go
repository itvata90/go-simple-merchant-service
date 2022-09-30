package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	domain "merchant-service/internal/team_member/entity"

	q "github.com/core-go/sql"
)

const (
	teamMembersTable = "team_members"
)

func NewTeamMemberAdapter(db *sql.DB) *TeamMemberSQLAdapter {
	return &TeamMemberSQLAdapter{DB: db}
}

type TeamMemberSQLAdapter struct {
	DB *sql.DB
}

func (r *TeamMemberSQLAdapter) GetTeamMembers(ctx context.Context, pageSize int, pageIndex int) ([]domain.TeamMember, int64, error) {
	var total int64
	var members []domain.TeamMember
	queryMembers := "select id, username, first_name, last_name, birth_date, nationality, contact_email, contact_phone_no, province, district, street, merchant_code, tax_id, role, created_at, updated_at from team_members limit ? offset ?"
	err := q.Query(ctx, r.DB, nil, &members, queryMembers, pageSize, pageIndex*pageSize)
	if err != nil {
		return nil, total, err
	}
	if len(members) > 0 {
		queryTotal := "select count(*) as total from team_members"
		row := r.DB.QueryRow(queryTotal)
		err = row.Scan(&total)
		if err != nil {
			return nil, total, err
		}

		return members, total, nil
	}

	return members, total, err
}

func (r *TeamMemberSQLAdapter) GetTeamMemberById(ctx context.Context, id string) (*domain.TeamMember, error) {
	var member []domain.TeamMember
	queryMember := fmt.Sprintf("select id, username, first_name, last_name, birth_date, nationality, contact_email, contact_phone_no, province, district, street, merchant_code, tax_id, role, created_at, updated_at from team_members where id = %s limit 1", q.BuildParam(1))
	err := q.Query(ctx, r.DB, nil, &member, queryMember, id)
	if err != nil {
		return nil, err
	}
	if len(member) > 0 {
		return &member[0], nil
	}
	return nil, nil
}

func (r *TeamMemberSQLAdapter) GetTeamMemberByMerchantCode(ctx context.Context, merchantCode string, pageSize int, pageIndex int) ([]domain.TeamMember, int64, error) {
	var total int64
	var members []domain.TeamMember
	queryMember := "select id, username, first_name, last_name, birth_date, nationality, contact_email, contact_phone_no, province, district, street, merchant_code, role, created_at, updated_at from team_members where merchant_code = ? limit ? offset ?"
	err := q.Query(ctx, r.DB, nil, &members, queryMember, merchantCode, pageSize, pageIndex*pageSize)
	if err != nil {
		return nil, total, err
	}
	if len(members) > 0 {
		queryTotal := "select count(*) as total from team_members where merchant_code = ?"
		row := r.DB.QueryRow(queryTotal, merchantCode)
		err = row.Scan(&total)
		if err != nil {
			return nil, total, err
		}

		return members, total, nil
	}

	return members, total, err

}

func (r *TeamMemberSQLAdapter) CreateTeamMember(ctx context.Context, teamMember *domain.TeamMember) (int64, error) {
	tx := GetTx(r, ctx)

	if notValid, err := inValidField(tx, "contact_phone_no", teamMember.ContactPhoneNo, ""); err != nil {
		return -1, err
	} else if notValid {
		return 0, fmt.Errorf("duplicate phone %s", teamMember.ContactPhoneNo)
	}
	if notValid, err := inValidField(tx, "contact_email", teamMember.ContactEmail, ""); err != nil {
		return -1, err
	} else if notValid {
		return 0, fmt.Errorf("duplicate email %s", teamMember.ContactEmail)
	}

	query, args := q.BuildToInsert(teamMembersTable, teamMember, q.BuildParam)
	res, err := tx.ExecContext(ctx, query, args...)

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

func (r *TeamMemberSQLAdapter) UpdateTeamMember(ctx context.Context, teamMember *domain.TeamMember) (int64, error) {
	tx := GetTx(r, ctx)

	if notValid, err := inValidField(tx, "contact_phone_no", teamMember.ContactPhoneNo, teamMember.Id); err != nil {
		return 0, err
	} else if notValid {
		return 0, fmt.Errorf("duplicate phone %s", teamMember.ContactPhoneNo)
	}
	if notValid, err := inValidField(tx, "contact_email", teamMember.ContactEmail, teamMember.Id); err != nil {
		return 0, err
	} else if notValid {
		return 0, fmt.Errorf("duplicate email %s", teamMember.ContactEmail)
	}

	query, args := q.BuildToUpdate(teamMembersTable, teamMember, q.BuildParam)
	res, err := tx.ExecContext(ctx, query, args...)

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

func (r *TeamMemberSQLAdapter) DeleteTeamMember(ctx context.Context, id string) (int64, error) {
	tx := GetTx(r, ctx)

	query := fmt.Sprintf("delete from team_members where id = %s", q.BuildParam(1))
	res, err := tx.ExecContext(ctx, query, id)

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

func GetTx(r *TeamMemberSQLAdapter, ctx context.Context) *sql.Tx {
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
	query := fmt.Sprintf(`select if(count(*), 'true', 'false') as is_valid from team_members where %s = ? and not id = ?`, field)
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
