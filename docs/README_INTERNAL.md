# Documentação Interna - Arcana

## Estrutura do Projeto
Este documento serve para orientar desenvolvedores sobre a arquitetura interna.

- **core/**: Lógica pura. Sem prints, sem dependências de UI.
- **cli/**: Lógica de apresentação. Onde interage com o usuário.
- **presets/**: Dados estáticos e regras de negócio sobre formatos de senha.

## Roadmap de Desenvolvimento
1. Implementar lógica Core (crypto).
2. Criar Perfis básicos.
3. Conectar CLI ao Core.
4. Adicionar funcionalidades Stealth (QR).

