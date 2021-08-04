package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/qianxi/blog-backend/model"
	rdx "github.com/qianxi/blog-backend/redis"
)

func TestSet(t *testing.T) {
	t.Run("test_set", func(t *testing.T) {
		rdb := rdx.New()

		// r, _ := json.Marshal(model.Comment{
		// 	ID:    1,
		// 	Login: "21213",
		// })
		rdb.Set(context.TODO(), "a", &model.Comment{
			ID: 1,
		}, time.Hour)
	})
}
