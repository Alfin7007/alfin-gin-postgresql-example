package users

type Core struct {
	UserID   int
	Name     string
	Email    string
	Password string
}

type Data interface {
	InsertData(Core) error
	FindUser(Email string) (Core, error)
}

type Bussiness interface {
	Register(Core) error
	Login(Core) (id int, token string, err error)
}
