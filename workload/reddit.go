package workload

import (
	"context"
	"github.com/artikell/valkey-tpc/storage"
	"github.com/spf13/viper"
)

/*
Variable：
- User Count
- Post Count
Business:
- User Login
- Create Post
- View Post
- List Post
- Comment Post
*/
const workLoadName = "Reddit"

func init() {
	registerWorkLoad(workLoadName, &WorkLoad{
		paramsVerify: func(v *viper.Viper) error {
			return nil
		},
		action: []*Action{
			{name: "user_login", run: runUserLogin, weight: 200},
			{name: "create_post", run: runCreatePost, weight: 50},
			{name: "view_post", run: runViewPost, weight: 500},
			{name: "list_post", run: runListPost, weight: 200},
			{name: "comment_post", run: runCommentPost, weight: 50},
		},
	})
}

func runUserLogin(ctx context.Context, st *storage.Storage) error {
	//opt := getOption(ctx)
	//cli := getClient(ctx)

	return nil
}

func runCreatePost(ctx context.Context, st *storage.Storage) error {
	return nil
}

func runViewPost(ctx context.Context, st *storage.Storage) error {
	return nil
}

func runListPost(ctx context.Context, st *storage.Storage) error {
	return nil
}

func runCommentPost(ctx context.Context, st *storage.Storage) error {
	return nil
}
