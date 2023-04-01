package resolver

import (
	"context"
	"errors"
	"go-gql/config"
	"go-gql/model"
	"gorm.io/gorm"
)

const (
	PasswordAuthorization = iota + 1
	VerifyCodeAuthorization
)

type IVerify interface {
	Verify(ctx context.Context, params *LoginInput) (*model.User, error)
}

type PasswordLogin struct {
}

type VerifyCodeLogin struct {
}

func (*PasswordLogin) Verify(ctx context.Context, params *LoginInput) (*model.User, error) {
	if params.Input.Phone == "" || params.Input.Password == "" {
		return nil, errors.New("手机号或密码不能为空")
	}
	db := config.GetDB()
	// todo 通过PhoneNumber查询用户信息, 没有记录则证明用户不存在
	user, err := model.GetUser(ctx, db, model.WithPhone(params.Input.Phone))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("改手机号尚未注册用户")
		}
		return nil, err
	}
	// todo 拿当前的密码加密后进行比对
	if user.CheckPassword(params.Input.Password) {
		return user, nil
	}
	return nil, errors.New("登录失败, 密码错误")
}

func (*VerifyCodeLogin) Verify(ctx context.Context, params *LoginInput) (*model.User, error) {
	if params.Input.Phone == "" || params.Input.Password == "" {
		return nil, errors.New("用户名不存在或者密码错误")
	}

	// todo 通过PhoneNumber查询用户信息, 没有记录则证明用户不存在

	// todo 拿当前的密码加密后进行比对

	return nil, errors.New("暂时不支持该验证方式")
}

type VerifyBuilder struct {
	verifyType int32
}

func NewVerifyBuilder(verifyType int32) *VerifyBuilder {
	return &VerifyBuilder{
		verifyType: verifyType,
	}
}
func (v *VerifyBuilder) Builder() IVerify {
	switch v.verifyType {
	case PasswordAuthorization:
		return &PasswordLogin{}
	case VerifyCodeAuthorization:
		return &VerifyCodeLogin{}
	default:
		panic("no such login type support")
	}
}
