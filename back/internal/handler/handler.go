package handler

import (
	"net/http"
	"os"
	"time"

	"myapp/internal/middleware"
	"myapp/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	h := &Handler{DB: db}

	r.POST("/login", h.login)
	r.POST("/register_user", h.registerUser)

	//認証必須のグループ
	authorized_r := r.Group("/auth")
	authorized_r.Use(middleware.AuthMiddleware())
	{
		authorized_r.GET("/hello", hello_world)
	}
}

func hello_world(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func (h *Handler) login(c *gin.Context) {
	login_data := model.User{}

	if err := c.ShouldBindJSON(&login_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	// DBからユーザー取得
	var db_user model.Users
	if err := h.DB.Where("username = ?", login_data.Username).First(&db_user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or Password incorrect"})
		return
	}

	// bcryptでパスワード照合
	if err := bcrypt.CompareHashAndPassword(db_user.Password, []byte(login_data.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or Password incorrect"})
		return
	}

	// トークン作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": db_user.Username,
		"iss":  os.Getenv("JWT_ISS"),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
		"nbf":  time.Now().Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "token creation failed"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"message": "login successful",
	})
}

func (h *Handler) registerUser(c *gin.Context) {
	received_data := model.User{}

	if err := c.ShouldBindJSON(&received_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	// 同一ユーザー名が存在するか確認
	var existing model.User
	if err := h.DB.Where("username = ?", received_data.Username).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	// パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(received_data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "process error"})
		return
	}

	register_data := model.Users{
		Username: received_data.Username,
		Password: hash,
	}

	// 新規ユーザーを保存
	if err := h.DB.Create(&register_data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":  "User registration successful.",
		"username": gin.H{"username": received_data.Username},
	})
}
