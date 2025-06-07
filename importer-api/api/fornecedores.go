package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getFornecedoresByIdRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

// getFornecedoresById retrieves a fornecedor by ID
func (server *Server) getFornecedoresById(ctx *gin.Context) {
	var request getFornecedoresByIdRequest
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fornecedor, err := server.store.GetFornecedoresById(ctx, request.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, fornecedor)
}
