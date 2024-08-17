package entity

type House struct {
	Address   string
	Year      string
	Developer string
	HouseId   int
	CreateAt  string
	UpdateAt  string
}

type HouseUseCase interface {
	CreateHouse(*House) (*House, error)
}

type HouseRepoInterface interface {
	CreateHouse(*House) (*House, error)
}
