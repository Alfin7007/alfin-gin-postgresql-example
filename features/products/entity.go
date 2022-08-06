package products

type Core struct {
	ProductID uint
	Name      string
	Stock     uint
	Detail    string
	Price     uint
}

type Data interface {
	InsertData(Core) error
	SelectAll() ([]Core, error)
	SelectData(ID uint) (Core, error)
}

type Bussiness interface {
	InsertProduct(Core) error
	ShowAll() ([]Core, error)
	ShowProduct(ID uint) (Core, error)
}
