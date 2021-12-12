package observer

import (
	"errors"
	"fmt"

	"github.com/franszhafran/ppl-fp-2021/domain/event"
	"github.com/franszhafran/ppl-fp-2021/domain/repository"
)

type PembelianSuccess struct {
	BarangRepo repository.BarangRepository
}

func (ps *PembelianSuccess) Update(dataType string, rawData interface{}) error {
	data := rawData.(event.BarangTerjual)
	if dataType == "payment_received" {
		barang, err := ps.BarangRepo.FindByID(data.IDBarang)
		if err != nil {
			return err
		}
		barang.Stok -= data.Jumlah
		fmt.Printf("Mengurangi stok untuk barang %s sebanyak %d\n", data.IDBarang, data.Jumlah)
		return ps.BarangRepo.Persist(barang)
	}

	return errors.New("unknown event type received")
}
