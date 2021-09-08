package dao

import (
	"testing"

	"shyiran/my-gin-vue/model"
)

func TestAddPost(t *testing.T) {
	post := &model.Post{}
	AddPost(post)
}
