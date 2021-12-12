package builder

import (
	"errors"
	"fmt"
	"time"

	"github.com/franszhafran/ppl-fp-2021/domain/model"
)

type PembelianBuilder struct {
	id               string
	idPelanggan      string
	idBarang         string
	nomorInvoice     string
	jumlahBarang     int
	nominal          int
	waktu            time.Time
	status           model.StatusPembelian
	alamatPengiriman string
	timeSet          bool
}

func (p *PembelianBuilder) WithIDPelanggan(id string) *PembelianBuilder {
	p.idPelanggan = id
	return p
}

func (p *PembelianBuilder) WithIDBarang(id string) *PembelianBuilder {
	p.idBarang = id
	return p
}

func (p *PembelianBuilder) WithJumlahBarang(jumlah int) *PembelianBuilder {
	p.jumlahBarang = jumlah
	return p
}

func (p *PembelianBuilder) WithNominal(nominal int) *PembelianBuilder {
	p.nominal = nominal
	return p
}

func (p *PembelianBuilder) WithAlamat(alamat string) *PembelianBuilder {
	p.alamatPengiriman = alamat
	return p
}

func (p *PembelianBuilder) WithStatus(status model.StatusPembelian) *PembelianBuilder {
	p.status = status
	return p
}

func (p *PembelianBuilder) WithWaktu(waktu time.Time) *PembelianBuilder {
	p.timeSet = true
	p.waktu = waktu
	return p
}

func (p *PembelianBuilder) Build() (*model.Pembelian, error) {
	if p.idPelanggan == "" || p.idBarang == "" || p.jumlahBarang == 0 || p.nominal == 0 || p.alamatPengiriman == "" {
		return nil, errors.New("Unsufficient data input to create model Pembelian")
	}

	if !p.timeSet {
		p.waktu = time.Now()
	}

	pembelian := model.Pembelian{
		ID:               "trx-001",
		IDBarang:         p.idBarang,
		IDPelanggan:      p.idPelanggan,
		JumlahBarang:     p.jumlahBarang,
		Nominal:          p.nominal,
		AlamatPengiriman: p.alamatPengiriman,
		Status:           p.status,
		Waktu:            p.waktu,
	}
	pembelian.NomorInvoice = fmt.Sprintf("%s-%d", p.idPelanggan, p.waktu.Unix())

	return &pembelian, nil
}

func (p *PembelianBuilder) Reset() {
	r := new(PembelianBuilder)
	p = r
}
