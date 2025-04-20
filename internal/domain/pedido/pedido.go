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

type Pedido struct {
	ID                 string `gorm:"primaryKey"`
	NomePeca           string
	Fornecedor         string
	Cliente            string
	Foto               string
	Valor              float64
	LocalArmazenamento string
	Documentos         pq.StringArray `gorm:"type:text[]"`
	DataSolicitacao    time.Time
	DataAprovacao      *time.Time
	DataEntrega        *time.Time
	Status             Status
}
