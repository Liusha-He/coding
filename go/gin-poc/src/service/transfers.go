package service

import (
	"errors"
	"fmt"
	"net/http"
	"simple-bank/src/auth"
	"simple-bank/src/dao"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64   `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64   `json:"to_account_id" binding:"required,min=1"`
	Currency      string  `json:"currency" binding:"required, currency"`
	Amount        float64 `json:"amount" binding:"required,min=1.0"`
}

func (server *Server) validateAccount(ctx *gin.Context, accountID int64, currency string) (dao.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err = fmt.Errorf("Account [%d] currency mismatch: expected %s, but got %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	return account, true
}

// Create Transfer 	godoc
// @Summary 		create transfer
// @Schemes 		http
// @Description 	Takes an transfer json and store in DB, Returned saved json.
// @Tags 			services
// @Produce 		json
// @Param 			transfer  body	transferRequest true  "transfer json"
// @Success 		200 {object} dao.Transfer
// @Router 			/api/v1/services/transfers [post]
func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, valid := server.validateAccount(ctx, req.FromAccountID, req.Currency)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*auth.Payload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account does not belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, valid = server.validateAccount(ctx, req.ToAccountID, req.Currency)
	if !valid {
		return
	}

	arg := dao.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
