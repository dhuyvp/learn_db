package models

import "time"

type Person struct {
	FirstName string    `db:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string    `db:"last_name,omitempty" json:"last_name,omitempty"`
	Age       int       `db:"age,omitempty" json:"age,omitempty"`
	CreateAt  time.Time `db:"create_at,omitempty" json:"create_at,omitempty"`
}
