package repository

import "time"

type Task struct {
	ID         uint `storm:"id,increment"`
	Name       string
	Status     bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}
