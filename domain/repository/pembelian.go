package repository

import (
	"github.com/franszhafran/ppl-fp-2021/domain/model"
)

type PembelianRepository interface {
	Persist(*model.Pembelian) error
}
