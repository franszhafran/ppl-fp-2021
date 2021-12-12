package model

import "time"

type StatusPembelian int

const (
	STATUS_MENUNGGU_PEMBAYARAN = iota
	STATUS_DALAM_PENGEMASAN
	STATUS_DIKIRIM
)

type Pembelian struct {
	ID               string
	IDPelanggan      string
	IDBarang         string
	NomorInvoice     string
	JumlahBarang     int
	Nominal          int
	Waktu            time.Time
	Status           StatusPembelian
	AlamatPengiriman string
}
