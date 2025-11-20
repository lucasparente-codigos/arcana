#!/bin/bash

# Nome do binário
BINARY_NAME="arcana"
OUTPUT_DIR="bin"

echo "📦 Baixando dependências..."
go mod tidy

echo "📂 Preparando diretório de saída..."
mkdir -p $OUTPUT_DIR

echo "🔨 Compilando Arcana..."
# Flags de segurança e otimização:
# -s -w: Remove tabela de símbolos e debug info (menor binário, dificulta engenharia reversa)
# -trimpath: Remove caminhos absolutos do sistema de arquivos do build (privacidade)
if go build -ldflags="-s -w" -trimpath -o $OUTPUT_DIR/$BINARY_NAME main.go; then
    chmod +x $OUTPUT_DIR/$BINARY_NAME
    echo "✅ Build concluído com sucesso!"
    echo "👉 Executável: ./$OUTPUT_DIR/$BINARY_NAME"
    echo "💡 Teste agora: ./$OUTPUT_DIR/$BINARY_NAME --help"
else
    echo "❌ Erro fatal durante a compilação."
    exit 1
fi
