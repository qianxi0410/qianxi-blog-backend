package redis

import "fmt"

func PostCount() string {
	return "blog_post_count"
}

func TagCount(tag string) string {
	return fmt.Sprintf("%s_tag_count", tag)
}

func Post(id string) string {
	return fmt.Sprintf("%s_post", id)
}

func PostWithPageAndSize(page, size int) string {
	return fmt.Sprintf("%d_%d_page_size", page, size)
}

func PostWithPageAndSizeAndTag(page, size int, tag string) string {
	return fmt.Sprintf("%d_%d_page_size_%s", page, size, tag)
}
