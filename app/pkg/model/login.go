package model

import (
	"OCsemmerApp/pkg/db"
	"OCsemmerApp/pkg/domain"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// sidechannel

type Login struct {
	ID       string
	Name     string
	PassWord string
}

// 最初にuserのパスワードを設定
func UpdatePass(id, password string) error {
	_, err := db.Conn.Exec("UPDATE users SET password = $1 WHERE id = $2", password, id)
	if err != nil {
		return err
	}
	return nil
}

// 駄目なログインフォームのバックエンドコード
func SearchUser(request domain.LoginRequest) error {
	row := db.Conn.QueryRow("SELECT * FROM users WHERE name = $1", request.Name)
	login := Login{}
	// ユーザーがいない場合パスワード検証をせず返してしまう
	if err := row.Scan(&login.ID, &login.Name, &login.PassWord); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("ユーザーが存在しません")
		}
		log.Println(err)
		return nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(login.PassWord), []byte(request.Pass)); err != nil {
		log.Println(err)
		return fmt.Errorf("パスワードが違います")
	}
	return nil
}
