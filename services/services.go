package services

// Services 服务集合的接口
type Services interface {
	Mount()
	SetService(Service)
}

// Service 服务的通用接口
type Service interface {
	Register(Services)
	BindAbility(a Ability)
}

// Ability 将服务能力链接到服务上, 然后由服务来绑定能力.
// 这些服务能力为通用接口, 实际由 App 内实现, 作用是与服务解耦, 保证引用从服务到 App 内引用.
type Ability interface {
	Link(s Service)
}
