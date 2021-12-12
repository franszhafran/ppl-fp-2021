package repository

import (
	"github.com/franszhafran/ppl-fp-2021/domain/model"
)

type BarangRepository interface {
	Persist(*model.Barang) error
	FindByID(string) (*model.Barang, error)
	Delete(string) error
}
