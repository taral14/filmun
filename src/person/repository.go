package person

import (
	"github.com/jmoiron/sqlx"
	"github.com/taral14/filmun/src/entity"
)

type rowPerson struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	OriginalName string `db:"original_name"`
	KpInternalId int    `db:"kp_internal_id"`
	ImageUrl     string `db:"image_url"`
	KpImageUrl   string `db:"kp_image_url"`
}

type MysqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository(db *sqlx.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (r MysqlRepository) FindOneById(id int) (entity.Person, error) {
	return r.queryOne("SELECT * FROM tbl_person WHERE id=?", id)
}

func (r MysqlRepository) FindAll(limit, offset int) ([]entity.Person, error) {
	return r.queryAll("SELECT * FROM tbl_person LIMIT ? OFFSET ?", limit, offset)
}

func (r MysqlRepository) FindByName(name string, limit, offset int) ([]entity.Person, error) {
	return r.queryAll("SELECT * FROM tbl_person WHERE name LIKE ? LIMIT ? OFFSET ?", "%"+name+"%", limit, offset)
}

func (r MysqlRepository) queryOne(query string, params ...interface{}) (entity.Person, error) {
	row := rowPerson{}
	err := r.db.Get(&row, query, params...)
	if err != nil {
		return entity.Person{}, err
	}
	return mapToPerson(row)
}

func (r MysqlRepository) queryAll(query string, params ...interface{}) ([]entity.Person, error) {
	rows := []rowPerson{}
	err := r.db.Select(&rows, query, params...)
	if err != nil {
		return []entity.Person{}, err
	}
	persons := make([]entity.Person, 0, len(rows))
	for _, row := range rows {
		person, err := mapToPerson(row)
		if err != nil {
			return persons, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

func mapToPerson(row rowPerson) (entity.Person, error) {
	return entity.Person{
		ID:           row.ID,
		Name:         row.Name,
		OriginalName: row.OriginalName,
		ImageUrl:     row.ImageUrl,
		KpInternalId: row.KpInternalId,
	}, nil
}
