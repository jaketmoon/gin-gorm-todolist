package Todo

type ModuleTodo struct{}

func (u *ModuleTodo) Init() {}

func (u *ModuleTodo) GetName() string {
	return "Todo"
}
