package key

import "fmt"

func PostsCount() string {
	return "blog_post_count"
}

func PostsCountWithTag(tag string) string {
	return fmt.Sprintf("%s_tag_count", tag)
}

func Post(id int64) string {
	return fmt.Sprintf("%d_post", id)
}

func Posts(page, size int64) string {
	return fmt.Sprintf("%d_%d_page_size", page, size)
}

func PostsWithTag(page, size int64, tag string) string {
	return fmt.Sprintf("%d_%d_page_size_%s", page, size, tag)
}
