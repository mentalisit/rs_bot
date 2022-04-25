package dbase

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "rsbot"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	//defer database.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	//log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	//defer database.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 7*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	//log.Printf("Connected to DB %s successfully\n", dbname)
	if no == 1 {
		err = createTableConfig(db)
		err = createTableNumkz(db)
		err = createTableRsevent(db)
		err = createTableSborkz(db)
		err = createTableSubscribe(db)
		err = createTableTimer(db)
		fmt.Println("Таблицы созданы ошибок вроде нет ")
		fmt.Println(err)
	}

	return db, nil
}

func init() {
	db, err := DbConnection()
	if err != nil {
		log.Printf("Error %s when getting database connection", err)
		return
	}
	defer db.Close()
	//log.Printf("Successfully connected to database")
	//err = createTableConfig(db)
	//err =createTableNumkz(db)
	//err =createTableRsevent(db)
	//err =createTableSborkz(db)
	//err =createTableSubscribe(db)
	//err =createTableTimer(db)
	if err != nil {
		log.Printf("Create product table failed with error %s", err)
		return
	}
}
