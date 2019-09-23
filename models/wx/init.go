package wx

import "github.com/jinzhu/gorm"

var Db *gorm.DB

func ModelInit(db *gorm.DB) {
	Db = db
}
