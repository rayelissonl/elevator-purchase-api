package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/rayelissonl/elevator-purchase-api/internal/infrastructure/api/handlers"
)

func Setup(router *gin.Engine, pedidoHandler *handlers.PedidoHandler) {
	api := router.Group("/api")
	{
		// Rota para criação de pedido
		api.POST("/pedidos", pedidoHandler.Create)

		// Rota para listar todos os pedidos
		api.GET("/pedidos", pedidoHandler.List)

		// Rota para buscar pedido por ID
		api.GET("/pedidos/:id", pedidoHandler.FindByID)

		// Rota para atualizar o status do pedido
		api.PUT("/pedidos/:id/status", pedidoHandler.UpdateStatus)
	}
}
