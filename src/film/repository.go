package film

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/taral14/filmun/src/entity"
)

type MysqlRepository struct {
	db *sqlx.DB
}

type rowFilm struct {
	ID             int     `db:"id"`
	Name           string  `db:"name"`
	Url            string  `db:"url"`
	OriginalName   string  `db:"original_name"`
	KpInternalId   int     `db:"kp_internal_id"`
	ImdbInternalId string  `db:"imdb_internal_id"`
	Year           int     `db:"year"`
	CountSeasons   int     `db:"count_seasons"`
	IsSeries       int     `db:"is_series"`
	ImdbMark       float32 `db:"imdb_mark"`
	ImdbMarkVotes  int     `db:"imdb_mark_votes"`
	ImageUrl       string  `db:"image_url"`
	KpImageUrl     string  `db:"kp_image_url"`
	Description    string  `db:"description"`
	Status         int     `db:"status"`
	PremiereAt     int64   `db:"premiere_at"`
	Release720pAt  int64   `db:"release_720p_at"`
	Release1080pAt int64   `db:"release_1080p_at"`
	CreatedAt      int64   `db:"created_at"`
}

func NewMysqlRepository(db *sqlx.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (r MysqlRepository) FindByActorId(actorId int, limit, offset int) ([]entity.Film, error) {
	sql := "SELECT t.* FROM tbl_film t JOIN tbl_film_actor a ON a.film_id=t.id WHERE a.person_id=? ORDER BY t.premiere_at DESC LIMIT ? OFFSET ?"
	return r.queryAll(sql, actorId, limit, offset)
}

func (r MysqlRepository) FindAll(limit, offset int) ([]entity.Film, error) {
	sql := "SELECT t.* FROM tbl_film t WHERE t.is_series=0 ORDER BY t.premiere_at DESC LIMIT ? OFFSET ?"
	return r.queryAll(sql, limit, offset)
}

func (r MysqlRepository) FindByDirectorId(directorId int, limit, offset int) ([]entity.Film, error) {
	sql := "SELECT t.* FROM tbl_film t JOIN tbl_film_director d ON d.film_id=t.id WHERE d.person_id=? ORDER BY t.premiere_at DESC LIMIT ? OFFSET ?"
	return r.queryAll(sql, directorId, limit, offset)
}

func (r MysqlRepository) queryOne(query string, params ...interface{}) (entity.Film, error) {
	row := rowFilm{}
	err := r.db.Get(&row, query, params...)
	if err != nil {
		return entity.Film{}, err
	}
	return mapToFilm(row)
}

func (r MysqlRepository) queryAll(query string, params ...interface{}) ([]entity.Film, error) {
	rows := []rowFilm{}
	err := r.db.Select(&rows, query, params...)
	if err != nil {
		return []entity.Film{}, err
	}
	films := make([]entity.Film, 0, len(rows))
	for _, row := range rows {
		person, err := mapToFilm(row)
		if err != nil {
			return films, err
		}
		films = append(films, person)
	}
	return films, nil
}

func mapToFilm(row rowFilm) (entity.Film, error) {
	return entity.Film{
		ID:             row.ID,
		Name:           row.Name,
		OriginalName:   row.OriginalName,
		Imdb:           entity.FilmImdb{Mark: row.ImdbMark, Votes: row.ImdbMarkVotes},
		IsSeries:       row.IsSeries,
		ImageUrl:       row.ImageUrl,
		Description:    row.Description,
		Status:         entity.FilmStatus(row.Status),
		Release720pAt:  time.Unix(row.Release720pAt, 0),
		Release1080pAt: time.Unix(row.Release1080pAt, 0),
		PremiereAt:     time.Unix(row.PremiereAt, 0),
		CreatedAt:      time.Unix(row.CreatedAt, 0),
		Year:           row.Year,
		CountSeasons:   row.CountSeasons,
		KpInternalId:   row.KpInternalId,
		ImdbInternalId: row.ImdbInternalId,
		Url:            row.Url,
		KpImageUrl:     row.KpImageUrl,
	}, nil
}
