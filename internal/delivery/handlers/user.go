package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tp-project/internal/models/user"
)

type ServiceUser interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserHandler struct {
	service ServiceUser
}

func NewUserHandler(userService ServiceUser) UserHandler {
	return UserHandler{
		service: userService,
	}
}

func (h *UserHandler) GetAllUsers(g *gin.Context) {
	ctx := g.Request.Context()
	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) CreateUser(g *gin.Context) {
	ctx := g.Request.Context()
	var user models.User
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateUser(ctx, user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *UserHandler) GetUserByID(g *gin.Context) {
	ctx := g.Request.Context()
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}
	user, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) GetUserByEmail(g *gin.Context) {
	ctx := g.Request.Context()
	email := g.Param("email")
	user, err := h.service.GetUserByEmail(ctx, email)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) UpdateUser(g *gin.Context) {
	ctx := g.Request.Context()
	var user models.User
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateUser(ctx, user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) DeleteUser(g *gin.Context) {
	ctx := g.Request.Context()
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
		return
	}

	if err := h.service.DeleteUser(ctx, id); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}
