package repository

import (
	"github.com/franszhafran/ppl-fp-2021/domain/model"
	"github.com/franszhafran/ppl-fp-2021/infrastructure/database"
)

type TestingPembelianRepository struct {
	DB *database.DBClient
}

func (p *TestingPembelianRepository) Persist(pembelian *model.Pembelian) error {
	p.DB.Query("persist into pembelian ... ()")
	return nil
}
