package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen-dataloader-prisma/dataloader"
	"gqlgen-dataloader-prisma/db"
	"gqlgen-dataloader-prisma/graph/generated"
)

func (r *mutationResolver) CreateOnePost(ctx context.Context, authorUsername string, title string) (*db.PostModel, error) {
	return r.Prisma.Post.CreateOne(db.Post.Title.Set(title), db.Post.Author.Link(db.User.Username.Equals(authorUsername))).Exec(ctx)
}

func (r *mutationResolver) CreateOneUser(ctx context.Context, username string) (*db.UserModel, error) {
	return r.Prisma.User.CreateOne(db.User.Username.Set(username)).Exec(ctx)
}

func (r *postResolver) Author(ctx context.Context, obj *db.PostModel) (*db.UserModel, error) {
	user, err := dataloader.For(ctx).UserByUsername.Load(obj.AuthorUsername)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *queryResolver) Users(ctx context.Context) ([]db.UserModel, error) {
	return r.Prisma.User.FindMany().Exec(ctx)
}

func (r *queryResolver) Posts(ctx context.Context) ([]db.PostModel, error) {
	return r.Prisma.Post.FindMany().Exec(ctx)
}

func (r *userResolver) Posts(ctx context.Context, obj *db.UserModel) ([]db.PostModel, error) {
	return dataloader.For(ctx).PostsByAuthorUsername.Load(obj.Username)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
