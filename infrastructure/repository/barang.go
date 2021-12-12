package repository

import (
	"github.com/franszhafran/ppl-fp-2021/domain/model"
	"github.com/franszhafran/ppl-fp-2021/infrastructure/database"
)

type TestingBarangRepository struct {
	DB *database.DBClient
}

func (b *TestingBarangRepository) Persist(barang *model.Barang) error {
	// implement here
	b.DB.Query("insert into ... ()")
	return nil
}

func (b *TestingBarangRepository) FindByID(id string) (*model.Barang, error) {
	barang := model.Barang{
		Stok: 50,
	}
	b.DB.Query("find by id ... ()")
	return &barang, nil
}

func (b *TestingBarangRepository) Delete(id string) error {
	// implement here
	b.DB.Query("delete by id ... ()")
	return nil
}
