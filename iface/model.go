package iface

type Model interface {
	In(request Request, r map[string]any) (Model, error)
	Out(data any, r map[string]any) (any, error)
	SetModel(i Model)
}
