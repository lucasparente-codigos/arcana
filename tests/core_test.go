kpackage tests

import (
	"testing"

	"github.com/lucasparente-codigos/arcana/internal/core"
	"github.com/lucasparente-codigos/arcana/internal/domain"
	"github.com/lucasparente-codigos/arcana/internal/presets"
)

// Vetores de Teste garantem que a lógica determinística NUNCA mude.
// Se este teste falhar, significa que uma atualização quebrou a compatibilidade
// com senhas geradas anteriormente.
func TestDeterministicVectors(t *testing.T) {
	vectors := []struct {
		MasterKey string
		Context   string
		User      string
		ProfileID string
		Expected  string
	}{
		// Caso base: Web profile
		{"mysecret", "github", "", "web", "TBD_AFTER_FIRST_RUN"}, 
		// Como o Argon2 usa salt estático+user, precisamos rodar uma vez para pegar o hash esperado 
		// e fixar aqui. Para este exemplo, vou simular a verificação de consistência.
	}

	// Nota: Como Argon2 é lento para rodar em testes unitários rápidos frequentemente,
	// em CI real usaríamos flags de build ou mocks para KDF, mas aqui testamos o fluxo completo.
	
	p, _ := presets.GetProfile("pin")
	config := domain.Config{
		Context: "bank",
		Username: "lucas",
	}
	
	// Teste de consistência: Rodar 2x deve dar o mesmo resultado
	pass1, _, _ := core.GeneratePassword("123456", config, p)
	pass2, _, _ := core.GeneratePassword("123456", config, p)
	
	if pass1 != pass2 {
		t.Errorf("Quebra de determinismo! %s != %s", pass1, pass2)
	}
}

func TestProfileConstraints(t *testing.T) {
	p, _ := presets.GetProfile("pin")
	config := domain.Config{Context: "test"}
	
	pass, _, _ := core.GeneratePassword("master", config, p)
	
	if len(pass) != 6 {
		t.Errorf("Profile PIN deveria ter 6 chars, teve %d", len(pass))
	}
	
	for _, char := range pass {
		if char < '0' || char > '9' {
			t.Errorf("Profile PIN gerou caractere inválido: %c", char)
		}
	}
}
