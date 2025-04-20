// internal/domain/pedido/repository.go
package pedido

type Repository interface {
	Create(pedido *Pedido) error
	FindByID(id string) (*Pedido, error)
	FindAll() ([]Pedido, error)
	UpdateStatus(id string, status Status) error
}
