package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/admin/requests"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
)

type MenuController struct {
	app.Controller
}

var Menu = new(MenuController)

// Index 获取列表页面
func (*MenuController) Index() (services.Response, error) {
	return res.Ok("admin/menu/index", app.Data{}), nil
}

// Create 获取添加页面
func (*MenuController) Create(req *requests.MenuRequest) (services.Response, error) {
	return res.Ok("admin/menu/index", app.Data{
		"name": req.Name,
	}), nil
}

// Store 添加数据
func (*MenuController) Store(req *requests.MenuRequest, model *models.AdminMenu) (services.Response, error) {
	model.Name = *req.Name
	result := db.Create(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Created("创建成功"), nil
}

// Edit 获取修改页面
func (*MenuController) Edit(req *requests.MenuRequest, model *models.AdminMenu) (services.Response, error) {
	return res.Ok("admin/menu/index", app.Data{
		"model": model,
	}), nil
}

// Update 修改数据
func (*MenuController) Update(req *requests.MenuRequest, model *models.AdminMenu) (services.Response, error) {
	model.Name = *req.Name
	result := db.Save(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Ok("修改成功"), nil
}

// Show 获取详情
func (*MenuController) Show(req *requests.MenuRequest, model *models.AdminMenu) (services.Response, error) {
	return res.Ok("admin/menu/index", app.Data{
		"model": model,
	}), nil
}

// Destroy 销毁数据
func (*MenuController) Destroy(model *models.AdminMenu) (services.Response, error) {
	result := db.Delete(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.NoContent("删除成功"), nil
}
