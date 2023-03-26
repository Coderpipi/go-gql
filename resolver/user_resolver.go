package resolver

import (
	"go-gql/model"
	"time"
)

type UserResolver struct {
	*model.User
}

type UserParams struct {
	Id  *int32  `json:"id"`
	Sex *string `json:"sex"`
	Age *string `json:"age"`
}

func (r *UserResolver) ID() int32 {
	return int32(r.User.ID)
}

func (r *UserResolver) Username() string {
	return r.User.UserName
}

func (r *UserResolver) Age() int32 {
	return int32(r.User.Age)
}

func (r *UserResolver) Sex() string {
	return r.User.Sex
}

func (r *UserResolver) CreatedAt() time.Time {
	return r.User.CreatedAt
}

func (r *UserResolver) UpdatedAt() time.Time {
	return r.User.UpdatedAt
}

func wrapUserResolver(user *model.User) *UserResolver {
	if user == nil {
		return nil
	}
	return &UserResolver{
		user,
	}
}
