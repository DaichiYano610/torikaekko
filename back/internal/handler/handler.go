package handler

import (
	"net/http"
	"time"

	"myapp/internal/middleware"
	"myapp/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/login", login)
	r.POST("/register_user", register_user)

	//認証必須のグループ
	authorized_r := r.Group("/")
	authorized_r.Use(middleware.AuthMiddleware())
	{
		authorized_r.GET("/hello", hello_world)
	}
}

func hello_world(c *gin.Context) {
	tmp := []model.User{
		{Username: "admin", Password: "admin"},
		{Username: "admin", Password: "admin"},
	}

	c.IndentedJSON(http.StatusOK, tmp)
}

func login(c *gin.Context) {
	login_data := model.User{}

	if err := c.ShouldBindJSON(&login_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invailed JSON format"})
		return
	}

	// 認証（今はハードコード、将来はDB）
	if login_data.Username != "admin" || login_data.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Password error"})
		return
	}

	// 有効期限付きJWTトークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": login_data.Username,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(24 * time.Hour).Unix(), // トークンは24時間で期限切れ
		"nbf":  time.Now().Unix(),
	})

	secret := "admin" // 実際は .env 等に保管する
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token creation failed"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}

func register_user(c *gin.Context) {
	register_data := model.User{}

	if err := c.ShouldBindJSON(&register_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invailed JSON format"})
		return
	}

	//DBに同じユーザ情報があるか確認
	//あった場合、登録なし
	//ない場合、登録

	c.JSON(http.StatusOK, gin.H{"register": register_data})
}
