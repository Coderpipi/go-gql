package resolver

import (
	"context"
	"fmt"
	"go-gql/config"
	"go-gql/model"
	"go-gql/utils"
)

func (*QueryResolver) Hello() string {
	return "Hello World"
}

// User 对应Schema中定义的FindByID方法，如果方法的error不为空，将响应500错误码
func (*QueryResolver) User(ctx context.Context, params UserQueryParams) (*UserResolver, error) {
	db := config.GetDB()
	ms, e := model.GetUserByID(ctx, db, uint(params.Id))
	return wrapUserResolver(ms), e
}

// Users 对应Schema中定义的GetUsers方法，如果方法的error不为空，将响应500错误码
func (*QueryResolver) Users(ctx context.Context, params UserQueryParams) ([]*UserResolver, error) {
	db := config.GetDB()
	ids := make([]uint, 0)
	for _, id := range params.Ids {
		ids = append(ids, uint(id))
	}
	ms, e := model.GetUsers(ctx, db, ids)

	return wrapUserResolvers(ms), e
}

// Login 登录
func (*QueryResolver) Login(ctx context.Context, params LoginInput) (*Token, error) {
	verify := NewVerifyBuilder(params.Input.LoginType)
	user, err := verify.Builder().Verify(ctx, &params)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	// todo 验证通过, 颁发token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, err

	}
	return &Token{
		TokenString: token,
	}, nil
}
