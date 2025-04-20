package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rayelissonl/elevator-purchase-api/internal/application/usecases"
	"github.com/rayelissonl/elevator-purchase-api/internal/domain/pedido"
)

type PedidoHandler struct {
	CreateUC *usecases.CreatePedidoUseCase
	ListUC   *usecases.ListPedidosUseCase
	FindUC   *usecases.FindPedidoByIDUseCase
	UpdateUC *usecases.UpdatePedidoStatusUseCase
}

func NewPedidoHandler(
	createUC *usecases.CreatePedidoUseCase,
	listUC *usecases.ListPedidosUseCase,
	findUC *usecases.FindPedidoByIDUseCase,
	updateUC *usecases.UpdatePedidoStatusUseCase,
) *PedidoHandler {
	return &PedidoHandler{
		CreateUC: createUC,
		ListUC:   listUC,
		FindUC:   findUC,
		UpdateUC: updateUC,
	}
}

func (h *PedidoHandler) Create(c *gin.Context) {
	var input pedido.Pedido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.CreateUC.Execute(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pedido"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Pedido criado com sucesso!"})
}

func (h *PedidoHandler) List(c *gin.Context) {
	pedidos, err := h.ListUC.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pedidos)
}

func (h *PedidoHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	pedido, err := h.FindUC.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}
	c.JSON(http.StatusOK, pedido)
}

func (h *PedidoHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Status pedido.Status `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UpdateUC.Execute(id, body.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado com sucesso"})
}
