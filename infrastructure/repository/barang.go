package repository

import "github.com/franszhafran/ppl-fp-2021/domain/model"

type BarangRepository struct {
}

func (b *BarangRepository) Persist(barang *model.Barang) error {
	// implement here
	return nil
}

func (b *BarangRepository) FindByID(id string) (*model.Barang, error) {
	barang := model.Barang{
		Stok: 50,
	}
	return &barang, nil
}

func (b *BarangRepository) Delete(id string) error {
	// implement here
	return nil
}
