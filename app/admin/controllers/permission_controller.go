package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/admin/requests"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
)

type PermissionController struct {
	app.Controller
}

var Permission = new(PermissionController)

// Index 获取列表页面
func (*PermissionController) Index() (services.Response, error) {
	return res.Ok("admin/permission/index", app.Data{}), nil
}

// Create 获取添加页面
func (*PermissionController) Create(req *requests.PermissionRequest) (services.Response, error) {
	return res.Ok("admin/permission/index", app.Data{
		"name": req.Name,
	}), nil
}

// Store 添加数据
func (*PermissionController) Store(req *requests.PermissionRequest, model *models.AdminPermission) (services.Response, error) {
	model.Name = req.Name
	result := db.Create(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Created("创建成功"), nil
}

// Edit 获取修改页面
func (*PermissionController) Edit(req *requests.PermissionRequest, model *models.AdminPermission) (services.Response, error) {
	return res.Ok("admin/permission/index", app.Data{
		"model": model,
	}), nil
}

// Update 修改数据
func (*PermissionController) Update(req *requests.PermissionRequest, model *models.AdminPermission) (services.Response, error) {
	model.Name = req.Name
	result := db.Save(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Ok("修改成功"), nil
}

// Detail 获取详情
func (*PermissionController) Detail(req *requests.PermissionRequest, model *models.AdminPermission) (services.Response, error) {
	return res.Ok("admin/permission/index", app.Data{
		"model": model,
	}), nil
}

// Destroy 销毁数据
func (*PermissionController) Destroy(model *models.AdminPermission) (services.Response, error) {
	result := db.Delete(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.NoContent("删除成功"), nil
}
