package service

type Auto struct {
	ErrorHandler func(error)
}

func (a *Auto) Register(service Service, deps []string) {

}

//this will block until all services registed quit
func (a *Auto) Run() {

}
