package exception

type Err struct {
}

func New() *Err {
	return &Err{}
}
