package repository

import (
	"database/sql"
	"time"

	userModel "../../models/user"
	_ "github.com/go-sql-driver/mysql"
)

// Create :: 新增會員
func Create(account string, password string) bool {
	db := db()
	result, err := db.Exec(
		"INSERT INTO user (Account, Password) VALUES (?, ?)",
		account,
		password,
	)

	count, err := result.RowsAffected()
	defer db.Close()
	return err == nil && count > 0
}

// Get :: 取得會員資料
func Get(account string) *userModel.UserDao {
	db := db()
	row := db.QueryRow("SELECT * FROM user WHERE Account = ?", account)
	var user userModel.UserDao
	err := row.Scan(&user.Account, &user.Password)
	defer db.Close()
	if err != nil {
		return nil
	}

	return &user
}

// GetList :: 取得會員列表
func GetList() []userModel.UserDao {
	db := db()
	rows, err := db.Query("SELECT * FROM user")
	checkError(err)

	var users []userModel.UserDao
	for rows.Next() {
		var user userModel.UserDao
		err = rows.Scan(&user.Account, &user.Password)
		checkError(err)
		users = append(users, user)
	}
	rows.Close()
	defer db.Close()
	return users
}

// Delete :: 刪除會員
func Delete(account string) bool {
	db := db()
	result, err := db.Exec("DELETE FROM user WHERE Account = ?", account)
	checkError(err)

	count, err := result.RowsAffected()
	defer db.Close()
	return err == nil && count > 0
}

// ModifyPassword :: 刪除會員
func ModifyPassword(account string, password string) bool {
	db := db()
	result, err := db.Exec("UPDATE user SET Password = ? WHERE Account = ?", password, account)
	checkError(err)

	count, err := result.RowsAffected()
	defer db.Close()
	return err == nil && count > 0
}

// db :: DB 連線
func db() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/golang?charset=utf8")
	checkError(err)

	err = db.Ping()
	checkError(err)

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

// checkError :: 檢查例外
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
