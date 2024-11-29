package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/admin/requests"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
)

type RoleController struct {
	app.Controller
}

var Role = new(RoleController)

// Index 获取列表页面
func (*RoleController) Index() (services.Response, error) {
	return res.Ok("admin/role/index", app.Data{}), nil
}

// Create 获取添加页面
func (*RoleController) Create(req *requests.RoleRequest) (services.Response, error) {
	return res.Ok("admin/role/index", app.Data{
		"name": req.Name,
	}), nil
}

// Store 添加数据
func (*RoleController) Store(req *requests.RoleRequest, model *models.AdminRole) (services.Response, error) {
	model.Name = req.Name
	result := db.Create(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Created("创建成功"), nil
}

// Edit 获取修改页面
func (*RoleController) Edit(req *requests.RoleRequest, model *models.AdminRole) (services.Response, error) {
	return res.Ok("admin/role/index", app.Data{
		"model": model,
	}), nil
}

// Update 修改数据
func (*RoleController) Update(req *requests.RoleRequest, model *models.AdminRole) (services.Response, error) {
	model.Name = req.Name
	result := db.Save(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Ok("修改成功"), nil
}

// Show 获取详情
func (*RoleController) Show(req *requests.RoleRequest, model *models.AdminRole) (services.Response, error) {
	return res.Ok("admin/role/index", app.Data{
		"model": model,
	}), nil
}

// Destroy 销毁数据
func (*RoleController) Destroy(model *models.AdminRole) (services.Response, error) {
	result := db.Delete(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.NoContent("删除成功"), nil
}
