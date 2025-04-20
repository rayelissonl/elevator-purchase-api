// internal/infrastructure/persistence/pedido_repository.go
package persistence

import (
	"gorm.io/gorm"

	"github.com/rayelissonl/elevator-purchase-api/internal/domain/pedido"
)

type PedidoGormRepo struct {
	DB *gorm.DB
}

func NewPedidoRepository(db *gorm.DB) pedido.Repository {
	return &PedidoGormRepo{DB: db}
}

func (r *PedidoGormRepo) Create(p *pedido.Pedido) error {
	return r.DB.Create(p).Error
}

func (r *PedidoGormRepo) FindByID(id string) (*pedido.Pedido, error) {
	var p pedido.Pedido
	err := r.DB.First(&p, "id = ?", id).Error
	return &p, err
}

func (r *PedidoGormRepo) FindAll() ([]pedido.Pedido, error) {
	var pedidos []pedido.Pedido
	err := r.DB.Find(&pedidos).Error
	return pedidos, err
}

func (r *PedidoGormRepo) UpdateStatus(id string, status pedido.Status) error {
	return r.DB.Model(&pedido.Pedido{}).Where("id = ?", id).Update("status", status).Error
}
