package database

import (
	"encoding/json"
	"fmt"
	"learn_db/app/models"
	"learn_db/pkg/utils"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var columnDB = map[string][]*models.Columns{}

type Variable struct {
	Db *sqlx.DB
	Tx *sqlx.Tx
}

func New() (*Variable, error) {

	db, err := MySQLConnection()

	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected!")

	return &Variable{
		Db: db,
		Tx: nil,
	}, nil
}

func MySQLConnection() (*sqlx.DB, error) {
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("mysql", mysqlConnURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}
	return db, nil
}

func (e *Variable) InsertObject(tableName string, object interface{}, isUseTx bool) (int, error) {
	var data map[string]interface{}
	lastID := 0
	jsonString, err := json.Marshal(object)
	if err != nil {
		return 0, err
	}

	json.Unmarshal(jsonString, &data)
	var listColumn = make([]*models.Columns, 0)
	pointMap := map[string]interface{}{}

	if columnDB[tableName] == nil {
		queryColumn := "SELECT COLUMN_NAME as FieldName,DATA_TYPE as DataType, " +
			" CHARACTER_MAXIMUM_LENGTH as MaxLength,CASE WHEN COLUMN_KEY='PRI' THEN 1 ELSE 0 END AS IsIdentity, " +
			" CASE WHEN IS_NULLABLE='YES' THEN 1 ELSE 0 END as IsNullable,Extra FROM Information_schema.Columns " +
			" WHERE Table_Name= " + tableName + " AND Table_Schema=database() ORDER BY Column_Name"

		queryColumn = "SELECT * FROM " + tableName

		//queryColumn = "SELECT COLUMN_NAME as FieldName,DATA_TYPE as DataType,  CHARACTER_MAXIMUM_LENGTH as MaxLength,CASE WHEN COLUMN_KEY='PRI' THEN 1 ELSE 0 END AS IsIdentity,  CASE WHEN IS_NULLABLE='YES' THEN 1 ELSE 0 END as IsNullable,Extra FROM Information_schema.Columns  WHERE Table_Name= \"Person\" AND Table_Schema=database() ORDER BY Column_Name"
		e.Db.Select(&listColumn, queryColumn)

		if err != nil {
			return 0, err
		}
		columnDB[tableName] = listColumn
	} else {
		listColumn = columnDB[tableName]
	}

	fmt.Println(listColumn)

	builderColumn := strings.Builder{}
	builderValue := strings.Builder{}
	for i := 0; i < len(listColumn); i++ {
		if listColumn[i].IsIdentity == "1" && listColumn[i].Extra == "auto_increnent" {
			continue
		}
		var fieldName = listColumn[i].FieldName
		if data[fieldName] == nil {
			continue
		}

		log.Fatal(fmt.Println(fieldName))

		builderColumn.WriteString("," + fieldName)
		builderValue.WriteString(",:" + fieldName)
		processDataBeforeSave(data, listColumn[i])
		pointMap[fieldName] = data[fieldName]
	}
	var stringCol = builderColumn.String()
	var stringValue = builderValue.String()
	var query = "INSERT INTO " + tableName + " (" + stringCol[1:] + ") VALUES (" + stringValue[1:] + ") RETURNING id;"

	stmt, err := e.Db.PrepareNamed(query)

	if isUseTx {
		stmt, err = e.Tx.PrepareNamed(query)
	}
	if err != nil {
		return 0, err
	}

	err = stmt.Get(&lastID, pointMap)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func processDataBeforeSave(data map[string]interface{}, column *models.Columns) {
	if column.DataType == "varchar" {
		value := data[column.FieldName].(string)
		data[column.FieldName] = value

		if len(value) > column.MaxLength {
			data[column.FieldName] = value[0:column.MaxLength]
		}
	}
}
