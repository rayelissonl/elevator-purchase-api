package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rayelissonl/elevator-purchase-api/internal/application/usecases"
	"github.com/rayelissonl/elevator-purchase-api/internal/domain/pedido"
	"github.com/rayelissonl/elevator-purchase-api/internal/infrastructure/api/handlers"
	"github.com/rayelissonl/elevator-purchase-api/internal/infrastructure/api/routes"
	"github.com/rayelissonl/elevator-purchase-api/internal/infrastructure/persistence"

	_ "github.com/rayelissonl/elevator-purchase-api/docs"
)

func main() {
	// @title           Elevator Purchase API
	// @version         1.0
	// @description     API para gestão de pedidos de compra para manutenção de elevadores.
	// @termsOfService  http://swagger.io/terms/

	// @contact.name   Suporte
	// @contact.email  suporte@elevator.com

	// @host      localhost:8080
	// @BasePath  /api

	// Conectar ao banco
	dsn := "host=localhost user=postgres password=postgres dbname=erp_pedidos_compra_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	if err := db.AutoMigrate(&pedido.Pedido{}); err != nil {
		log.Fatalf("Erro ao realizar a migração: %v", err)
	}

	// Injeta repositórios e use cases
	pedidoRepo := persistence.NewPedidoRepository(db)
	createUC := &usecases.CreatePedidoUseCase{Repo: pedidoRepo}
	listUC := &usecases.ListPedidosUseCase{Repo: pedidoRepo}
	findUC := &usecases.FindPedidoByIDUseCase{Repo: pedidoRepo}
	updateUC := &usecases.UpdatePedidoStatusUseCase{Repo: pedidoRepo}
	deleteUC := &usecases.DeletePedidoUseCase{Repo: pedidoRepo}

	// Injeta os handlers
	pedidoHandler := handlers.NewPedidoHandler(createUC, listUC, findUC, updateUC, deleteUC)

	// Inicia o Gin
	r := gin.Default()

	// Registra as rotas separadamente
	routes.Setup(r, pedidoHandler)

	// Sobe o servidor
	r.Run(":8080")
}
