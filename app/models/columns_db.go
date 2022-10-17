package models

type Columns struct {
	FieldName  string `db:"fieldname" json:"field_name"`
	DataType   string `db:"datatype" json:"data_type"`
	MaxLength  int    `db:"maxlength" json:"max_length"`
	IsIdentity string `db:"isidentity" json:"is_identity"`
	IsNullable string `db:"isnullable" json:"is_nullable"`
	Extra      string `db:"extra" json:"extra"`
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
