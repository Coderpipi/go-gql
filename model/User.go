package model

import (
	"context"
	"errors"
	"go-gql/common"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PasswordCost = 12
	Unknown      = iota
	Normal
	Vip
	Admin
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:username;type:varchar(255);not null;default ''" json:"UserName"`
	Phone    string `gorm:"column:phone;type:varchar(11);not null;default ''" json:"phone"`
	Password string `gorm:"column:password;type:varchar(255);not null;" json:"password"`
	Type     int8   `gorm:"column:type;type:tinyint;not null; default 0" json:"type"`
	Age      int8   `gorm:"column:age;type:int(8);" json:"age"`
	Sex      string `gorm:"column:sex;type:varchar(10);not null;default ''" json:"sex"`
}

func (u *User) TableName() string {
	return "user"
}

func WithUserID(id uint) common.Option {
	return common.OptionFunc(func(options *common.Options) {
		options.Query["id"] = id
	})
}

func WithPhone(phone string) common.Option {
	return common.OptionFunc(func(options *common.Options) {
		options.Query["phone"] = phone
	})
}

func GetUser(ctx context.Context, db *gorm.DB, options ...common.Option) (*User, error) {
	opts := &common.Options{Query: make(map[string]interface{})}
	for _, opt := range options {
		opt.Apply(opts)
	}
	u := new(User)
	if err := db.WithContext(ctx).Model(u).Where(opts.Query).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByID(ctx context.Context, db *gorm.DB, userID uint) (*User, error) {
	u := new(User)
	u.ID = userID
	if err := db.WithContext(ctx).Model(u).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUsers(ctx context.Context, db *gorm.DB, ids []uint) ([]*User, error) {
	users := make([]*User, 0)
	if err := db.WithContext(ctx).Model(new(User)).Find(&users, ids).Error; err != nil {
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
	if err := db.WithContext(ctx).Model(user).Where("phone = ?", user.Phone).First(new(User)).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		if err := db.WithContext(ctx).Model(user).Save(user).Error; err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, errors.New("该手机号已经被用户使用了")
}

func UpdateUser(ctx context.Context, db *gorm.DB, user *User) (*User, error) {
	if user == nil {
		return nil, errors.New("user 信息不能为空")
	}
	var count int64 = 0
	if err := db.WithContext(ctx).Model(user).Where("id = ?", user.ID).Count(&count).Error; err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("当前用户不存在")
	}
	if err := db.WithContext(ctx).Model(user).UpdateColumns(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(ctx context.Context, db *gorm.DB, id uint) (*User, error) {
	user := &User{}
	user.ID = id
	if err := db.WithContext(ctx).Model(user).Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
