package main

import (
	"fmt"
	"time"

	"github.com/franszhafran/ppl-fp-2021/domain/builder"
	"github.com/franszhafran/ppl-fp-2021/domain/model"
	observerIf "github.com/franszhafran/ppl-fp-2021/domain/observer"
	"github.com/franszhafran/ppl-fp-2021/infrastructure/observer"
	"github.com/franszhafran/ppl-fp-2021/infrastructure/repository"
	"github.com/franszhafran/ppl-fp-2021/presentation/usecase"
)

func checkBuilder() *model.Pembelian {
	pembelianBuilder := new(builder.PembelianBuilder)
	pembelian, err := pembelianBuilder.
		WithIDBarang("gta-v-2021").
		WithIDPelanggan("user-001").
		WithAlamat("Jl Setiabudi No 1 South Jakarta").
		WithJumlahBarang(1).
		WithNominal(250000).
		Build()
	if err == nil {
		fmt.Printf("%+v\n", pembelian)
	}
	return pembelian
}

func terimaPayment(pembelian *model.Pembelian) {
	fmt.Println("")
	fmt.Println("== Observer: stok akan berkurang ketika terjadi pembelian ==")
	uc := usecase.PembelianUsecase{
		EventListeners: []observerIf.Listener{
			&observer.PembelianSuccess{
				BarangRepo: new(repository.BarangRepository),
			},
		},
	}

	pembayaran := model.Pembayaran{
		ID:               "payment-001",
		IDPembelian:      pembelian.ID,
		Metode:           "bank",
		StatusPembayaran: "paid",
		Waktu:            time.Now(),
	}
	err := uc.Verify(*pembelian, pembayaran)
	if err != nil {
		fmt.Println("terjadi kesalahan")
	}
}

func main() {
	pembelian := checkBuilder()
	terimaPayment(pembelian)
}
