package models

type UserModel struct {
	ID       uint64 `json:"id.omitempty"`
	Nome     string `json:"name.omitempty"`
	Nick     string `json:"nick.omitempty"`
	Email    string `json:"email.omitempty"`
	Senha    string `json:"senha.omitempty"`
	// CriadoEm time.Time `json:"criadoEm.omitempty"`
}