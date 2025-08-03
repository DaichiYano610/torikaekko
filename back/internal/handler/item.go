package handler

import (
	"fmt"
	"myapp/internal/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadGiveItem(c *gin.Context) {
	// 認証済みユーザーID取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	name := c.PostForm("itemName")
	want := c.PostForm("itemWant")

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid multipart form"})
		return
	}

	files := form.File["itemImages"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no image uploaded"})
		return
	}

	var imagePaths []string
	for _, file := range files {
		if !strings.HasSuffix(strings.ToLower(file.Filename), ".jpg") && !strings.HasSuffix(strings.ToLower(file.Filename), ".jpeg") {
			continue
		}

		filename := fmt.Sprintf("item_images/%d_%s", time.Now().UnixNano(), file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			continue
		}
		imagePaths = append(imagePaths, filename)
	}

	if len(imagePaths) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no valid images saved"})
		return
	}

	item := model.Item{
		Name:       name,
		Want:       want,
		ImagePaths: strings.Join(imagePaths, ","), // カンマで結合
		UserID:     userID.(uint),
	}

	if err := h.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "item uploaded successfully",
		"item_id": item.ID,
	})
}
