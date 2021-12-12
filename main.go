package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/franszhafran/ppl-fp-2021/application/usecase"
	"github.com/franszhafran/ppl-fp-2021/domain/builder"
	"github.com/franszhafran/ppl-fp-2021/domain/model"
	observerIf "github.com/franszhafran/ppl-fp-2021/domain/observer"
	"github.com/franszhafran/ppl-fp-2021/infrastructure/database"
	"github.com/franszhafran/ppl-fp-2021/infrastructure/observer"
	"github.com/franszhafran/ppl-fp-2021/infrastructure/repository"
)

func checkBuilder() *model.Pembelian {
	fmt.Println("== Builder: membuat model pembelian ==")
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

	barangRepo := repository.TestingBarangRepository{
		DB: database.GetDBClient(),
	}

	// strategy pattern -> bisa menggunakan custom repository di BarangRepo
	uc := usecase.PembelianUsecase{
		EventListeners: []observerIf.Listener{
			&observer.PembelianSuccess{
				BarangRepo: &barangRepo,
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

func simpanPembelian(pembelian *model.Pembelian, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("")
	fmt.Println("== Try to trigger db conn ==")

	barangRepo := repository.TestingBarangRepository{
		DB: database.GetDBClient(),
	}
	pembelianRepo := repository.TestingPembelianRepository{
		DB: database.GetDBClient(),
	}

	// strategy pattern -> bisa menggunakan custom repository di BarangRepo
	uc := usecase.PembelianUsecase{
		EventListeners: []observerIf.Listener{
			&observer.PembelianSuccess{
				BarangRepo: &barangRepo,
			},
		},
		PembelianRepo: &pembelianRepo,
	}

	uc.Create(*pembelian)
}

func main() {
	var wg sync.WaitGroup
	pembelian := checkBuilder()
	wg.Add(1)
	go simpanPembelian(pembelian, &wg)
	terimaPayment(pembelian)
	wg.Wait()
}
