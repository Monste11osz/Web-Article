package mysql

import (
	"WebPrac/pkg/models"
	"database/sql"
	"errors"
)

type WebPracMod struct {
	DB *sql.DB
}

func (s *WebPracMod) Insert(title, content, expires string) (int, error) {
	strt := `INSERT INTO WebPrac (title, content, created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL  ? DAY))`
	res, err := s.DB.Exec(strt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (s *WebPracMod) Get(id int) (*models.Article, error) {
	strt := `SELECT id, title, content, created, expires FROM WebPrac
	WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := s.DB.QueryRow(strt, id)
	a := &models.Article{}
	err := row.Scan(&a.ID, &a.Title, &a.Content, &a.Created, &a.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNotReport
		} else {
			return nil, err
		}
	}
	return a, nil
}

func (s *WebPracMod) LastUs() ([]*models.Article, error) {
	strt := `SELECT id, title, content, created, expires FROM WebPrac
			WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := s.DB.Query(strt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var artic []*models.Article
	for rows.Next() {
		a := &models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.Created, &a.Expires)
		if err != nil {
			return nil, err
		}
		artic = append(artic, a)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return artic, nil
}
