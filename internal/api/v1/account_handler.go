package v1

import (
	"log/slog"
	"net/http"

	"github.com/CaioDGallo/granite-identity/internal/logger"
	"github.com/CaioDGallo/granite-identity/internal/service"
	"github.com/gin-gonic/gin"
)

var accountService *service.AccountService

func RegisterRoutes(router *gin.Engine) {
	accountService = service.NewAccountService()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/accounts", createAccount)
		v1.GET("/accounts/:id", getAccount)
	}
}

func createAccount(c *gin.Context) {
	logger.GetLogger().Info("handling create account request")

	var req service.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := accountService.CreateAccount(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.GetLogger().Info("account created", slog.String("account_id", account.ID.String()))
	c.JSON(http.StatusOK, account)
}

func getAccount(c *gin.Context) {
	logger.GetLogger().Info("handling get account request")

	id := c.Param("id")
	account, err := accountService.GetAccountByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.GetLogger().Info("account retrieved", slog.String("account_id", account.ID.String()))
	c.JSON(http.StatusOK, account)
}
