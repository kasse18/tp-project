package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tp-project/internal/models/tattoo"
)

type ServiceTattoo interface {
	GetAllTattoos(ctx context.Context) ([]models.Tattoo, error)
	GetTattooByID(ctx context.Context, id int) (*models.Tattoo, error)
	CreateTattoo(ctx context.Context, tattoo models.Tattoo) (int, error)
	UpdateTattoo(ctx context.Context, tattooId int, tattoo models.Tattoo) error
	DeleteTattoo(ctx context.Context, id int) error
}

type TattooHandler struct {
	service ServiceTattoo
}

func NewTattooHandler(tattooService ServiceTattoo) TattooHandler {
	return TattooHandler{
		service: tattooService,
	}
}

func (h TattooHandler) GetAllTattoos(g *gin.Context) {
	ctx := g.Request.Context()
	tattoos, err := h.service.GetAllTattoos(ctx)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"tattoos": tattoos})
}

func (h TattooHandler) CreateTattoo(g *gin.Context) {
	ctx := g.Request.Context()
	var tattoo models.Tattoo
	if err := g.BindJSON(&tattoo); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateTattoo(ctx, tattoo)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

func (h TattooHandler) GetTattooByID(g *gin.Context) {
	ctx := g.Request.Context()
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	tattoo, err := h.service.GetTattooByID(ctx, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"tattoo": tattoo})
}

func (h TattooHandler) UpdateTattoo(g *gin.Context) {
	ctx := g.Request.Context()
	var tattoo models.Tattoo
	if err := g.BindJSON(&tattoo); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateTattoo(ctx, tattoo.ID, tattoo); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"tattoo": tattoo})
}

func (h TattooHandler) DeleteTattoo(g *gin.Context) {
	ctx := g.Request.Context()
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}

	if err := h.service.DeleteTattoo(ctx, id); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}
