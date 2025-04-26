package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/rayelissonl/elevator-purchase-api/internal/infrastructure/api/handlers"
)

func Setup(router *gin.Engine, pedidoHandler *handlers.PedidoHandler) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		// Rota para criação de pedido
		api.POST("/pedidos", pedidoHandler.Create)

		api.GET("/pedidos", pedidoHandler.List)

		// Rota para buscar pedido por ID
		api.GET("/pedidos/:id", pedidoHandler.FindByID)

		// Rota para atualizar o status do pedido
		api.PUT("/pedidos/:id/status", pedidoHandler.UpdateStatus)

		api.DELETE("/pedidos/:id", pedidoHandler.Delete)

		//api.PUT("/pedidos/:id", pedidoHandler.Update)
	}
}
