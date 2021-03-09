package model

import (
	"wallet/ELGamal"
)

type NewWallet struct {
	Name string `json:"name" form:"name"`
	Id   string `json:"id" form:"id"`
	Str  string `json:"str" form:"str"`
}

type BctoEx struct {
	// 此处的 G1 是否多余
	G1     string          `json:"g1"`
	G2     string          `json:"g2"`
	P      string          `json:"p"`
	H      string          `json:"h"`
	Amount ELGamal.Account `json:"amount"`
}

type ExchangeCoin struct {
	Receipt struct {
		Cmv    string `json:"cmv"`
		Epkrc1 string `json:"epkrc1"`
		Epkrc2 string `json:"epkrc2"`
		Hash   string `json:"hash"` //此次购币交易的交易哈希
	} `json:"receipt"`
	Priv ELGamal.PrivateKey `json:"priv"`
}
