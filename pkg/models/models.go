package models

import (
	"errors"
	"time"
)

var ErrNotReport = errors.New("models: Такой статьи нет")

type Article struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
