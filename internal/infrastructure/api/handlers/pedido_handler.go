package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rayelissonl/elevator-purchase-api/internal/application/usecases"
	"github.com/rayelissonl/elevator-purchase-api/internal/domain/pedido"

	_ "github.com/rayelissonl/elevator-purchase-api/docs"
)

type PedidoHandler struct {
	CreateUC *usecases.CreatePedidoUseCase
	ListUC   *usecases.ListPedidosUseCase
	FindUC   *usecases.FindPedidoByIDUseCase
	UpdateUC *usecases.UpdatePedidoStatusUseCase
	DeleteUC *usecases.DeletePedidoUseCase
}

func NewPedidoHandler(
	createUC *usecases.CreatePedidoUseCase,
	listUC *usecases.ListPedidosUseCase,
	findUC *usecases.FindPedidoByIDUseCase,
	updateUC *usecases.UpdatePedidoStatusUseCase,
	deleteUC *usecases.DeletePedidoUseCase,
) *PedidoHandler {
	return &PedidoHandler{
		CreateUC: createUC,
		ListUC:   listUC,
		FindUC:   findUC,
		UpdateUC: updateUC,
		DeleteUC: deleteUC,
	}
}

// CreatePedido godoc
// @Summary Cria um novo pedido
// @Description Cria um novo pedido no sistema
// @Tags pedidos
// @Accept json
// @Produce json
// @Param pedido body pedido.Pedido true "Dados do Pedido"
// @Success 201 {object} pedido.Pedido
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
//
// @Router /pedidos [post]
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

// ListPedidosHandler godoc
// @Summary      Listar Pedidos
// @Description  Lista todos os pedidos cadastrados
// @Tags         pedidos
// @Accept       json
// @Produce      json
// @Success      200  {array}  pedido.Pedido
// @Router       /pedidos [get]
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

// DeletePedido godoc
// @Summary Deleta um pedido
// @Description Deleta um pedido pelo ID
// @Tags pedidos
// @Accept json
// @Produce json
// @Param id path string true "ID do Pedido"
// @Success 204 {string} string "Pedido deletado com sucesso"
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /pedidos/{id} [delete]
func (h *PedidoHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.DeleteUC.Execute(id) // Vamos assumir que você criou esse usecase
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}
	c.Status(http.StatusNoContent) // 204 - sucesso sem conteúdo
}

// UpdatePedido godoc
// @Summary Atualiza um pedido
// @Description Atualiza os dados de um pedido existente
// @Tags pedidos
// @Accept json
// @Produce json
// @Param id path string true "ID do Pedido"
// @Param pedido body pedido.Pedido true "Dados atualizados do Pedido"
// @Success 200 {object} pedido.Pedido
// @Failure 400 {object} handlers.ErrorResponse
// @Failure 404 {object} handlers.ErrorResponse
// @Failure 500 {object} handlers.ErrorResponse
// @Router /pedidos/{id} [put]
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
