package model

import "time"

type Pembayaran struct {
	ID               string
	IDPembelian      string
	Metode           string
	StatusPembayaran string
	Waktu            time.Time
}
