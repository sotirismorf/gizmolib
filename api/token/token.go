package token

import (
	// "database/sql"
	"net/http"
	"time"
	"strconv"

	"github.com/sotirismorf/microservice/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Service struct {
	queries *database.Queries
}

type credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewService(queries *database.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.POST("/token", s.GenerateToken)
}

func GenerateToken(user_id uint) (string, error) {

	token_lifespan,err := strconv.Atoi("1")

	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("this_is_a_secret"))

}

func LoginCheck(username string, password string) (string,error) {
	
	var err error

	token, err := GenerateToken(1)

	if err != nil {
		return "",err
	}

	return token,nil
	
}

func (s *Service) GenerateToken(c *gin.Context) {
	// Parse request
	var request credentials
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// token, err := models.LoginCheck(u.Username, u.Password)
	token, err := LoginCheck("sotiris", "password")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	// Build response
	// c.IndentedJSON(http.StatusCreated, response)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
