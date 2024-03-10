package model

import (
	"OCsemmerApp/pkg/db"
	"OCsemmerApp/pkg/domain"
	"database/sql"
	"fmt"
)

// ユーザーが入力した製品名ですべてのユーザーの履歴が検索できてしまうやつ
func SearchData(name string) (*domain.DataResponse, error) {
	rows, err := db.Conn.Query(
		fmt.Sprintf(
			`SELECT
				T1.product_id,
				T1.product_name,
				T2.name,
				T1.purchase_date
			FROM 
				purchase AS T1
				JOIN users AS T2
					ON T1.user_id = T2.id
			WHERE
				T2.name = 'guest'
				AND T1.product_name LIKE '%%%s%%'`, name))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var datalist domain.DataResponse
	for rows.Next() {
		data, err := convertToData(rows)
		if err != nil {
			return nil, err
		}
		datalist = append(datalist, data)
	}
	return &datalist, nil
}

func convertToData(rows *sql.Rows) (*domain.Data, error) {
	data := domain.Data{}
	if err := rows.Scan(
		&data.ProductID,
		&data.ProductName,
		&data.UserName,
		&data.PurchaseDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}
