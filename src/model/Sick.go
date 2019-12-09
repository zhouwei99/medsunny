package model

import (
	"database/sql"
	"log"
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

var SicknessTableName = "public.data_sickness"

func FetchSickById(id int64) (sick *Sickness, err error) {
	sick = &Sickness{}
	row := Db.QueryRow("SELECT id, uuid, name, symptom, nursing, medicals, duration, risks, create_at from public.data_sickness where id = $1", id)
	ScanRow(row, sick)
	return sick, nil
}

func FetchBySymptom(symptom string) (sicks []Sickness, err error) {
	syp := "%" + symptom + "%"
	rows, err := Db.Query("SELECT id, uuid, name, symptom, nursing, medicals, duration, risks, create_at FROM public.data_sickness where symptom like $1", syp)
	if err != nil {
		log.Println(err)
	}
	ScanRows(rows, sicks)
	return
}

func (sick *Sickness) Save() (err error) {
	if sick.Id == 0 {
		err = sick.doCreate()
	} else {
		err = sick.doUpdate()
	}
	return
}

func (sick *Sickness) doCreate() (err error) {
	stt := "insert into public.data_sickness (uuid, name, symptom, nursing, medicals, duration, risks, create_at) values ($1, $2, $3, $4, $5, $6, $7) returning id"
	stmt, err := Db.Prepare(stt)
	if stmt == nil {
		log.Println("statement is nil in save sick, " + sick.Name)
		return
	}
	row := stmt.QueryRow(createUUID(), sick.Name, sick.Symptom, sick.Nursing, sick.Medicals, sick.Duration, sick.Risks, time.Time{})
	err = row.Scan(sick.Id, sick.CreateAt)
	return
}

func (sick *Sickness) doUpdate() (err error) {
	var sk *Sickness
	sk, err = FetchSickById(sick.Id)
	if err != nil {
		log.Printf("update sickness error, on reading sickness info by id: %d", sick.Id)
		log.Println(err)
		return
	}
	stt := "update public.data_sickness set name=$2, symptom=$3, nursing=$4, medicals=$5, duration=$6, risks=$7 where id = $1"
	stmt, err := Db.Prepare(stt)
	if stmt == nil {
		log.Println("statement is nil in save sick, " + sick.Name)
		return
	}
	_ = stmt.QueryRow(sk.Id, sick.Name, sick.Symptom, sick.Nursing, sick.Medicals, sick.Duration, sick.Risks)
	return
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
		log.Println("scan row error", err)
	}
}
