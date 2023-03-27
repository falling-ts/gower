package controllers

import (
	"github.com/falling-ts/gower/app"
	"github.com/falling-ts/gower/app/api/requests"
	"github.com/falling-ts/gower/app/models"
	"github.com/falling-ts/gower/services"
)

type {{.UpCamel}}Controller struct {
	app.Controller
}

var {{.UpCamel}} = new({{.UpCamel}}Controller)

// Index 获取列表页面
func (*{{.UpCamel}}Controller) Index(req *requests.{{.UpCamel}}Request) (services.Response, error) {
	return res.Ok("获取成功", app.Data{}), nil
}

// Store 添加数据
func (*{{.UpCamel}}Controller) Store(req *requests.{{.UpCamel}}Request, model *models.{{.UpCamel}}) (services.Response, error) {
	model.Name = req.Name
	result := db.Create(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Created("创建成功"), nil
}

// Update 修改数据
func (*{{.UpCamel}}Controller) Update(req *requests.{{.UpCamel}}Request, model *models.{{.UpCamel}}) (services.Response, error) {
	model.Name = req.Name
	result := db.Save(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.Ok("修改成功"), nil
}

// Show 获取详情
func (*{{.UpCamel}}Controller) Show(req *requests.{{.UpCamel}}Request, model *models.{{.UpCamel}}) (services.Response, error) {
	return res.Ok("获取成功", app.Data{}), nil
}

// Destroy 销毁数据
func (*{{.UpCamel}}Controller) Destroy(model *models.{{.UpCamel}}) (services.Response, error) {
	result := db.Delete(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.NoContent("删除成功"), nil
}
