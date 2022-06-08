package main

import (
	"database/sql"
	"log"
)

func getOne(tNo int) (a teacher, err error) {
	a = teacher{}
	//scan将数据拷贝到参数变量
	err = db.QueryRow("SELECT tNo,username,password,role dbo FROM teacher WHERE tNo=?", tNo).
		Scan(&a.tNo, &a.username, &a.password, &a.role)
	return
}

func getMany(tNo int) (teachers []teacher, err error) {
	rows, err := db.Query("SELECT tNo,username,password,role dbo FROM teacher WHERE tNo>?", tNo)
	for rows.Next() {
		a := teacher{}
		err = rows.Scan(&a.tNo, &a.username, &a.password, &a.role)
		if err != nil {
			log.Fatalln(err.Error())
		}
		teachers = append(teachers, a)
	}
	return
}

func (t *teacher) Update() (err error) {
	_, err = db.Exec("UPDATE teacher SET tNo=?,username=?,password=?,role=?", t.tNo, t.username, t.password, t.role)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}

//添加有问题
func (t *teacher) Insert() (err error) {
	statement := `INSERT INTO teacher
(username,
password,
role)
VALUES
(@username,
@password,
@role);
SELECT isNull(SCOPE_IDENTITY(),-1);`
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("username", t.username),
		sql.Named("password", t.password),
		sql.Named("role", t.role)).Scan(&t.tNo)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}
