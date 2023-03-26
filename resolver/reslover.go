package resolver

import (
	"context"
	"go-gql/config"
	"go-gql/model"
)

type QueryResolver struct {
}

// GetUserByID 对应Schema中定义的FindByID方法，如果方法的error不为空，将响应500错误码
func (_ *QueryResolver) GetUserByID(ctx context.Context, params UserParams) (*UserResolver, error) {
	db := config.GetDB()
	ms, e := model.GetUserByID(ctx, db, uint(*params.Id))
	return wrapUserResolver(ms), e
}

// FindAll 对应Schema中定义的FindByID方法，如果方法的error不为空，将响应500错误码
func (_ *QueryResolver) FindAll(ctx context.Context, params UserParams) ([]*UserResolver, error) {
	//db := config.GetDB()
	//ms, e := model.GetUserByID(ctx, db, uint(*params.Id))
	return nil, nil
}
