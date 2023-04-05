package resolver

import (
	"context"
	"go-gql/config"
	"go-gql/model"
	"gorm.io/gorm"
)

func (*MutationResolver) CreateUser(ctx context.Context, params *UserInput) (*UserResolver, error) {
	db := config.GetDB()
	user, e := model.CreateUser(ctx, db, &model.User{UserName: params.Input.Username, Sex: params.Input.Sex, Age: int8(params.Input.Age), Phone: params.Input.Phone, Type: int8(params.Input.Type)})
	return wrapUserResolver(user), e
}

func (*MutationResolver) UpdateUser(ctx context.Context, params *UserInput) (*UserResolver, error) {
	db := config.GetDB()
	user, e := model.UpdateUser(ctx, db, &model.User{Model: gorm.Model{ID: uint(params.Id)}, UserName: params.Input.Username, Sex: params.Input.Sex, Age: int8(params.Input.Age), Phone: params.Input.Phone, Type: int8(params.Input.Type)})
	return wrapUserResolver(user), e
}

func (*MutationResolver) DeleteUser(ctx context.Context, params *DeleteInput) (*UserResolver, error) {
	db := config.GetDB()
	user, e := model.DeleteUser(ctx, db, uint(params.Id))
	return wrapUserResolver(user), e
}
