package dao

import (
	"connection"
	"entity"
	"log"
)

// GetUserByID 根据id，查询用户
func GetUserByID(id int32) (*entity.User, error) {
	user := new(entity.User)
	db, error := connection.GetDB()
	if error != nil {
		log.Fatal("can't get db connection!")
		return user, error
	}
	defer db.Close()
	rows, error := db.Query("select * from user where id = ?;", id)
	if error != nil {
		log.Fatal("query user error:" + error.Error())
		return user, error
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&user.ID, &user.Name)
	}
	return user, nil
}
