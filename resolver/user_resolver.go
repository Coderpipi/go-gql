package resolver

import (
	"go-gql/model"
	"time"
)

type UserResolver struct {
	*model.User
}
type UserInput struct {
	Input struct {
		Username string `json:"username"`
		Sex      string `json:"sex"`
		Age      int32  `json:"age"`
	} `json:"input"`
}
type DeleteInput struct {
	Id int32 `json:"id"`
}
type UserParams struct {
	Id       *int32   `json:"id"`
	Username *string  `json:"username"`
	Ids      []*int32 `json:"ids"`
	Sex      *string  `json:"sex"`
	Age      *int32   `json:"age"`
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

func wrapUserResolvers(ms []*model.User) []*UserResolver {
	users := make([]*UserResolver, 0, len(ms))
	for _, user := range ms {
		users = append(users, &UserResolver{user})
	}

	return users
}
