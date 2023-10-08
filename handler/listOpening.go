package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/goOpportunities/schemas"
)


// @BasePath /api/v1

// @Summary List Opening
// @Description List all job openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "openings listed", openings)
}