package usecase

import "github.com/franszhafran/ppl-fp-2021/domain/model"

type PembelianUsecase interface {
	Create(model.Pembelian) error
	Verify(model.Pembelian, model.Pembayaran) error
}
