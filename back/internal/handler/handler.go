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
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	var user model.User
	if err := h.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or Password incorrect"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or Password incorrect"})
		return
	}

	// トークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":    user.Username,
		"user_id": user.ID,
		"iss":     os.Getenv("JWT_ISS"),
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"nbf":     time.Now().Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data": gin.H{
			"token": tokenString,
			"user":  user.Username,
		},
	})
}

func (h *Handler) registerUser(c *gin.Context) {
	var input model.User

	// JSONデータをバインド
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	// ユーザー名の重複チェック
	var existing model.User
	if err := h.DB.Where("username = ?", input.Username).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password hash error"})
		return
	}

	// 登録用データ作成
	newUser := model.User{
		Username: input.Username,
		Password: string(hashedPassword), // stringにして保存もOK
	}

	// DBに保存
	if err := h.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"message":  "User registration successful.",
		"username": newUser.Username,
	})
}
