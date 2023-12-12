package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const logfile = "/etc/mhd/gosample/logs/replicate.log"
const connstring = "host=192.168.50.40 port=5432 user=pgremote password=pgremoteuser11 database=postgres sslmode=require"

var Logfile, LogfileErr = os.Create(logfile)
var logs = log.New(Logfile, "::: DB :::", log.Lshortfile)

func AddProfile(srcPort, destination, destPort string) error {
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		logs.Panicf("Error openening database :   %s", err)
		return err
	}

	defer db.Close()

	insertQuery := "INSERT INTO app.profile(src_port,destination,dest_port) VALUES(%s)"
	formatString := fmt.Sprintf("'%s','%s','%s'", srcPort, destination, srcPort)
	fullInsert := fmt.Sprintf(insertQuery, formatString)

	_, err = db.Exec(fullInsert)
	if err != nil {
		logs.Printf("Error inserting values :  %s", err)
		return err
	}
	return nil

}
