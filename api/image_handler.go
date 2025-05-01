package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// ImageHandler handles image-related requests
type ImageHandler struct {
	uploadDir string
}

// NewImageHandler creates a new image handler
func NewImageHandler(uploadDir string) *ImageHandler {
	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create upload directory: %v", err))
	}

	return &ImageHandler{
		uploadDir: uploadDir,
	}
}

// RegisterRoutes registers the routes for the image handler
func (h *ImageHandler) RegisterRoutes(router *gin.RouterGroup) {
	images := router.Group("/images")
	{
		images.POST("/upload", h.UploadImage)
		images.DELETE("", h.DeleteImage)
	}
}

// @Summary Upload an image
// @Description Upload an image file and return its URL
// @Tags images
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file"
// @Param type formData string true "Image type (restaurant_logo, restaurant_cover, dish)"
// @Param id formData string true "ID of the entity (restaurant_id or dish_id)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/images/upload [post]
func (h *ImageHandler) UploadImage(c *gin.Context) {
	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Validate file type
	imageType := c.PostForm("type")
	if imageType != "restaurant_logo" && imageType != "restaurant_cover" && imageType != "dish" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image type"})
		return
	}

	// Get entity ID
	entityID := c.PostForm("id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entity ID is required"})
		return
	}

	// Validate file size (max 5MB)
	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size too large (max 5MB)"})
		return
	}

	// Validate file extension
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only jpg, jpeg, and png are allowed"})
		return
	}

	// Generate filename based on type and ID
	filename := fmt.Sprintf("%s_%s%s", imageType, entityID, ext)
	filepath := filepath.Join(h.uploadDir, filename)

	// Save the file
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Return the URL
	imageURL := fmt.Sprintf("/images/%s", filename)
	c.JSON(http.StatusOK, gin.H{
		"url": imageURL,
	})
}

// @Summary Delete an image
// @Description Delete an uploaded image
// @Tags images
// @Produce json
// @Param url query string true "Image URL"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/images [delete]
func (h *ImageHandler) DeleteImage(c *gin.Context) {
	imageURL := c.Query("url")
	if imageURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image URL is required"})
		return
	}

	// Extract filename from URL
	filename := filepath.Base(imageURL)
	filepath := filepath.Join(h.uploadDir, filename)

	// Delete the file
	if err := os.Remove(filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Image deleted successfully",
	})
}

// DeleteImageByURL deletes an image file by its URL
func (h *ImageHandler) DeleteImageByURL(imageURL string) error {
	if imageURL == "" {
		return fmt.Errorf("image URL is required")
	}

	// Extract filename from URL
	filename := filepath.Base(imageURL)
	filepath := filepath.Join(h.uploadDir, filename)

	// Delete the file
	if err := os.Remove(filepath); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}
