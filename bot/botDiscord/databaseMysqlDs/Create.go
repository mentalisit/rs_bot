package databaseMysqlDs

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func createTableSborkz(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS sborkz(
		id int primary key auto_increment,
		name VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
    	nameid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
    	mention VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		guildid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
    	lvlkz VARCHAR(20) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		chatid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		mesid VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		time VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		date VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		numberkz INT(11) NULL DEFAULT NULL,
		numberevent INT(11) NULL DEFAULT NULL,
		eventpoints INT(31) NULL DEFAULT NULL,
		timedown INT(11) UNSIGNED NULL DEFAULT NULL,
		active INT(11) NULL DEFAULT NULL
		)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}

	return nil
}

func createTableNumkz(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS numkz(
		id int primary key auto_increment,
    	lvlkz VARCHAR(20) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		number INT(11) NOT NULL DEFAULT '0',
		chatid VARCHAR(20) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci'
		)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}

func createTableRsevent(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS rsevent(
		id int primary key auto_increment,
		chatid VARCHAR(20) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		numevent INT(11) NULL DEFAULT NULL,
		activeevent INT(11) NULL DEFAULT NULL,
		number INT(11) NULL DEFAULT NULL
	)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}

	return nil
}

func CreateTableTempTop(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS temptop(
		id int primary key auto_increment,
		name VARCHAR(30) NULL DEFAULT NULL COLLATE 'utf8mb4_unicode_ci',
		numkz INT(11) NULL DEFAULT NULL
	)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}

	return nil
}

func createTableTEST(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS product(
		product_id int primary key auto_increment, 
		product_name text,
		product_price int,
		created_at datetime default CURRENT_TIMESTAMP, 
		updated_at datetime default CURRENT_TIMESTAMP
		)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS product(
		product_id int primary key auto_increment, 
		product_name text,
		product_price int,
		created_at datetime default CURRENT_TIMESTAMP, 
		updated_at datetime default CURRENT_TIMESTAMP
		)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Ошибка создания таблицы %s ", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("ошибка чтения строк  %s ", err)
		return err
	}
	if rows != 0 {
		log.Printf("что-то пошло не так: %d", rows)
	}
	return nil
}
