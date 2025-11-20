package presets

import (
	"errors"
	"github.com/lucasparente-codigos/arcana/internal/domain"
)

const (
	AlphaLower = "abcdefghijklmnopqrstuvwxyz"
	AlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers    = "0123456789"
	Symbols    = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	
	// Charsets combinados
	Alphanumeric = AlphaLower + AlphaUpper + Numbers
	AllPrintable = Alphanumeric + Symbols
)

var profiles = map[string]domain.Profile{
	"web": {
		ID:          "web",
		Description: "Padrão balanceado (16 chars, símbolos seguros)",
		Length:      16,
		Alphabet:    Alphanumeric + "!@#$%^&*",
		Require:     "Alphanumeric + Common Symbols",
	},
	"sysadmin": {
		ID:          "sysadmin",
		Description: "Paranóico/Infra (32 chars, full ASCII)",
		Length:      32,
		Alphabet:    AllPrintable,
		Require:     "Full ASCII",
	},
	"legacy": {
		ID:          "legacy",
		Description: "Compatibilidade antiga (14 chars, sem símbolos)",
		Length:      14,
		Alphabet:    Alphanumeric,
		Require:     "Alphanumeric Only",
	},
	"pin": {
		ID:          "pin",
		Description: "Código numérico (6 dígitos)",
		Length:      6,
		Alphabet:    Numbers,
		Require:     "Numbers Only",
	},
}

func GetProfile(id string) (domain.Profile, error) {
	if p, ok := profiles[id]; ok {
		return p, nil
	}
	return domain.Profile{}, errors.New("profile not found: " + id)
}

func ListProfiles() map[string]domain.Profile {
	return profiles
}
