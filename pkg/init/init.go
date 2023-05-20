package init

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func ConnectUsers() *sql.DB {
// 	db, err := sql.Open("mysql", "root:123456@(localhost)/users")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// See "Important settings" section.
// 	db.SetConnMaxLifetime(time.Minute * 3)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(10)

// 	return db
// }

func InitDb() (*gorm.DB, error) {
	dsn := "root:123456@tcp(localhost)/users?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
