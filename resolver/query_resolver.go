package resolver

import (
	"context"
	"go-gql/config"
	"go-gql/model"
)

func (*QueryResolver) Hello() *string {
	s := "Hello World"
	return &s
}

// GetUserByID 对应Schema中定义的FindByID方法，如果方法的error不为空，将响应500错误码
func (*QueryResolver) GetUserByID(ctx context.Context, params UserParams) (*UserResolver, error) {
	db := config.GetDB()
	ms, e := model.GetUserByID(ctx, db, uint(*params.Id))
	return wrapUserResolver(ms), e
}

// GetUsers 对应Schema中定义的GetUsers方法，如果方法的error不为空，将响应500错误码
func (*QueryResolver) GetUsers(ctx context.Context, params UserParams) ([]*UserResolver, error) {
	db := config.GetDB()
	ids := make([]uint, 0)
	for _, id := range params.Ids {
		ids = append(ids, uint(*id))
	}
	ms, e := model.GetUsers(ctx, db, ids)

	return wrapUserResolvers(ms), e
}
