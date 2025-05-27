package controller

import (
	"net/http"
	"strconv"

	"github.com/alishercodecrafter/orderpackscalculator/internal/model"
	"github.com/gin-gonic/gin"
)

// PacksService defines the interface for pack calculation services
type PacksService interface {
	// GetPacks returns all available packs
	GetPacks() model.Packs
	// AddPack adds a new pack
	AddPack(pack model.Pack) error
	// RemovePack removes a pack by its size
	RemovePack(packSize model.PackSize) error
	// CalculatePacks calculates the optimal number of packs needed for an order
	CalculatePacks(orderSize int) (model.CalculationResponse, error)
}

// PacksController handles HTTP requests
type PacksController struct {
	// PacksService is the service for pack calculations
	service PacksService
}

// NewPacksController creates a new PacksController
func NewPacksController(service PacksService) *PacksController {
	return &PacksController{
		service: service,
	}
}

// GetPacks returns all packs
// @Summary Get all packs
// @Description Get a list of all available packs
// @Produce json
// @Success 200 {array} model.Pack "List of packs"
// @Router /api/packs [get]
func (c *PacksController) GetPacks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.GetPacks())
}

// AddPack adds a new pack
// @Summary Add pack
// @Description Add a new pack with its properties
// @Accept json
// @Produce json
// @Param request body model.AddPackRequest true "Pack to add"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 400 {object} map[string]string "Error response"
// @Router /api/packs [post]
func (c *PacksController) AddPack(ctx *gin.Context) {
	var req model.AddPackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})

		return
	}

	if req.Pack.Size <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Pack size must be greater than zero"})

		return
	}

	if err := c.service.AddPack(req.Pack); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// RemovePack removes a pack
// @Summary Remove pack
// @Description Remove a pack by its size value
// @Produce json
// @Param size path int true "Pack size to remove"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 400 {object} map[string]string "Error response"
// @Router /api/packs/{size} [delete]
func (c *PacksController) RemovePack(ctx *gin.Context) {
	sizeStr := ctx.Param("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pack size"})

		return
	}

	if err := c.service.RemovePack(model.PackSize(size)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// GetIndex renders the main page
// @Summary Render main page
// @Description Get the main page of the Pack Calculator app
// @Produce html
// @Success 200 {string} string "HTML page"
// @Router / [get]
func (c *PacksController) GetIndex(ctx *gin.Context) {
	packs := c.service.GetPacks()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"packs": packs,
	})
}

// CalculatePacks calculates the number of packs needed
// @Summary Calculate packs
// @Description Calculate the optimal number of packs needed for an order
// @Accept json
// @Produce json
// @Param request body model.CalculationRequest true "Order size"
// @Success 200 {object} model.CalculationResponse "Calculation result"
// @Failure 400 {object} map[string]string "Error response"
// @Router /api/calculate [post]
func (c *PacksController) CalculatePacks(ctx *gin.Context) {
	var req model.CalculationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})

		return
	}

	if req.OrderSize <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Order size must be greater than zero"})

		return
	}

	result, err := c.service.CalculatePacks(req.OrderSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, result)
}
