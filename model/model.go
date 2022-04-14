package model

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID        uint64    `gorm:"primaryKey;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Redirect  string    `json:"redirect" gorm:"not null"`
	Goly      string    `json:"goly" gorm:"unique; not null"`
	Clicked   uint64    `json:"clicked"`
	Random    bool      `json:"random"`
}

func Setup() {
	dsn := "host=localhost user=admin password=test port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}

}
