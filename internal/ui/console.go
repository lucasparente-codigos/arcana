package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/lucasparente-codigos/arcana/internal/domain"
	"github.com/skip2/go-qrcode"
	"golang.org/x/term"
)

// ReadSecret lê a entrada do usuário sem ecoar no terminal
func ReadSecret(prompt string) (string, error) {
	fmt.Print(prompt)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // Nova linha após enter
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

// DisplayResult gerencia como o segredo é mostrado (Texto ou QR)
func DisplayResult(password string, audit *domain.AuditResult, config domain.Config) {
	if audit != nil {
		printAudit(audit)
	}

	if config.Stealth {
		displayQR(password)
	} else {
		fmt.Printf("\n🔑 Senha Gerada: \033[1;32m%s\033[0m\n", password)
	}

	if config.Timeout > 0 {
		startCountdown(config.Timeout)
	}
}

func printAudit(a *domain.AuditResult) {
	fmt.Println("\n--- 🛡️  AUDITORIA DE GERAÇÃO ---")
	fmt.Printf("Algoritmo:   %s\n", a.Algorithm)
	fmt.Printf("Contexto:    %s\n", a.ContextUsed)
	fmt.Printf("Perfil:      %s\n", a.ProfileUsed)
	fmt.Printf("Parâmetros:  Time=%d, Mem=%dKB, Threads=%d\n", a.ArgonTime, a.ArgonMemory, a.ArgonThreads)
	
	entropyQual := "Fraca"
	if a.EntropyBits > 50 { entropyQual = "Razoável" }
	if a.EntropyBits > 80 { entropyQual = "Forte" }
	if a.EntropyBits > 120 { entropyQual = "Excessiva (Excelente)" }

	fmt.Printf("Entropia:    %.2f bits (%s)\n", a.EntropyBits, entropyQual)
	fmt.Println("--------------------------------")
}

func displayQR(content string) {
	fmt.Println("\n📱 Modo Stealth (QR Code):")
	// Nível M de recuperação de erro é suficiente para terminais
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		fmt.Println("Erro ao gerar QR:", err)
		fmt.Println("Fallback Texto:", content)
		return
	}
	// Imprime string ASCII do QR
	fmt.Println(qr.ToString(false))
	fmt.Println("Escaneie agora. A tela será limpa em breve.")
}

func startCountdown(seconds int) {
	fmt.Printf("\nLimpando tela em %d segundos...", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	ClearScreen()
}

// ClearScreen tenta limpar o terminal usando comandos ANSI e syscalls
func ClearScreen() {
	// ANSI Reset
	fmt.Print("\033[H\033[2J")
	
	// Comando do sistema como fallback
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// CopyToClipboard tenta copiar (implementação básica cross-platform)
// Nota: Em produção, bibliotecas específicas são melhores, mas aqui evitamos deps pesadas.
func CopyToClipboard(text string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		// Tenta xclip, se falhar tenta xsel (pode requerer instalação)
		cmd = exec.Command("xclip", "-selection", "clipboard")
	default:
		return fmt.Errorf("clipboard não suportado nativamente neste SO")
	}
	
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}
