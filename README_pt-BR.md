<div align="center">

# 🦅 KESTREL

**Plataforma de Avaliação de Segurança e Testes de Invasão**

*Observe. Analise. Valide.*

Mais que um scanner — uma plataforma completa de Security Engineering.

[![Status](https://img.shields.io/badge/status-em%20desenvolvimento-orange)]()
[![Go](https://img.shields.io/badge/core-Go-00ADD8?logo=go)]()
[![Python](https://img.shields.io/badge/automação-Python-3776AB?logo=python)]()
[![License: MIT](https://img.shields.io/badge/licença-MIT-blue.svg)](LICENSE)
[![Roadmap](https://img.shields.io/badge/roadmap-v2.0-blueviolet)]()

🇧🇷 Português | [🇺🇸 English](README.md)

</div>

---

## ⚠️ Uso somente autorizado

O Kestrel é construído para ser usado **exclusivamente** contra alvos que você possui ou para os quais tem **permissão explícita por escrito** para testar — labs, CTFs, programas de bug bounty ou sua própria infraestrutura. Validação de escopo, listas de permissão/bloqueio e proteção contra alvos fora do escopo são funcionalidades centrais da plataforma, não um adendo. Veja [`SECURITY.md`](SECURITY.md) para a política completa de escopo e o processo de divulgação responsável.

---

## 🎯 Por que o Kestrel

A maioria dos projetos de pentest para portfólio é um único script que roda um scan e imprime resultados. O Kestrel é construído como uma plataforma de verdade:

- Gerencia **escopo** com limites de autorização aplicados
- Separa achados **detectados** de achados **validados/confirmados**
- Mantém uma **trilha de auditoria** de evidências (requisições, respostas, timestamps, notas do analista)
- Produz um **relatório que um cliente de fato poderia receber** — não um dump de terminal

> "Técnica + Automação + Inteligência = Segurança Real"

## 📌 Status

🚧 **Em desenvolvimento inicial** — Fase 00 (Planejamento) concluída, Fase 01 (Fundação do Projeto) em andamento.
Veja [`CHANGELOG.md`](CHANGELOG.md) para o progresso detalhado.

---

## ✨ Funcionalidades (planejadas)

| Categoria | Capacidade |
|---|---|
| **Gestão de Alvos** | Domínios, IPs, URLs e ranges CIDR autorizados, com listas de permissão/bloqueio |
| **Reconhecimento** | DNS, enumeração de subdomínios, descoberta de serviços HTTP e certificados |
| **Enumeração** | Portas, serviços, versões, tecnologias → inventário completo da superfície de ataque |
| **Avaliação de Vulnerabilidades** | Detecção alinhada a OWASP Top 10 / API Top 10 / WSTG, com pontuação CVSS |
| **Fluxo de Validação** | `Achado Potencial → Validação → Confirmado / Falso Positivo`, com rastreio de confiança |
| **Coleta de Evidências** | Captura de requisição/resposta, timestamps e notas do analista para cada achado confirmado |
| **Relatórios** | Sumário executivo, metodologia, achados, evidências e recomendações automatizados — HTML / PDF / JSON |
| **Controle de Acesso** | RBAC (admin, analista, visualizador), logs de auditoria, rate limiting |
| **Dashboard** | Frontend web para projetos, alvos, scans, ativos, achados e relatórios |

---

## 📚 Padrões e Referências

A metodologia de avaliação do Kestrel é construída diretamente sobre padrões da indústria, não sobre heurísticas improvisadas:

- **OWASP Top 10:2025** — Categorias de risco de aplicações web
- **OWASP API Security Top 10** — Categorias de risco específicas de APIs
- **OWASP WSTG** (Web Security Testing Guide) — Metodologia de testes
- **CWE** (Common Weakness Enumeration) — Classificação de fraquezas
- **CVSS** (Common Vulnerability Scoring System) — Pontuação de severidade

<details>
<summary><strong>Catálogo de vulnerabilidades (mapeamento OWASP Top 10:2025)</strong></summary>

| ID | Categoria | Exemplos |
|---|---|---|
| A01 | Broken Access Control | IDOR, BOLA, BFLA |
| A02 | Security Misconfiguration | Headers, CORS, exposição de `.git` |
| A03 | Supply Chain Failures | Dependências, CI/CD |
| A04 | Cryptographic Failures | TLS, hashing, secrets |
| A05 | Injection | XSS, SQLi, SSRF, LFI |
| A06 | Insecure Design | Lógica de negócio, condições de corrida |
| A07 | Authentication Failures | Força bruta, gestão de sessão |
| A08 | Data Integrity Failures | Deserialização, verificação de integridade |
| A09 | Logging & Alerting Failures | Logs, monitoramento |
| A10 | Exceptional Conditions | Tratamento de erros, DoS |

</details>

---

## 🏗️ Arquitetura

O Kestrel começa como um **monólito modular** — um único serviço Go implantável, com módulos internos claramente separados (Scanner Engine, Auth Service, Reports Service) — em vez de microsserviços desde o início. Racional completo e diagramas: [`ARCHITECTURE.md`](ARCHITECTURE.md).

A construção está organizada em cinco blocos:

| Bloco | Foco |
|---|---|
| **1 · Fundação** | Escopo do projeto, threat model, arquitetura, boilerplate do repositório |
| **2 · Núcleo da Plataforma** | Arquitetura de backend, banco de dados, autenticação, gestão de alvos |
| **3 · Motor de Segurança** | Reconhecimento → enumeração → avaliação de vulnerabilidades → validação → relatórios |
| **4 · Interface e Automação** | Dashboard frontend, automação de segurança em Python |
| **5 · Entrega / Escala** | Docker, DevSecOps, Kubernetes, AWS, Terraform, observabilidade, hardening, testes, documentação |

---

## 🧰 Stack técnica

| Camada | Escolha | Por quê |
|---|---|---|
| **API Principal** | Go (Chi/Gin) | Performance, tipagem forte, concorrência para orquestração de scans |
| **Automação de segurança** | Python | Scripts de recon/parsing onde o ecossistema é imbatível — nunca duplica responsabilidades do Go |
| **Banco de dados** | PostgreSQL | Integridade relacional para alvos, scans, achados, evidências |
| **Cache/Fila** | Redis | Adicionado quando a orquestração de scans realmente precisar dele (não desde o dia um) |
| **Frontend** | React + Next.js + TypeScript | Consome a API exclusivamente |
| **Containers** | Docker / Docker Compose | Paridade entre local e CI |
| **Orquestração** | Kubernetes | Introduzido quando houver motivo concreto para orquestrar em escala |
| **IaC** | Terraform | Infra como código — VPC, EKS, IAM, RDS, S3, load balancer |
| **Cloud** | AWS | VPC, EKS, RDS, ElastiCache, S3, IAM, CloudWatch |
| **CI/CD** | GitHub Actions | Lint → testes → SAST → dependency scan → secret scan → build → deploy |

Kubernetes, Terraform e AWS são introduzidos mais adiante no roadmap, quando houver motivo concreto para orquestrar e implantar nessa escala — veja [`ARCHITECTURE.md`](ARCHITECTURE.md).

---

## 🗺️ Roadmap (v2.0)

<details open>
<summary><strong>Bloco 1 · Fundação</strong></summary>

| Fase | Foco | Status |
|---|---|---|
| 00 | Planejamento — objetivo, escopo, threat model, arquitetura | ✅ |
| 01 | Fundação do projeto — estrutura do repo, boilerplate, endpoint `/health` | 🚧 |

</details>

<details>
<summary><strong>Bloco 2 · Núcleo da Plataforma</strong></summary>

| Fase | Foco |
|---|---|
| 02 | Arquitetura de backend (handler → service → domain → repository, SOLID, DI) |
| 03 | Banco de dados e migrações (PostgreSQL, entidades, relacionamentos, índices) |
| 04 | Autenticação (JWT, refresh tokens, hashing de senha, RBAC, audit logs) |
| 05 | Gestão de alvos (validação de escopo, allow/deny list, CIDR/domínio/IP/URL) |

</details>

<details>
<summary><strong>Bloco 3 · Motor de Segurança</strong></summary>

| Fase | Foco |
|---|---|
| 06 | Reconhecimento (DNS, subdomínios, HTTP, tech discovery, certificados) |
| 07 | Enumeração (portas, serviços, versões, inventário de superfície de ataque) |
| 08 | Avaliação de vulnerabilidades (OWASP Top 10:2025 / API Top 10 / WSTG, CWE, CVSS) |
| 09 | Validação de segurança (achado potencial → validação → confirmado/rejeitado) |
| 10 | Relatórios (sumário executivo, escopo, metodologia, achados, evidências) |

</details>

<details>
<summary><strong>Bloco 4 · Interface e Automação</strong></summary>

| Fase | Foco |
|---|---|
| 11 | Dashboard frontend (projetos, alvos, scans, ativos, achados, relatórios) |
| 12 | Automação de segurança em Python (recon, parsers, HTTP, utils de relatório) |

</details>

<details>
<summary><strong>Bloco 5 · Entrega / Escala</strong></summary>

| Fase | Foco |
|---|---|
| 13 | Docker (multi-stage, non-root, healthcheck, secrets) |
| 14 | DevSecOps (SAST, SCA, secret scan, SBOM, provenance) |
| 15 | Kubernetes (Deploy, Service, Ingress, ConfigMap/Secret, HPA, RBAC, network policy) |
| 16 | AWS (VPC, EKS, RDS, ElastiCache, S3, IAM, CloudWatch, load balancer) |
| 17 | Terraform (IaC, VPC, EKS, RDS, S3, load balancer) |
| 18 | Observabilidade (logs, métricas, traces, Prometheus, Grafana, dashboards) |
| 19 | Hardening de segurança (app/API/DB, Docker/K8s, threat modeling, TLS) |
| 20 | Testes (unitários, integração, API, E2E, regressão de segurança, fuzzing) |
| 21 | Documentação (arquitetura, API/OpenAPI, threat model, ADRs, runbooks) |
| 22 | Deploy em ambiente de produção (CI/CD, monitoramento, rollback, health checks, auto-scaling) |
| 23 | Portfólio (GitHub, demo, releases, screenshots, cases) |

</details>

---

## 🚀 Como começar

> Chegando na Fase 01 — a API ainda não existe.

```bash
# Placeholder — atualizado quando a Fase 01 for lançada
git clone https://github.com/kestrel/kestrel.git
cd kestrel
make setup
```

## 📁 Estrutura do projeto

```
kestrel/
├── cmd/                # Entrypoints
├── internal/
│   ├── handler/        # Camada HTTP
│   ├── service/         # Lógica de negócio
│   ├── domain/          # Entidades centrais
│   └── repository/      # Acesso a dados
├── automation/          # Automação de segurança em Python
├── web/                 # Frontend React + Next.js
├── deployments/         # Docker, Kubernetes, Terraform
└── docs/                 # Arquitetura, ADRs, runbooks
```

## 🤝 Contribuindo

Este projeto atualmente segue um roadmap de desenvolvimento solo (veja o [Roadmap](#️-roadmap-v20)). Issues e discussões são bem-vindas — veja [`CONTRIBUTING.md`](CONTRIBUTING.md) quando publicado.

## 📄 Licença

MIT — veja [`LICENSE`](LICENSE).

## 🔒 Segurança

Encontrou uma vulnerabilidade no próprio Kestrel, ou tem dúvidas sobre escopo autorizado? Veja [`SECURITY.md`](SECURITY.md) para o processo de divulgação responsável.

---

<div align="center">
<sub>Kestrel — Construído para um futuro mais seguro.</sub>
</div>
