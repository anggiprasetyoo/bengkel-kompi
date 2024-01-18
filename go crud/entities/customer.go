package entities

type Customer struct {
	Id          int64
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	NomorHp     string `validate:"required" label:"Nomor HP"`
	Merk        string `validate:"required"`
	Alamat      string `validate:"required"`
	Masalah     string `validate:"required"`
}
