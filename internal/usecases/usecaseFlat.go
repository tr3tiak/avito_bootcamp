package usecases

import "avito_bootcamp/internal/entity"

type UsecaseFlat struct {
	repo entity.RepoFlat
}

func InitUseCaseFlat(rf entity.RepoFlat) UsecaseFlat {
	return UsecaseFlat{
		repo: rf,
	}
}

func (uf UsecaseFlat) CreateFlat(flat *entity.Flat) error {
	err := uf.repo.Post(flat)
	if err != nil {
		return err
	}
	return nil
}

func (uf UsecaseFlat) UpdateStatusFlat(flat *entity.Flat) error {
	err := uf.repo.UpdateStatus(flat)
	if err != nil {
		return err
	}
	return nil
}
