package tests

import (
	"encoding/json"
	"fmt"
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/models"
	"gorm.io/gorm"
	"testing"
)

func TestDB(t *testing.T) {
	fmt.Println("----------------TestDB 开始----------------")

	assert := getAssert(t)
	password, err := passwd.Hash("123")
	assert.Nil(err)

	user := &models.TestUser{
		Username: util.Ptr(util.Nanoid(5)).(*string),
		Password: password,
		Email:    util.Ptr(util.Nanoid(5) + "@test.com").(*string),
	}

	err = db.Transaction(func(tx *gorm.DB) error {

		result := tx.Create(user)
		assert.Nil(result.Error)

		userInfo := &models.TestUserInfo{
			Nickname: util.Ptr("测试01").(*string),
			Avatar:   util.Ptr("/static/images/avatar.png").(*string),
			User:     user,
		}
		result = tx.Create(userInfo)
		assert.Nil(result.Error)

		categories := []models.TestCategory{
			{
				Name:     "分类1",
				ParentID: util.Ptr(uint(0)).(*uint),
				User:     user,
				Children: []models.TestCategory{
					{
						Name: "子分类1",
						User: user,
					},
				},
			},
			{
				Name:     "分类2",
				ParentID: util.Ptr(uint(0)).(*uint),
				User:     user,
				Children: []models.TestCategory{
					{
						Name: "子分类2",
						User: user,
					},
					{
						Name: "子分类3",
						User: user,
					},
				},
			},
		}

		result = tx.Create(categories)
		assert.Nil(result.Error)

		var category *models.TestCategory
		category = (func() *models.TestCategory {
			for _, lvl1 := range categories {
				for _, lvl2 := range lvl1.Children {
					return &lvl2
				}
			}
			panic("分类不存在")
		})()

		articles := []models.TestArticle{
			{
				Title:    "标题1",
				Content:  util.Ptr("内容1").(*string),
				Category: category,
				User:     user,
			},
			{
				Title:    "标题2",
				Content:  util.Ptr("内容2").(*string),
				Category: category,
				User:     user,
			},
		}
		result = tx.Create(articles)
		assert.Nil(result.Error)

		article := &articles[0]
		comments := []models.TestComment{
			{
				Content:  util.Ptr("评论1").(*string),
				User:     user,
				Article:  article,
				ParentID: util.Ptr(uint(0)).(*uint),
				Children: []models.TestComment{
					{
						Content: util.Ptr("评论2").(*string),
						User:    user,
						Article: article,
					},
					{
						Content: util.Ptr("评论3").(*string),
						User:    user,
						Article: article,
					},
				},
			},
		}
		result = tx.Create(comments)
		assert.Nil(result.Error)

		return nil
	})
	assert.Nil(err)

	id := user.ID
	user = new(models.TestUser)
	db.Preload("UserInfo").Preload("Categories", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Children").Preload("Articles", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Comments", func(db *gorm.DB) *gorm.DB {
				return db.Preload("User.UserInfo").Preload("Children.User.UserInfo")
			})
		})
	}).Preload("Articles", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User.UserInfo").Preload("Children.User.UserInfo")
		})
	}).Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Parent.User.UserInfo").Preload("Article", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User.UserInfo")
		})
	}).First(user, id)

	out, err := user.SetModel(user).Out(app.Rule{
		"_skips": app.Skips{"password", "deleted_at"},
		"user_info": app.Rule{
			"nickname": "Nickname",
			"avatar": struct {
				Args []string
				Func func(avatar string) string
			}{
				Args: []string{"avatar"},
				Func: func(avatar string) string {
					return config.App.Url + avatar
				},
			},
		},
		"categories": app.Rule{
			"name": "Name",
			"children": app.Rule{
				"name": "Name",
				"articles": app.Rule{
					"title": "Title",
					"comments": app.Rule{
						"content": "Content",
						"user": app.Rule{
							"username": "Username",
							"nickname": func(user models.TestUser) string {
								return *user.UserInfo.Nickname
							},
							"avatar": func(user models.TestUser) string {
								return fmt.Sprintf("%s%s", config.App.Url, *user.UserInfo.Avatar)
							},
						},
						"children": app.Rule{
							"content": "Content",
							"user": app.Rule{
								"username": "Username",
								"nickname": func(user models.TestUser) string {
									return *user.UserInfo.Nickname
								},
								"avatar": func(user models.TestUser) string {
									return fmt.Sprintf("%s%s", config.App.Url, *user.UserInfo.Avatar)
								},
							},
						},
					},
				},
			},
		},
		"articles": app.Rule{
			"title":   "Title",
			"content": "Content",
			"comments": app.Rule{
				"content": "Content",
				"user": app.Rule{
					"username": "Username",
					"nickname": func(user models.TestUser) string {
						return *user.UserInfo.Nickname
					},
					"avatar": func(user models.TestUser) string {
						return fmt.Sprintf("%s%s", config.App.Url, *user.UserInfo.Avatar)
					},
				},
				"children": app.Rule{
					"content": "Content",
					"user": app.Rule{
						"username": "Username",
						"nickname": func(user models.TestUser) string {
							return *user.UserInfo.Nickname
						},
						"avatar": func(user models.TestUser) string {
							return fmt.Sprintf("%s%s", config.App.Url, *user.UserInfo.Avatar)
						},
					},
				},
			},
		},
		"comments": app.Rule{
			"content": "Content",
			"parent": app.Rule{
				"user": app.Rule{
					"username": "Username",
					"nickname": func(user models.TestUser) string {
						return *user.UserInfo.Nickname
					},
					"avatar": func(user models.TestUser) string {
						return fmt.Sprintf("%s%s", config.App.Url, *user.UserInfo.Avatar)
					},
				},
			},
			"article": app.Rule{
				"id":    "ID",
				"title": "title",
				"user": app.Rule{
					"username": "Username",
					"nickname": func(user models.TestUser) string {
						return *user.UserInfo.Nickname
					},
					"avatar": func(user models.TestUser) string {
						return fmt.Sprintf("%s%s", config.App.Url, *user.UserInfo.Avatar)
					},
				},
			},
		},
	})
	assert.Nil(err)

	prettyJSON, err := json.MarshalIndent(out, "", "  ")
	assert.Nil(err)
	fmt.Println("DB Find Result: ")
	fmt.Println(string(prettyJSON))

	fmt.Println("----------------TestDB 结束----------------")
}
