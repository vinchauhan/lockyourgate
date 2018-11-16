package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "vchauhan"
	password = ""
	dbname   = "lockyourgate_dev"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	//fmt.Println(psqlInfo)
	dbconn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	row := dbconn.Exec(`
	INSERT INTO orders(user_id, amount, description)
	WHERE id=$1`, 3)

	var id int
	var fname, lname, email string
	err = row.Scan(&id, &fname, &lname, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No rows in the resultset")
		} else {
			panic(err)
		}

	}

	log.Fatalf("The value retrieved are %s, %s, %s", fname, lname, email)

}
