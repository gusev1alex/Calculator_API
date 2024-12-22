package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gusev1alex/Calculator_API.git/internal/calculator"
)

type Server struct {
	httpServer *gin.Engine
}

func NewServer() *Server {
	// Инициализируем Gin-сервер
	server := gin.Default()

	// Определяем маршруты
	server.POST("/api/v1/calculate", calculateHandler)

	return &Server{
		httpServer: server,
	}
}

func (s *Server) Run(addr string) error {
	// Запускаем сервер на указанном адресе
	return s.httpServer.Run(addr)
}

func calculateHandler(c *gin.Context) {
	var req struct {
		Expression string `json:"expression"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// Пример вычисления (в реальном случае вызовите метод из internal/calculator)
	result, err := calculator.Calc(req.Expression)
	if err != nil {
		if err.Error() == "invalid expression" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": strconv.FormatFloat(result, 'f', 3, 64)})
}
