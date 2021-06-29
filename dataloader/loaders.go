package dataloader

import (
	"context"
	"gqlgen-dataloader-prisma/db"
	"net/http"
	"time"
)

const loadersKey = "dataloaders"

type Loaders struct {
	UserByUsername        UserLoader
	PostsByAuthorUsername PostLoader
}

func Middleware(prismaClient *db.PrismaClient, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserLoader{
				fetch: func(usernames []string) ([]db.UserModel, []error) {
					//Fetch all the users by username. They do not follow the order of the "usernames" slice
					usersInDB, err := prismaClient.User.FindMany(db.User.Username.In(usernames)).Exec(r.Context())

					//Sort the username to be in the same order as the "usernames" slice
					userByUsername := map[string]db.UserModel{}
					for _, user := range usersInDB {
						userByUsername[user.Username] = user
					}
					users := make([]db.UserModel, len(usernames))
					for i, username := range usernames {
						users[i] = userByUsername[username]
					}
					//Return the sorted result
					return users, []error{err}
				},
				wait:     1 * time.Millisecond,
				maxBatch: 100,
			},
			PostLoader{
				fetch: func(authorUsernames []string) ([][]db.PostModel, []error) {
					users, err := prismaClient.User.FindMany(db.User.Username.In(authorUsernames)).With(db.User.Posts.Fetch()).Exec(r.Context())
					postsByAuthorUsername := map[string][]db.PostModel{}
					for _, user := range users {
						postsByAuthorUsername[user.Username] = user.Posts()
					}
					posts := make([][]db.PostModel, len(authorUsernames))
					for i, username := range authorUsernames {
						posts[i] = postsByAuthorUsername[username]
					}
					return posts, []error{err}
				},
				wait:     1 * time.Millisecond,
				maxBatch: 100,
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
