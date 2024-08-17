package entity

type Flat struct {
	Id      int
	HouseId int
	Price   int
	Rooms   int
	Status  string
}

type UsecaseFlatInterface interface {
	CreateFlat(*Flat) error
	UpdateStatusFlat(*Flat) error
}
type RepoFlat interface {
	Post(*Flat) error
	UpdateStatus(*Flat) error
}
