package chat

import (
	. "chat/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func create_connection() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect)

	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func migrate(db *gorm.DB) {
	// Migrate the schema
	// テーブルが既存でない場合のみ実行される
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Message{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").AddForeignKey("room_id", "rooms(id)", "RESTRICT", "RESTRICT")
}
