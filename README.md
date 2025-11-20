# Arcana

Arcana é uma ferramenta CLI escrita em Go para geração de senhas de alta entropia, combinando auditoria transparente, modos determinísticos, perfis inteligentes e um modo stealth com exibição via QR code. O projeto foi concebido com foco em simplicidade, segurança real e portabilidade total, distribuído como binário único.

---

## 📌 Visão Geral

Arcana resolve um problema essencial: gerar senhas fortes e consistentes sem depender de armazenamento persistente ou de ferramentas pesadas. Seu modo determinístico permite a criação de senhas reproduzíveis a partir de uma frase-mestra e de um contexto, enquanto seus perfis prontos atendem diferentes cenários operacionais — de autenticação web a ambientes restritivos.

Recursos principais:

* Modo determinístico consciente (frase-mestra + contexto)
* Perfis inteligentes (web-safe, sysadmin, legacy-safe, PIN)
* Modo stealth com exibição opcional em QR code
* Estimativa de entropia e avisos de insegurança
* Explicação auditável do processo de geração
* Timer opcional para limpeza da saída

---

## ⚙️ Instalação

### Requisitos

* Go 1.22+
* Ambiente Unix-like ou Windows compatível com scripts shell (opcional)

### Instalação via `go install`

```
go install github.com/lucasparente-codigos/arcana@latest
```

### Instalação manual (build local)

```
git clone https://github.com/lucasparente-codigos/arcana
cd arcana
go build -o arcana
```

---

## 🚀 Exemplos de Uso

### Gerar uma senha aleatória padrão

```
arcana generate
```

### Usar um perfil específico

```
arcana generate --profile web
```

### Gerar senha determinística

```
arcana generate --master "MinhaFraseMestra" --context "github-prod"
```

### Ativar modo stealth (QR code)

```
arcana generate --stealth
```

### Explicar como a senha foi construída

```
arcana generate --explain
```

### Limpar o terminal após X segundos

```
arcana generate --clean 10
```

---

## 🧩 Flags Disponíveis

| Flag                 | Descrição                                   |
| -------------------- | ------------------------------------------- |
| `--profile <nome>`   | Seleciona um dos perfis inteligentes.       |
| `--master <frase>`   | Frase-mestra para modo determinístico.      |
| `--context <valor>`  | Contexto para geração determinística.       |
| `--stealth`          | Exibe a senha como QR code.                 |
| `--clean <segundos>` | Limpa a saída automaticamente após o tempo. |
| `--explain`          | Mostra detalhes auditáveis da geração.      |
| `--length <n>`       | Define tamanho customizado da senha.        |

---

## 🔐 Segurança & Auditoria

Arcana não armazena frases-mestras, contextos ou senhas.
O modo determinístico utiliza combinações de hashing criptográfico e parâmetros auditáveis via `--explain`.
O QR code é gerado localmente, sem chamadas externas.

---

## 🗺️ Roadmap

* Suporte a plugins externos opcionais
* Integração com gerenciadores baseados em pipe (`pass`, `gopass`)
* Perfis adicionais para ambientes de compliance rígida
* Modo de geração em lote

---

## 📄 Licença

Disponível sob licença MIT. Para detalhes, consulte o arquivo LICENSE no repositório.
