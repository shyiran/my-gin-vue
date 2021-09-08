package dao

import (
	"testing"
	"time"

	"shyiran/my-gin-vue/model"
)

func TestAddCategory(t *testing.T) {
	c := &model.Category{
		// ID:       1,
		Name:     "Mathematics",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	AddCategory(c)
}

func TestUpdateCategory(t *testing.T) {

}
