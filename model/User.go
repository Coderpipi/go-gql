package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:username;type:varchar(255);not null;default ''" json:"UserName"`
	Age      int8   `gorm:"column:age;type:int(8);" json:"age"`
	Sex      string `gorm:"column:sex;type:varchar(10);not null; default ''" json:"sex"`
	Password string `gorm:"column:sex;type:varchar(255);not null;" json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func GetUserByID(ctx context.Context, db *gorm.DB, userID uint) (*User, error) {
	u := new(User)
	u.ID = userID
	if err := db.Model(u).First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
