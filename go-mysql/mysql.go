package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("yang:x@tcp(localhost:3306)/test1")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("Error connecting database!")
		panic(err)
	}

	// 最大空闲连接数，默认不配置，是 2 个最大空闲连接
	db.SetMaxIdleConns(5)
	// 最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	// 空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)

	err = db.Ping()
	if err != nil {
		log.Println("Ping failed database!")
		_ = db.Close()
		if err != nil {
			return
		}
		panic(err)
	}
	DB = db
}

// insert values into the table
func insert() {
	rst, err := DB.Exec("insert into user (username, sex, email) values(?, ?, ?)",
		"Messi", "Male", "messi@gmail.com")
	if err != nil {
		log.Println("sql insert error!")
		panic(err)
	}
	id, err := rst.LastInsertId()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully insert id:", id)
}

type User struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

// query user from the table
func query(id int) (*User, error) {
	rows, err := DB.Query("select * from user where id=? limit 1", id)
	if err != nil {
		log.Println("sql query error:", err)
		return nil, errors.New(err.Error())
	}

	user := new(User)
	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.UserName, &user.Sex, &user.Email)
		if err != nil {
			log.Println("sql scan error:", err)
			return nil, errors.New(err.Error())
		}
	}
	return user, nil
}

// update user from the table
func update(id int, username string) {
	rst, err := DB.Exec("update user set username=? where id=?", username, id)
	if err != nil {
		log.Println("failed to update the user:", err)
		return
	}
	affected, _ := rst.RowsAffected()
	log.Println("Successfully updated the row:", affected)
}

// delete a user from the table
func delete(id int) {
	rst, err := DB.Exec("delete from user where id=?", id)
	if err != nil {
		log.Println("sql delete error:", err)
		return
	}
	affected, _ := rst.RowsAffected()
	log.Println("Successfully deleted the row:", affected)
}

func main() {
	defer DB.Close()

	//insert()

	//user, err := query(1)
	//if err != nil {
	//	log.Println("query error:", err)
	//	panic(err)
	//}
	//log.Println("Get a user:", *user) // {1 Messi Male messi@gmail.com}

	//update(1, "Ronaldo")

	delete(1)
}
