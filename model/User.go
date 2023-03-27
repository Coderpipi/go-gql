package model

import (
	"context"
	"errors"
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

func GetUsers(ctx context.Context, db *gorm.DB, ids []uint) ([]*User, error) {
	users := make([]*User, 0)
	if err := db.Model(new(User)).Find(&users, ids).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(ctx context.Context, db *gorm.DB, user *User) (*User, error) {
	if user == nil {
		return nil, errors.New("user 信息不能为空")
	}

	if user.UserName == "" {
		return nil, errors.New("username 不能为空")

	}

	if err := db.Model(user).Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
