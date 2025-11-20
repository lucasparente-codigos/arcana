package domain

// Profile define as regras para geração da senha
type Profile struct {
	ID          string
	Description string
	Length      int
	Alphabet    string
	Require     string // Descrição textual de requisitos (ex: "symbols")
}

// Config armazena as opções passadas via CLI
type Config struct {
	ProfileID    string
	Context      string
	Username     string
	LengthOverride int
	Stealth      bool
	Explain      bool
	Timeout      int
	RawContext   bool // Se true, não normaliza o contexto (case sensitive)
}

// AuditResult contém dados para o comando --explain
type AuditResult struct {
	Algorithm     string
	ArgonMemory   uint32
	ArgonTime     uint32
	ArgonThreads  uint8
	EntropyBits   float64
	ProfileUsed   string
	ContextUsed   string
	SaltSignature string // Hash parcial do salt para verificação
}
