<div align="center">

# 🦅 KESTREL

**Plataforma de Avaliação de Segurança & Testes de Penetração**

*Observe. Analise. Valide.*

**Mais do que um scanner — uma plataforma completa de Engenharia de Segurança.**

[![Status](https://img.shields.io/badge/status-desenvolvimento%20ativo-orange)]()
[![Go](https://img.shields.io/badge/core-Go-00ADD8?logo=go)]()
[![PostgreSQL](https://img.shields.io/badge/database-PostgreSQL-336791?logo=postgresql)]()
[![Python](https://img.shields.io/badge/automation-Python-3776AB?logo=python)]()
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Roadmap](https://img.shields.io/badge/roadmap-v2.0-blueviolet)]()

</div>

---

## ⚠️ Uso autorizado exclusivamente

O Kestrel foi desenvolvido para ser utilizado **exclusivamente** contra alvos que você possui ou para os quais tenha **permissão explícita por escrito** para realizar testes — incluindo laboratórios autorizados, ambientes CTF, programas de Bug Bounty e infraestrutura própria.

Autorização e escopo são tratados como **controles de segurança fundamentais**, e não como configurações opcionais.

O Kestrel foi projetado com:

* Autorização explícita de alvos
* Validação e aplicação de escopo
* Regras de permissão e negação
* Proteção contra operações fora de escopo
* Controle de acesso baseado em funções
* Rastreabilidade das atividades de avaliação

Consulte [`SECURITY.md`](SECURITY.md) para conhecer a política de segurança completa e o processo de divulgação responsável.

---

## 🎯 Por que o Kestrel?

A maioria dos projetos de Pentest para portfólio consiste em um único script que executa uma varredura e imprime os resultados.

O Kestrel está sendo desenvolvido como uma **plataforma real de avaliação de segurança**, com um ciclo estruturado desde a autorização até a geração do relatório.

A plataforma é construída em torno de:

* **Avaliação orientada por escopo**, com limites de autorização aplicados
* **Detecção ≠ Validação ≠ Exploração**
* Inventário persistente da **superfície de ataque**
* Evidências associadas aos achados
* Avaliações de segurança reproduzíveis
* Rastreabilidade das atividades
* Arquitetura modular de segurança
* Geração automatizada de relatórios
* Integração futura com DevSecOps e ambientes de nuvem

O objetivo não é simplesmente encontrar algo suspeito.

O objetivo é responder:

> **O que foi observado, por que isso importa, pode ser validado, quais evidências sustentam o achado e como ele deve ser corrigido?**

### Ciclo de avaliação

```text
Autorização
      │
      ▼
Escopo & Alvo
      │
      ▼
Descoberta
      │
      ▼
Enumeração
      │
      ▼
Detecção
      │
      ▼
Possível Achado
      │
      ▼
Validação de Segurança
      │
 ┌────┴────┐
 ▼         ▼
Rejeitado  Confirmado
              │
              ▼
           Evidência
              │
              ▼
            Impacto
              │
              ▼
           Relatório
              │
              ▼
         Remediação
              │
              ▼
           Reteste
```

> **Detecção ≠ Validação ≠ Exploração**

Uma detecção representa uma observação ou possível vulnerabilidade.

A validação determina se essa observação é realmente reproduzível e relevante.

A exploração, quando aplicável e explicitamente autorizada, constitui uma atividade de segurança separada.

---

## 📌 Status

🚧 **Desenvolvimento ativo**

O Kestrel concluiu a base do projeto até a etapa atual de **Enumeração & Inventário da Superfície de Ataque**.

### Progresso atual

* ✅ Fase 00 — Planejamento
* ✅ Fase 01 — Fundação do Projeto
* ✅ Fase 02 — Arquitetura do Backend
* ✅ Fase 03 — Banco de Dados & Migrations
* ✅ Fase 04 — Autenticação & RBAC
* ✅ Fase 05 — Gerenciamento de Alvos & Aplicação de Escopo
* ✅ Fase 06 — Reconhecimento
* ✅ Fase 07 — Enumeração & Inventário da Superfície de Ataque
* 🚧 Fase 08 — Avaliação de Vulnerabilidades

O backend atualmente possui uma API autenticada, persistência com PostgreSQL, gerenciamento de alvos com escopo, reconhecimento, enumeração e inventário da superfície de ataque.

Consulte [`CHANGELOG.md`](CHANGELOG.md) para acompanhar o histórico detalhado do desenvolvimento.

---

# ✨ Funcionalidades

## Implementadas

| Categoria                  | Funcionalidade                                 | Status |
| -------------------------- | ---------------------------------------------- | :----: |
| **Autenticação**           | Tokens JWT de acesso e renovação               |    ✅   |
| **Autorização**            | RBAC com `admin` / `analyst` / `viewer`        |    ✅   |
| **Gerenciamento de Alvos** | Ciclo de vida de alvos autorizados             |    ✅   |
| **Aplicação de Escopo**    | Limites de avaliação baseados em autorização   |    ✅   |
| **Reconhecimento**         | Resolução DNS                                  |    ✅   |
| **Reconhecimento**         | Descoberta passiva de subdomínios via `crt.sh` |    ✅   |
| **Reconhecimento**         | Detecção de serviços HTTP                      |    ✅   |
| **Enumeração**             | Varredura concorrente de portas TCP            |    ✅   |
| **Enumeração**             | Coleta passiva de banners                      |    ✅   |
| **Superfície de Ataque**   | Agregação de ativos entre execuções            |    ✅   |
| **Persistência**           | Dados de avaliação armazenados no PostgreSQL   |    ✅   |

## Em desenvolvimento

| Categoria                         | Funcionalidade                               | Status |
| --------------------------------- | -------------------------------------------- | :----: |
| **Avaliação de Vulnerabilidades** | Motor de detecção                            |   🚧   |
| **Severidade**                    | Avaliação baseada em CVSS                    |   🚧   |
| **Ciclo de Achados**              | Potencial → Validação → Confirmado/Rejeitado |   🚧   |
| **Evidências**                    | Coleta e persistência de evidências          |    ⏳   |
| **Relatórios**                    | Geração automatizada de relatórios           |    ⏳   |

## Planejadas

| Categoria           | Funcionalidade                           |
| ------------------- | ---------------------------------------- |
| **Relatórios**      | HTML / PDF / JSON                        |
| **Dashboard**       | Interface web com React + Next.js        |
| **Automação**       | Automação de segurança com Python        |
| **DevSecOps**       | SAST / SCA / Secret Scanning / SBOM      |
| **Infraestrutura**  | Docker / Kubernetes / Terraform          |
| **Cloud**           | Implantação na AWS                       |
| **Observabilidade** | Logs / métricas / traces                 |
| **Testes**          | Regressão de segurança / Fuzzing / E2E   |
| **Documentação**    | OpenAPI / ADRs / Runbooks / Threat Model |

---

# 📚 Padrões & Referências

A metodologia de avaliação de segurança do Kestrel é construída sobre padrões e referências consolidados da indústria.

| Padrão                        | Finalidade                                       |
| ----------------------------- | ------------------------------------------------ |
| **OWASP Top 10**              | Principais riscos de segurança em aplicações web |
| **OWASP API Security Top 10** | Riscos específicos de segurança em APIs          |
| **OWASP WSTG**                | Metodologia para testes de segurança web         |
| **CWE**                       | Classificação de fraquezas de software           |
| **CVSS**                      | Avaliação da severidade de vulnerabilidades      |
| **PTES**                      | Metodologia para testes de penetração            |

Esses padrões fornecem estrutura e terminologia para a plataforma, mas não substituem a análise profissional durante uma avaliação de segurança.

---

## 🔎 Taxonomia de vulnerabilidades

O motor de vulnerabilidades do Kestrel foi projetado para associar achados a classificações de segurança estabelecidas.

Exemplos:

| Categoria                              | Exemplos                                           |
| -------------------------------------- | -------------------------------------------------- |
| **Controle de Acesso Quebrado**        | IDOR, BOLA, BFLA                                   |
| **Configuração Insegura**              | Headers, CORS, arquivos expostos                   |
| **Segurança da Cadeia de Suprimentos** | Dependências vulneráveis, falhas em CI/CD          |
| **Falhas Criptográficas**              | TLS fraco, hashing inseguro, secrets expostos      |
| **Injeção**                            | XSS, SQL Injection, Command Injection, LFI         |
| **Design Inseguro**                    | Falhas de lógica de negócio, condições de corrida  |
| **Falhas de Autenticação**             | Gerenciamento de sessão, controles de autenticação |
| **Falhas de Integridade**              | Desserialização insegura, validação de integridade |
| **Logs & Monitoramento**               | Logging insuficiente, ausência de alertas          |
| **Condições Excepcionais**             | Tratamento inadequado de erros e resiliência       |

---

# 🏗️ Arquitetura

O Kestrel utiliza uma arquitetura de **monólito modular**.

Em vez de introduzir microsserviços prematuramente, a plataforma utiliza um único serviço Go implantável, com domínios internos claramente separados.

```text
                         ┌──────────────────────┐
                         │       Cliente        │
                         │ CLI / Web / API      │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │       Handler        │
                         │      HTTP / API      │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │       Service        │
                         │    Regra de Negócio  │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │       Domain         │
                         │  Conceitos de Segurança │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │     Repository       │
                         │   Persistência       │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │      PostgreSQL      │
                         └──────────────────────┘
```

### Domínios internos

```text
internal/
├── auth/
├── project/
├── target/
├── scan/
├── asset/
├── recon/
├── enum/
├── finding/
├── evidence/
└── report/
```

A arquitetura separa intencionalmente as responsabilidades de segurança para permitir que novos recursos sejam adicionados sem transformar o Kestrel em uma coleção de scripts de varredura fortemente acoplados.

As decisões arquiteturais e os diagramas completos estão documentados em [`ARCHITECTURE.md`](ARCHITECTURE.md).

---

# 🧩 Arquitetura de Segurança

A autorização faz parte do próprio fluxo de avaliação.

Um alvo não deve automaticamente se tornar elegível para testes apenas porque foi cadastrado no banco de dados.

```text
┌────────────────────┐
│   Alvo Criado      │
└─────────┬──────────┘
          │
          ▼
┌────────────────────┐
│   Não Autorizado   │
└─────────┬──────────┘
          │
          │ Aprovação explícita
          ▼
┌────────────────────┐
│     Autorizado     │
└─────────┬──────────┘
          │
          ▼
┌────────────────────┐
│ Avaliação Permitida│
└────────────────────┘
```

O modelo de segurança é estruturado em múltiplas camadas:

```text
Autenticação
      │
      ▼
Autorização
      │
      ▼
Escopo do Alvo
      │
      ▼
Permissão de Avaliação
      │
      ▼
Operação de Segurança
      │
      ▼
Auditoria / Evidência
```

Esse modelo tem como objetivo reduzir atividades acidentais fora de escopo e fornecer rastreabilidade durante todo o ciclo de avaliação.

---

# 🗺️ Roadmap — v2.0

<details open>
<summary><strong>Bloco 1 · Fundação</strong></summary>

| Fase | Foco                                                        | Status |
| ---- | ----------------------------------------------------------- | :----: |
| 00   | Planejamento — objetivos, escopo, threat model, arquitetura |    ✅   |
| 01   | Fundação do projeto — repositório, estrutura, `/health`     |    ✅   |

</details>

<details open>
<summary><strong>Bloco 2 · Núcleo da Plataforma</strong></summary>

| Fase | Foco                                                                  | Status |
| ---- | --------------------------------------------------------------------- | :----: |
| 02   | Arquitetura do backend — handler → service → domain → repository      |    ✅   |
| 03   | Banco de dados & migrations — PostgreSQL, entidades e relacionamentos |    ✅   |
| 04   | Autenticação — JWT, refresh tokens, hashing, RBAC                     |    ✅   |
| 05   | Gerenciamento de alvos — autorização e aplicação de escopo            |    ✅   |

</details>

<details open>
<summary><strong>Bloco 3 · Motor de Segurança</strong></summary>

| Fase | Foco                                                               | Status |
| ---- | ------------------------------------------------------------------ | :----: |
| 06   | Reconhecimento — DNS, subdomínios e descoberta HTTP                |    ✅   |
| 07   | Enumeração — portas, serviços e inventário da superfície de ataque |    ✅   |
| 08   | Avaliação de vulnerabilidades — OWASP, CWE, CVSS                   |   🚧   |
| 09   | Validação de segurança — potencial → confirmado/rejeitado          |    ⏳   |
| 10   | Relatórios — metodologia, achados, evidências e remediação         |    ⏳   |

</details>

<details>
<summary><strong>Bloco 4 · Interface & Automação</strong></summary>

| Fase | Foco                              | Status |
| ---- | --------------------------------- | :----: |
| 11   | Dashboard web                     |    ⏳   |
| 12   | Automação de segurança com Python |    ⏳   |

</details>

<details>
<summary><strong>Bloco 5 · Entrega & Escala</strong></summary>

| Fase | Foco                                                           | Status |
| ---- | -------------------------------------------------------------- | :----: |
| 13   | Docker — imagens de produção, non-root e health checks         |    ⏳   |
| 14   | DevSecOps — SAST, SCA, secrets, SBOM e provenance              |    ⏳   |
| 15   | Kubernetes — deployment, networking, policies e scaling        |    ⏳   |
| 16   | AWS — VPC, EKS, RDS, S3, IAM e CloudWatch                      |    ⏳   |
| 17   | Terraform — infraestrutura como código                         |    ⏳   |
| 18   | Observabilidade — logs, métricas e traces                      |    ⏳   |
| 19   | Hardening — aplicação, API, banco e containers                 |    ⏳   |
| 20   | Testes — unitários, integração, E2E, regressão e fuzzing       |    ⏳   |
| 21   | Documentação — API, ADRs, threat model e runbooks              |    ⏳   |
| 22   | Deploy semelhante a produção — CI/CD, monitoramento e rollback |    ⏳   |
| 23   | Portfólio — releases, demo, screenshots e case studies         |    ⏳   |

</details>

---

# 🧰 Stack Tecnológica

| Camada                     | Tecnologia                   | Finalidade                                                  |
| -------------------------- | ---------------------------- | ----------------------------------------------------------- |
| **Core API**               | Go                           | Plataforma principal, concorrência e workflows de segurança |
| **HTTP Router**            | chi                          | Roteamento HTTP e middleware                                |
| **Banco de Dados**         | PostgreSQL                   | Persistência dos dados da plataforma                        |
| **Driver**                 | pgx                          | Acesso ao PostgreSQL em Go                                  |
| **Autenticação**           | JWT                          | Autenticação por tokens de acesso e renovação               |
| **Password Hashing**       | bcrypt                       | Proteção de credenciais                                     |
| **Migrations**             | golang-migrate               | Versionamento do schema                                     |
| **Automação de Segurança** | Python                       | Automação e ferramentas de segurança                        |
| **Frontend**               | React + Next.js + TypeScript | Dashboard web planejado                                     |
| **Containers**             | Docker / Docker Compose      | Ambientes reproduzíveis                                     |
| **CI/CD**                  | GitHub Actions               | Pipeline de segurança e entrega                             |
| **Cache / Queue**          | Redis                        | Planejado quando necessário para orquestração               |
| **Orquestração**           | Kubernetes                   | Orquestração de containers                                  |
| **Infraestrutura**         | Terraform                    | Infraestrutura como código                                  |
| **Cloud**                  | AWS                          | Implantação em nuvem                                        |
| **Observabilidade**        | Prometheus / Grafana         | Monitoramento planejado                                     |

> O Kestrel evita introduzir complexidade de infraestrutura antes que ela seja realmente necessária. Redis, Kubernetes, AWS e Terraform entram posteriormente no roadmap conforme os requisitos da plataforma evoluem.

---

# 🔬 Inventário da Superfície de Ataque

Um dos conceitos centrais do Kestrel é manter uma visão agregada dos ativos descobertos durante as avaliações de segurança.

```text
Alvo
 │
 ├── Domínios
 │    ├── example.com
 │    ├── api.example.com
 │    └── dev.example.com
 │
 ├── Endereços IP
 │
 ├── Serviços
 │    ├── HTTP
 │    ├── HTTPS
 │    └── SSH
 │
 ├── Portas
 │    ├── 22
 │    ├── 80
 │    └── 443
 │
 └── Tecnologias
      ├── Web Server
      ├── Framework
      └── Aplicação
```

Em vez de tratar cada scan como uma execução isolada, o Kestrel foi projetado para construir uma visão persistente da superfície de ataque observável do alvo ao longo do tempo.

---

# 🚀 Como começar

## Requisitos

* Go 1.21+
* Docker
* Docker Compose
* PostgreSQL
* [golang-migrate](https://github.com/golang-migrate/migrate)

## Clonar o repositório

```bash
git clone https://github.com/leomarqueseh/kestrel.git
cd kestrel
```

## Configurar o ambiente

```bash
cp .env.example .env
```

Configure as variáveis de ambiente necessárias de acordo com seu ambiente local.

## Iniciar o PostgreSQL

```bash
docker compose up -d
```

## Configurar a conexão com o banco

```bash
export DATABASE_URL="postgres://kestrel:kestrel@localhost:5432/kestrel?sslmode=disable"
```

## Executar as migrations

```bash
migrate \
  -database "$DATABASE_URL" \
  -path migrations \
  up
```

## Instalar dependências

```bash
go mod tidy
```

## Criar o administrador inicial

```bash
go run ./cmd/seed \
  -email admin@kestrel.local \
  -password "change-me"
```

> **Somente para desenvolvimento:** nunca utilize credenciais de exemplo em ambientes de produção.

## Iniciar o Kestrel

```bash
make run
```

A API estará disponível em:

```text
http://localhost:8080
```

### Health check

```bash
curl http://localhost:8080/health
```

Resposta esperada:

```json
{
  "status": "ok"
}
```

---

# 📁 Estrutura do Projeto

```text
kestrel/
├── api/                 # Definições e contratos da API
├── cmd/                 # Entrypoints da aplicação
├── configs/             # Configurações
├── deployments/         # Manifestos de deploy
├── docs/                # Documentação técnica
├── internal/
│   ├── auth/            # Autenticação e autorização
│   ├── project/         # Gerenciamento de projetos
│   ├── target/          # Gerenciamento de alvos e escopo
│   ├── scan/            # Orquestração de scans
│   ├── asset/           # Ativos da superfície de ataque
│   ├── recon/           # Reconhecimento
│   └── enum/            # Enumeração
├── migrations/          # Migrations do banco
├── pkg/                 # Pacotes reutilizáveis
├── scripts/             # Scripts auxiliares
├── automation/          # Automação com Python
├── tests/               # Suítes de testes
├── Makefile
├── ARCHITECTURE.md
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── README.md
└── SECURITY.md
```

---

# 🔐 Filosofia de Testes de Segurança

O Kestrel parte de um princípio simples:

> **Um achado é mais do que a saída de um scanner.**

Uma avaliação madura deve estabelecer uma cadeia clara:

```text
Observação
    ↓
Contexto
    ↓
Validação
    ↓
Evidência
    ↓
Impacto
    ↓
Remediação
    ↓
Reteste
```

Essa abordagem busca reduzir falsos positivos, melhorar a qualidade dos relatórios e criar um registro de avaliação que possa ser compreendido tanto por profissionais de segurança quanto por stakeholders.

---

# 🤝 Contribuindo

O Kestrel segue atualmente um roadmap estruturado de desenvolvimento.

Contribuições, discussões técnicas, relatos de bugs e feedback arquitetural são bem-vindos.

Antes de contribuir:

1. Leia [`CONTRIBUTING.md`](CONTRIBUTING.md).
2. Consulte o roadmap atual.
3. Mantenha as alterações limitadas ao módulo relevante.
4. Adicione ou atualize testes quando necessário.
5. Documente comportamentos relacionados à segurança.
6. Nunca introduza funcionalidades que ignorem controles de autorização ou escopo.

---

# 📄 Licença

O Kestrel é distribuído sob a [Licença MIT](LICENSE).

---

# 🔒 Segurança

Encontrou uma vulnerabilidade no próprio Kestrel?

Consulte [`SECURITY.md`](SECURITY.md) para conhecer o processo de divulgação responsável.

Evite divulgar publicamente detalhes sensíveis de uma vulnerabilidade antes que ela seja reportada e avaliada de maneira responsável.

---

<div align="center">

### 🦅 Kestrel

**Observe. Analise. Valide.**

*Construído para um futuro mais seguro.*

</div>
