package modelos

// Senha representa a senha que sera alterada
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
