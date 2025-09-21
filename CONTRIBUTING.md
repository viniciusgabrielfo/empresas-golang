# Guia de Contribuição

Obrigado por contribuir com o repositório **Empresas que utilizam Golang**! Este guia irá te ajudar a adicionar novas empresas de forma padronizada e eficiente.

## 📋 Como Contribuir

### 1. Fork e Clone
1. Faça um fork deste repositório
2. Clone o seu fork localmente:
   ```bash
   git clone https://github.com/SEU_USUARIO/empresas-golang.git
   cd empresas-golang
   ```

### 2. Adicionando uma Nova Empresa

#### Informações Necessárias
Para adicionar uma empresa, você precisa das seguintes informações:

- **Nome da empresa**: Nome oficial da empresa
- **Website**: Link para o site oficial da empresa
- **Stackshare** (opcional): Link do perfil da empresa no [Stackshare](https://stackshare.io)
- **Link de vagas**: Link direto para a página de vagas da empresa
- **Tipo de contratação**: NACIONAL, INTERNACIONAL ou NACIONAL/INTERNACIONAL

#### Como Encontrar os Links

**Link de vagas:**
- Procure por páginas como: `/careers`, `/jobs`, `/trabalhe-conosco`, `/vagas`
- Se não encontrar página específica, use o LinkedIn da empresa: `https://www.linkedin.com/company/NOME_DA_EMPRESA/jobs`
- Para empresas internacionais, procure por `/careers` ou `/jobs`

**Stackshare:**
- Acesse [stackshare.io](https://stackshare.io)
- Busque pelo nome da empresa
- Se encontrar, copie o link do perfil

#### Formato da Entrada

Cada empresa deve seguir este formato na tabela:

```markdown
| [Nome da Empresa](https://website.com) | [Clique aqui](https://stackshare.io/empresa) | [Clique aqui](https://empresa.com/vagas) | TIPO_CONTRATACAO |
```

**Exemplo:**
```markdown
| [TechCorp](https://techcorp.com) | [Clique aqui](https://stackshare.io/techcorp) | [Clique aqui](https://techcorp.com/careers) | NACIONAL |
```

### 3. Ordenação Alfabética

⚠️ **IMPORTANTE**: As empresas devem estar em ordem alfabética por nome.

**Ordenação automática:**
```bash
make sort
```

**Dicas para ordenação manual:**
- Ignore artigos como "A", "O", "De", "Da", "Do"
- Use o nome principal da empresa (ex: "Mercado Livre" vem antes de "MadeiraMadeira")
- Empresas com nomes em inglês vêm antes das em português quando começam com a mesma letra

### 4. Validação Antes do Commit

Antes de fazer o commit, verifique:

- [ ] A empresa está em ordem alfabética (use `make sort` para ordenar automaticamente)
- [ ] Todos os links funcionam (teste clicando neles)
- [ ] O formato da tabela está correto
- [ ] O tipo de contratação está correto
- [ ] O nome da empresa está linkado para o website oficial

### 5. Criando o Pull Request

1. Faça commit das suas mudanças:
   ```bash
   git add README.md
   git commit -m "feat: Adiciona [Nome da Empresa] à lista de empresas"
   ```

2. Push para o seu fork:
   ```bash
   git push origin main
   ```

3. Crie um Pull Request no GitHub

4. Preencha o template do PR com as informações solicitadas

## 🛠️ Ferramentas do Projeto

### Ordenação Automática
O projeto inclui um script em Go para ordenar automaticamente as empresas:

```bash
# Ordenar empresas automaticamente
make sort

# Ver todos os comandos disponíveis
make help
```

## 🔍 Validações Automáticas

O repositório possui validações automáticas que verificam:

- ✅ Se todos os links estão funcionando
- ✅ Se a formatação da tabela está correta
- ✅ Se a empresa foi adicionada em ordem alfabética

## ❓ Dúvidas Frequentes

### "Não encontrei o link de vagas da empresa"
- Use o LinkedIn da empresa: `https://www.linkedin.com/company/NOME_DA_EMPRESA/jobs`
- Ou deixe em branco e mencione no PR que não foi encontrado

### "A empresa não tem perfil no Stackshare"
- Deixe a coluna em branco (não coloque "N/A" ou "-")

### "Como saber se a empresa contrata brasileiros?"
- Verifique se há vagas remotas ou escritórios no Brasil
- Consulte o LinkedIn ou site de carreiras da empresa
- Se for internacional, marque como "INTERNACIONAL"

### "A empresa tem escritório no Brasil mas também contrata internacionalmente"
- Use "NACIONAL/INTERNACIONAL"

## 📝 Template de Pull Request

```markdown
## Adiciona [Nome da Empresa]

### Informações da Empresa
- **Nome**: [Nome da Empresa]
- **Website**: [https://empresa.com]
- **Stackshare**: [https://stackshare.io/empresa] (se aplicável)
- **Vagas**: [https://empresa.com/careers]
- **Tipo**: [NACIONAL/INTERNACIONAL/NACIONAL/INTERNACIONAL]

### Verificações
- [ ] Empresa adicionada em ordem alfabética
- [ ] Todos os links testados e funcionando
- [ ] Formato da tabela correto
- [ ] Tipo de contratação correto

### Observações
[Adicione qualquer observação relevante]
```

---

Obrigado por contribuir! 🚀
