package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "12.34.56.78"
	port     = 5432
	user     = "postgres"
	password = "*****"
	dbname   = "registry"
)

func connectDB() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}


func insertUser(db *sql.DB)  {
	stmt,err := db.Prepare("insert into public.harbor_user(user_id,username,email,password,realname,comment,deleted,reset_uuid,salt,sysadmin_flag,creation_time,update_time,password_version) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)")
	if err != nil {
		log.Fatal(err)
	}
	_,err = stmt.Exec("100","000000", "yb", "12345678AbC", "xx", "from LDAP", "f", "NULL", "NULL", "f", "2021-07-16 14:57:52.460499", "2021-07-16 14:57:52.460499", "sha256")
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("insert into user_tbl success")
	}
}



func query(db *sql.DB){
	var user_id,username,email string

	rows,err:=db.Query(" select user_id,username,email from harbor_user where user_id=$1","1")

	if err!= nil{
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next(){
		err:= rows.Scan(&user_id,&username,&email)

		if err!= nil{
			fmt.Println(err)
		}
	}

	err = rows.Err()
	if err!= nil{
		fmt.Println(err)
	}

	fmt.Println(user_id,username,email)
}

func main()  {
	db:=connectDB()
	//query(db)
	insertUser(db)

}

