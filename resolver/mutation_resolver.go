package resolver

import (
	"context"
	"go-gql/config"
	"go-gql/model"
)

func (*MutationResolver) CreateUser(ctx context.Context, params *UserInput) (*UserResolver, error) {
	db := config.GetDB()
	user, e := model.CreateUser(ctx, db, &model.User{UserName: params.Input.Username, Sex: params.Input.Sex, Age: int8(params.Input.Age)})
	return wrapUserResolver(user), e
}
