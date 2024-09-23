package model

import (
	"gorm.io/gorm"
)

/*
Timehash table structure:
Uid - users hash
email - users email
hash - 64 hash
param - Which function create this record. We need confirm email in signup and change current password
created - Where does record was created
livetime - hash life time
*/

type TimeHash struct {
	gorm.Model
	UID      string `gorm:"column:uid;type:varchar(60);NOT NULL;" json:"uid"`
	Mail     string `gorm:"column:email;type:varchar(254);DEFAULT '';" json:"email"`
	Hash     string `gorm:"column:hash;type:varchar(254);DEFAULT '';" json:"hash"`
	Param    string `gorm:"column:param;type:varchar(254);DEFAULT '';" json:"parametr"`
	Created  int    `gorm:"column:created;type:int;DEFAULT '0'" json:"created"`
	Lifetime int    `gorm:"column:lifetime;type:int;DEFAULT '0'" json:"lifetime"`
}
