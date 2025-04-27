# 🎯 API REST - Sistema de Monitoramento de Pedidos de Compra

Este projeto fornece uma API RESTful desenvolvida com Golang (Gin) para o gerenciamento e monitoramento de pedidos de compra. A API centraliza dados de pedidos, peças, fornecedores, clientes e serviços, além de se integrar com ferramentas de mensageria e monitoramento.

## 🚀 Projeto | 💻 Tecnologias

| Camada             | Projeto / Módulo       | Tecnologias Principais                                               |
| ------------------ | ---------------------- | -------------------------------------------------------------------- |
| **Backend**        | API REST               | 🟨 Golang (Gin) • 🧠 Arquitetura DDD • ☁️ AWS                        |
| **Banco de Dados** | Persistência           | 🐘 PostgreSQL • GORM                                                 |
| **Documentação**   | API RESTful            | 📄 Swagger / OpenAPI                                                 |
| **Mensageria**     | Comunicação Assíncrona | 🐇 RabbitMQ • 📬 Event-Driven Architecture                           |
| **Monitoramento**  | Observabilidade        | 📊 Grafana • 📈 Prometheus                                           |
| **DevOps**         | CI/CD + Orquestração   | 🐳 Docker + Docker Compose • ☸️ Kubernetes (EKS) • 🧪 GitHub Actions |

## 📂 Funcionalidades da API

- **Cadastro de Pedidos de Compra**

  - Nome da peça
  - Fornecedor
  - Cliente
  - Foto da peça
  - Valor
  - Local de armazenamento
  - Documentos anexos
  - Datas importantes (solicitação, aprovação e entrega)

- **Pesquisa e Consulta**

  - Pesquisa avançada por:
    - Nome da peça
    - Fornecedor
    - Cliente
    - Status do pedido

- **Acompanhamento do Processo de Compra**

  - Status do pedido:
    - **Cotado**
    - **Fazendo orçamento**
    - **Orçamento**
    - **Compra aprovada**
    - **Comprado**
  - Notificações automáticas
  - Registro de histórico

- **Visualização em Kanban**

  - Interface de visualização no formato **Kanban**
  - Colunas baseadas nos status do processo:
    - Cotado
    - Fazendo Orçamento
    - Orçamento
    - Compra Aprovada
    - Comprado
  - Arrastar e soltar (drag & drop) para atualizar o status do pedido.
  - Cards com informações resumidas (nome da peça, fornecedor, valor, datas).
  - Filtragem de cards por cliente, fornecedor e prioridade.

- **Gestão de Contratos e Serviços**

  - Vínculo de pedidos a contratos
  - Registro do tipo de serviço:
    - Manutenção
    - Venda de peças
    - Outros

- **Controle de Estoque**

  - Atualização automática
  - Notificações para estoques baixos
  - Reserva de peças para serviços

- **Gestão Financeira**

  - Registro de custos e previsões
  - Integração com contas a pagar
  - Relatórios financeiros detalhados

- **Relatórios e Dashboards**
  - Análise de pedidos
  - Indicadores de eficiência
  - Desempenho de fornecedores

## 🧪 Demonstração da API

Acesse a documentação interativa da API:
