package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/taral14/filmun/src/entity"
)

type MysqlRepository struct {
	db *sqlx.DB
}

type rowUser struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func NewMysqlRepository(db *sqlx.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (r MysqlRepository) FindByUsername(username string) (*entity.User, error) {
	sql := "SELECT t.* FROM tbl_user t WHERE t.username=?"
	return r.queryOne(sql, username)
}

func (r MysqlRepository) FindById(id int) (*entity.User, error) {
	sql := "SELECT t.* FROM tbl_user t WHERE t.id=?"
	return r.queryOne(sql, id)
}

func (r MysqlRepository) queryOne(query string, params ...interface{}) (*entity.User, error) {
	row := rowUser{}
	err := r.db.Get(&row, query, params...)
	if err != nil {
		return new(entity.User), err
	}
	return mapToUser(row)
}

func (r MysqlRepository) queryAll(query string, params ...interface{}) ([]*entity.User, error) {
	rows := []rowUser{}
	err := r.db.Select(&rows, query, params...)
	if err != nil {
		return []*entity.User{}, err
	}
	users := make([]*entity.User, 0, len(rows))
	for _, row := range rows {
		user, err := mapToUser(row)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func mapToUser(row rowUser) (*entity.User, error) {
	return &entity.User{
		ID:       row.ID,
		Username: row.Username,
		Password: row.Password,
	}, nil
}
