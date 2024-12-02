package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/admin/requests"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
)

type AdminController struct {
	app.Controller
	resource    string
	breadcrumbs []map[string]any
}

var (
	Admin = &AdminController{
		Controller: app.Controller{},
		resource:   "/admin/user",
		breadcrumbs: []map[string]any{
			{
				"name": "系统设置",
				"path": "#",
			},
			{
				"name": "员工管理",
				"path": "/admin/user",
			},
		},
	}
)

// Index 获取列表页面
func (a *AdminController) Index(req *requests.AdminIndexRequest) (services.Response, error) {
	return res.Ok("admin/user/index", app.Data{
		"breadcrumbs": a.breadcrumbs,
		"grid": map[string]any{
			"filter": map[string]any{
				"action": a.resource,
				"filters": []map[string]any{
					{
						"label": "用户名",
						"name":  "username",
						"value": req.Username,
						"type":  "text",
					},
					{
						"label": "邮箱",
						"name":  "email",
						"value": req.Email,
						"type":  "text",
					},
					{
						"label": "昵称",
						"name":  "nickname",
						"value": req.Nickname,
						"type":  "text",
					},
				},
			},
			"tools": map[string]any{
				"disableCreateButton": false,
			},
			"table": map[string]any{
				"pinRow": true,
				"pinCol": false,
			},
		},
	}), nil
}

// Create 获取添加页面
func (*AdminController) Create(req *requests.AdminRequest) (services.Response, error) {
	return res.Ok("admin/user/index", app.Data{
		"name": req.Name,
	}), nil
}

// Store 添加数据
func (*AdminController) Store(req *requests.AdminRequest, admin *models.AdminUser) (services.Response, error) {
	admin.Username = req.Name
	result := db.Create(admin)
	if result.Error != nil {
		return nil, exc.BadRequest(result.Error)
	}

	return res.Created("创建成功"), nil
}

// Edit 获取修改页面
func (*AdminController) Edit(model *models.AdminUser) (services.Response, error) {
	return res.Ok("admin/user/index", app.Data{
		"model": model,
	}), nil
}

// Update 修改数据
func (*AdminController) Update(req *requests.AdminRequest, model *models.AdminUser) (services.Response, error) {
	model.Username = req.Name
	result := db.Save(model)
	if result.Error != nil {
		return nil, exc.BadRequest(result.Error)
	}

	return res.Ok("修改成功"), nil
}

// Show 获取详情
func (*AdminController) Show(model *models.AdminUser) (services.Response, error) {
	return res.Ok("admin/user/index", app.Data{
		"model": model,
	}), nil
}

// Destroy 销毁数据
func (*AdminController) Destroy(model *models.AdminUser) (services.Response, error) {
	result := db.Delete(model)
	if result.Error != nil {
		return nil, exc.BadRequest(result.Error)
	}

	return res.NoContent("删除成功"), nil
}
