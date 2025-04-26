package pedido

import (
	"time"

	"github.com/lib/pq"
)

type Status string

const (
	Cotado           Status = "cotado"
	FazendoOrcamento Status = "fazendo_orcamento"
	Orcamento        Status = "orcamento"
	CompraAprovada   Status = "compra_aprovada"
	Comprado         Status = "comprado"
)

// Pedido representa um pedido de compra
// swagger:model Pedido
type Pedido struct {
	ID                 string  `gorm:"primaryKey"`
	NomePeca           string  `json:"nomePeca"`
	Fornecedor         string  `json:"fornecedor"`
	Cliente            string  `json:"cliente"`
	Foto               string  `json:"foto"`
	Valor              float64 `json:"valor"`
	LocalArmazenamento string  `json:"localArmazenamento"`
	//Documentos         []string `gorm:"type:text[]" json:"documentos"`
	Documentos      pq.StringArray `gorm:"type:text[]" json:"documentos"`
	DataSolicitacao time.Time      `json:"dataSolicitacao"`
	DataAprovacao   *time.Time     `json:"dataAprovacao,omitempty"`
	DataEntrega     *time.Time     `json:"dataEntrega,omitempty"`
	Status          Status         `json:"status"`
}
