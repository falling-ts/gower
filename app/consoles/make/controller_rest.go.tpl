package controllers

import (
	"{{.Module}}/app"
	"{{.Module}}/app/http/requests"
	"{{.Module}}/app/models"
	"{{.Module}}/services"
)

type {{.UpCamel}}Controller struct {
	app.Controller
}

var {{.UpCamel}} = new({{.UpCamel}}Controller)

// Index 获取列表页面
func (*{{.UpCamel}}Controller) Index(req *requests.{{.UpCamel}}Request) (services.Response, error) {
	return res.Ok("模板的 define 名称", app.Data{
		"name": req.Name,
	}), nil
}

// Create 获取添加页面
func (*{{.UpCamel}}Controller) Create(req *requests.{{.UpCamel}}Request) (services.Response, error) {
	return res.Ok("模板的 define 名称", app.Data{
		"name": req.Name,
	}), nil
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

// Edit 获取修改页面
func (*{{.UpCamel}}Controller) Edit(req *requests.{{.UpCamel}}Request, model *models.{{.UpCamel}}) (services.Response, error) {
	return res.Ok("模板的 define 名称", app.Data{
		"model": model,
	}), nil
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
	return res.Ok("模板的 define 名称", app.Data{
		"model": model,
	}), nil
}

// Destroy 销毁数据
func (*{{.UpCamel}}Controller) Destroy(model *models.{{.UpCamel}}) (services.Response, error) {
	result := db.Delete(model)
	if result.Error != nil {
		return nil, excp.BadRequest(result.Error)
	}

	return res.NoContent("删除成功"), nil
}
