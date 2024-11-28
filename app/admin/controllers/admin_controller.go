package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/admin/requests"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
)

type AdminController struct {
	app.Controller
}

var Admin = new(AdminController)

// Index 获取列表页面
func (*AdminController) Index() (services.Response, error) {
	return res.Ok("admin/user/index", app.Data{
		"breadcrumbs": []map[string]any{
			{
				"name": "系统设置",
				"path": "#",
			},
			{
				"name": "员工管理",
				"path": "/admin/user",
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
		return nil, excp.BadRequest(result.Error)
	}

	return res.Created("创建成功"), nil
}

// Edit 获取修改页面
func (*AdminController) Edit(req *requests.AdminRequest, model *models.AdminUser) (services.Response, error) {
	return res.Ok("admin/user/index", app.Data{
		"model": model,
	}), nil
}

// Update 修改数据
func (*AdminController) Update(req *requests.AdminRequest, model *models.AdminUser) (services.Response, error) {
	model.Username = req.Name
	result := db.Save(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Ok("修改成功"), nil
}

// Detail 获取详情
func (*AdminController) Detail(req *requests.AdminRequest, model *models.AdminUser) (services.Response, error) {
	return res.Ok("admin/user/index", app.Data{
		"model": model,
	}), nil
}

// Destroy 销毁数据
func (*AdminController) Destroy(model *models.AdminUser) (services.Response, error) {
	result := db.Delete(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.NoContent("删除成功"), nil
}
