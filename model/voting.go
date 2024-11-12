package model

type VoteProcess struct {
	IDKode []int          `json:"id_kode"`
	Vote   []NominasiVote `json:"vote"`
}

type NominasiVote struct {
	ID         int `json:"id"`
	JumlahVote int `json:"jumlah_vote"`
}
