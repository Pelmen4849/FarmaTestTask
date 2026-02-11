package handler

import (
	"net/http"
	"strconv"

	"medical_farm/back/internal/service"

	"github.com/gin-gonic/gin"
)

type DrugHandler struct {
	drugService service.DrugService
}

func NewDrugHandler(drugService service.DrugService) *DrugHandler {
	return &DrugHandler{drugService: drugService}
}

// GET /api/drugs?shop_id=1
func (h *DrugHandler) GetAvailableDrugs(c *gin.Context) {
	shopIDStr := c.Query("shop_id")
	if shopIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shop_id is required"})
		return
	}
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop_id"})
		return
	}

	drugs, err := h.drugService.ListAvailableDrugs(c.Request.Context(), shopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, drugs)
}

// GET /api/drugs/:id
func (h *DrugHandler) GetDrugByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	drug, err := h.drugService.GetDrugDetail(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if drug == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "drug not found"})
		return
	}
	c.JSON(http.StatusOK, drug)
}
