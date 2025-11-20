package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strings"

	"github.com/lucasparente-codigos/arcana/internal/domain"
	"golang.org/x/crypto/argon2"
)

const (
	StaticSalt = "ARCANA_STATIC_SALT_V1_DO_NOT_CHANGE"
	ArgonTime  = 1
	ArgonMem   = 64 * 1024 // 64MB
	ArgonThreads = 4
	KeyLen     = 32
)

// GeneratePassword executa o pipeline criptográfico
func GeneratePassword(masterKey string, config domain.Config, profile domain.Profile) (string, *domain.AuditResult, error) {
	// 1. Normalização
	ctxInput := config.Context
	if !config.RawContext {
		ctxInput = strings.ToLower(strings.TrimSpace(config.Context))
	}
	
	// 2. Preparação do Salt (Static + Username)
	// Isso garante que o mesmo usuário + mesma mestra gerem a mesma seed
	fullSalt := sha256.Sum256([]byte(StaticSalt + config.Username))

	// 3. Key Derivation Function (Argon2id)
	// Transforma a senha mestra em uma chave criptográfica de alta entropia
	derivedKey := argon2.IDKey([]byte(masterKey), fullSalt[:], ArgonTime, ArgonMem, ArgonThreads, KeyLen)

	// 4. Geração Determinística (HMAC-SHA256)
	// Assina o contexto e o ID do perfil com a chave derivada
	mac := hmac.New(sha256.New, derivedKey)
	mac.Write([]byte(ctxInput))
	mac.Write([]byte(profile.ID))
	hashBytes := mac.Sum(nil)

	// Determinar comprimento final
	length := profile.Length
	if config.LengthOverride > 0 {
		length = config.LengthOverride
	}

	// 5. Mapping (Bytes -> Charset) com Rejection Sampling
	password := mapBytesToCharset(hashBytes, profile.Alphabet, length)

	// 6. Auditoria (se solicitada)
	var audit *domain.AuditResult
	if config.Explain {
		entropy := float64(length) * math.Log2(float64(len(profile.Alphabet)))
		audit = &domain.AuditResult{
			Algorithm:     "Argon2id(Master) -> HMAC-SHA256(Context) -> UnbiasedMap",
			ArgonMemory:   ArgonMem,
			ArgonTime:     ArgonTime,
			EntropyBits:   entropy,
			ProfileUsed:   profile.ID,
			ContextUsed:   ctxInput,
			SaltSignature: hex.EncodeToString(fullSalt[:4]) + "...",
		}
	}

	return password, audit, nil
}

// mapBytesToCharset converte bytes aleatórios em caracteres do alfabeto
// usando rejection sampling para evitar viés de módulo.
func mapBytesToCharset(hash []byte, alphabet string, length int) string {
	var result strings.Builder
	
	// Se precisarmos de mais bytes do que o hash fornece,
	// usamos o hash como semente para gerar um stream (simplificado aqui via expansão SHA256 repetida)
	// Para senhas típicas (< 32 chars), o hash original basta, mas vamos garantir expansão.
	stream := hash
	
	idx := 0
	for result.Len() < length {
		if idx >= len(stream) {
			// Expande o stream fazendo hash do anterior
			newHash := sha256.Sum256(stream)
			stream = newHash[:]
			idx = 0
		}

		b := int(stream[idx])
		idx++

		// Unbiased Rejection Sampling
		// Rejeita bytes que caem na faixa "sobra" da divisão para garantir distribuição plana
		limit := 256 - (256 % len(alphabet))
		if b >= limit {
			continue
		}

		result.WriteByte(alphabet[b%len(alphabet)])
	}

	return result.String()
}
