package model

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

type Sickness struct {

	//ID
	Id int64 `json:"id"`

	//UUID
	Uuid string `json:"uuid"`

	//名称
	Name string `json:"name"`

	//症状
	Symptom string `json:"symptom"`

	//护理
	Nursing string `json:"nursing"`

	//常用药
	Medicals string `json:"medicals"`

	//病程
	Duration string `json:"duration"`

	//风险
	Risks string `json:"risks"`

	CreateAt time.Time `json:"create_at"`
}

var SicknessTableName = "sickness"

func FetchSickById(id int64) (sick Sickness, err error) {
	sick = Sickness{}
	row := Db.QueryRow("SELECT "+SickItems("")+" from "+SicknessTableName+" where id = $1", id)
	ScanRow(row, &sick)
	return Sickness{}, nil
}

func FetchBySymptom(symptom string) (sicks []Sickness, err error) {

	rows, err := Db.Query("SELECT "+SickItems("")+" FROM "+SicknessTableName+" where symptom like %$1%", symptom)
	if err != nil {
		log.Fatalln(err)
	}
	ScanRows(rows, sicks)
	return
}

func (sick *Sickness) Save() (err error) {
	stt := "insert into " + SicknessTableName + " (" + SickToSaveItems() + ") values ($1, $2, $3, $4, $5, $6) returning id, createAt"
	stmt, err := Db.Prepare(stt)
	if stmt == nil {
		log.Fatalln("statement is nil in save sick, " + sick.Name)
		return
	}
	row := stmt.QueryRow(sick.Name, sick.Symptom, sick.Nursing, sick.Medicals, sick.Duration, sick.Risks)
	err = row.Scan(sick.Id, sick.CreateAt)
	return
}

func SickToSaveItems() string {
	return "`name`, symptom, nursing, medicals, duration, risks"
}

func SickItems(tag string) string {
	if len(tag) > 0 {
		return strings.Join([]string{
			tag + ".id",
			tag + ".uuid",
			tag + ".`name`",
			tag + ".symptom",
			tag + ".nursing",
			tag + ".medicals",
			tag + ".duration",
			tag + ".risks",
			tag + ".createAt",
		}, " ")
	}
	return "id, uuid,`name`, symptom, nursing, medicals, duration, risks, createAt"
}

func ScanRows(rows *sql.Rows, sicks []Sickness) {
	for rows.Next() {
		sick := Sickness{}
		err := rows.Scan(sick.Id, sick.Uuid, sick.Name, sick.Symptom, sick.Nursing, sick.Medicals, sick.Medicals, sick.Duration, sick.Risks)
		if err != nil {
			log.Fatalln("scan row error", err)
		}
		sicks = append(sicks, sick)
	}
	err := rows.Close()
	if err != nil {
		log.Fatalln("sql statement close error, table : "+SicknessTableName, err)
	}
}

func ScanRow(row *sql.Row, sick *Sickness) {
	err := row.Scan(sick.Id, sick.Uuid, sick.Name, sick.Symptom, sick.Nursing, sick.Medicals, sick.Medicals, sick.Duration, sick.Risks)
	if err != nil {
		log.Fatalln("scan row error", err)
	}
}
