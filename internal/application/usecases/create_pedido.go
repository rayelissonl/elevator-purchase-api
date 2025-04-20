// internal/application/usecases/create_pedido.go
package usecases

import (
	"time"

	"github.com/google/uuid"

	"github.com/rayelissonl/elevator-purchase-api/internal/domain/pedido"
)

type CreatePedidoUseCase struct {
	Repo pedido.Repository
}

func (uc *CreatePedidoUseCase) Execute(input pedido.Pedido) error {
	input.ID = uuid.New().String()
	input.Status = pedido.Cotado
	input.DataSolicitacao = time.Now()
	return uc.Repo.Create(&input)
}

// ListPedidosUseCase - lista todos os pedidos
type ListPedidosUseCase struct {
	Repo pedido.Repository
}

func (uc *ListPedidosUseCase) Execute() ([]pedido.Pedido, error) {
	return uc.Repo.FindAll()
}

// FindPedidoByIDUseCase - busca pedido por ID
type FindPedidoByIDUseCase struct {
	Repo pedido.Repository
}

func (uc *FindPedidoByIDUseCase) Execute(id string) (*pedido.Pedido, error) {
	return uc.Repo.FindByID(id)
}

// UpdatePedidoStatusUseCase - atualiza o status de um pedido
type UpdatePedidoStatusUseCase struct {
	Repo pedido.Repository
}

func (uc *UpdatePedidoStatusUseCase) Execute(id string, status pedido.Status) error {
	return uc.Repo.UpdateStatus(id, status)
}
