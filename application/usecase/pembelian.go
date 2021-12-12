package usecase

import (
	"fmt"

	"github.com/franszhafran/ppl-fp-2021/domain/event"
	"github.com/franszhafran/ppl-fp-2021/domain/model"
	"github.com/franszhafran/ppl-fp-2021/domain/observer"
	"github.com/franszhafran/ppl-fp-2021/domain/repository"
)

// create(model.Barang) error
// 	verify(model.Pembelian, model.Pembayaran) error

type PembelianUsecase struct {
	EventListeners []observer.Listener
	PembelianRepo  repository.PembelianRepository
}

func (p *PembelianUsecase) Create(pembelian model.Pembelian) {
	p.PembelianRepo.Persist(&pembelian)
}

func (p *PembelianUsecase) Verify(pembelian model.Pembelian, pembayaran model.Pembayaran) error {
	fmt.Printf("Pembayaran diverifikasi untuk pembelian dengan ID %s\n", pembelian.ID)
	data := event.BarangTerjual{
		IDBarang: pembelian.IDBarang,
		Jumlah:   pembelian.JumlahBarang,
	}
	p.notify(data)

	return nil
}

func (p *PembelianUsecase) notify(data interface{}) error {
	for _, listener := range p.EventListeners {
		errLocal := listener.Update("payment_received", data)
		if errLocal != nil {
			return errLocal
		}
	}

	return nil
}
