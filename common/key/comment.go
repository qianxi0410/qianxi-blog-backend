package key

import "fmt"

func CommentCount() string {
	return "blog_comment_count"
}

func Comments(page, size int64) string {
	return fmt.Sprintf("comment:%d_%d_page_size", page, size)
}
