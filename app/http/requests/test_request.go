package requests

type TestRequest struct {
	request
	Test string `form:"test" json:"test" binding:"required" zh:"测试"`
}

type Test1Request struct {
	TestRequest
	Test1 int `form:"test1" json:"test2" binding:"required,number", zh:"测试1"`
}
