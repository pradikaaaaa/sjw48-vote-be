package model

type KodeVote struct {
	ID        uint   `json:"id"`
	Kode      string `json:"kode"`
	Jumlah    int    `json:"jumlah"`
	CreatedAt *string
	UpdatedAt *string
}

type RequestVote struct {
	JumlahVote int `json:"jumlah_vote"`
	Kuota      int `json:"kuota"`
}

type RequestKodeVote struct {
	KodeVote []string `json:"kode_vote"`
}

func (KodeVote) TableName() string {
	return "kode_vote"
}
