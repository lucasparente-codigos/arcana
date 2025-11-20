package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lucasparente-codigos/arcana/internal/core"
	"github.com/lucasparente-codigos/arcana/internal/domain"
	"github.com/lucasparente-codigos/arcana/internal/presets"
	"github.com/lucasparente-codigos/arcana/internal/ui"
)

var (
	version = "v1.0.0"
)

func main() {
	// 1. Configuração de Flags
	profileFlag := flag.String("p", "web", "Perfil: web, sysadmin, legacy, pin")
	lengthFlag := flag.Int("l", 0, "Sobrescrever comprimento da senha")
	userFlag := flag.String("u", "", "Username (adiciona entropia ao salt)")
	stealthFlag := flag.Bool("s", false, "Modo Stealth (QR Code)")
	explainFlag := flag.Bool("x", false, "Explicar processo de geração (Auditoria)")
	rawFlag := flag.Bool("raw", false, "Não normalizar o contexto (Case Sensitive)")
	timeoutFlag := flag.Int("t", 30, "Tempo (segundos) para limpar a tela")
	verFlag := flag.Bool("v", false, "Exibir versão")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "🔮 Arcana %s - Gerador de Segredos Determinístico\n", version)
		fmt.Fprintf(os.Stderr, "Uso: arcana [flags] <contexto>\n\nFlags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *verFlag {
		fmt.Printf("Arcana %s\n", version)
		os.Exit(0)
	}

	// 2. Validação de Argumentos
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Erro: Contexto é obrigatório.")
		fmt.Println("Exemplo: arcana -p web \"github.com\"")
		os.Exit(1)
	}
	contextStr := args[0]

	// 3. Carregar Perfil
	selectedProfile, err := presets.GetProfile(*profileFlag)
	if err != nil {
		fmt.Printf("Erro: Perfil '%s' desconhecido.\nPerfis disponíveis: web, sysadmin, legacy, pin\n", *profileFlag)
		os.Exit(1)
	}

	// 4. Input Seguro da Master Key
	fmt.Println("🔮 Arcana: Ambiente Seguro Inicializado")
	masterKey, err := ui.ReadSecret("Digite sua Frase-Mestra: ")
	if err != nil {
		fmt.Println("\nErro ao ler entrada:", err)
		os.Exit(1)
	}
	if len(masterKey) == 0 {
		fmt.Println("Erro: Frase-mestra não pode ser vazia.")
		os.Exit(1)
	}

	// 5. Execução Core
	config := domain.Config{
		ProfileID:      *profileFlag,
		Context:        contextStr,
		Username:       *userFlag,
		LengthOverride: *lengthFlag,
		Stealth:        *stealthFlag,
		Explain:        *explainFlag,
		Timeout:        *timeoutFlag,
		RawContext:     *rawFlag,
	}

	password, audit, err := core.GeneratePassword(masterKey, config, selectedProfile)
	
	// Limpar master key da memória imediatamente (best effort em Go)
	masterKey = "" 
	
	if err != nil {
		fmt.Printf("Erro fatal na geração: %v\n", err)
		os.Exit(1)
	}

	// 6. Apresentação
	ui.DisplayResult(password, audit, config)
}
