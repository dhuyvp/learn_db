package models

type Person struct {
	FirstName *string `db:"FirstName" json:"first_name"`
	LastName  *string `db:"LastName" json:"last_name"`
	Age       *int    `db:"Age" json:"age"`
	CreateAt  string  `db:"create_at" json:"create_at"`
	Id        *int    `db:"id" json:"id"`
}
