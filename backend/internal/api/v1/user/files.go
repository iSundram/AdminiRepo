
package user

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type FileController struct{}

func NewFileController() *FileController {
	return &FileController{}
}

// GetFiles returns files in a directory
func (fc *FileController) GetFiles(c *gin.Context) {
	path := c.DefaultQuery("path", "/")

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Files retrieved successfully",
		"data": gin.H{
			"current_path": path,
			"files": []gin.H{
				{
					"name":        "index.html",
					"type":        "file",
					"size":        "2.5KB",
					"modified":    "2024-01-15T14:30:00Z",
					"permissions": "644",
					"mime_type":   "text/html",
				},
				{
					"name":        "assets",
					"type":        "directory",
					"size":        "-",
					"modified":    "2024-01-14T10:20:00Z",
					"permissions": "755",
				},
				{
					"name":        "config.php",
					"type":        "file",
					"size":        "1.2KB",
					"modified":    "2024-01-16T09:15:00Z",
					"permissions": "644",
					"mime_type":   "application/x-php",
				},
			},
		},
	})
}

// UploadFile handles file upload
func (fc *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	path := c.DefaultPostForm("path", "/")
	filename := filepath.Join(path, file.Filename)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "File uploaded successfully",
		"data": gin.H{
			"filename": file.Filename,
			"path":     filename,
			"size":     file.Size,
		},
	})
}

// CreateFolder creates a new folder
func (fc *FileController) CreateFolder(c *gin.Context) {
	var req struct {
		Path       string `json:"path" binding:"required"`
		FolderName string `json:"folder_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	folderPath := filepath.Join(req.Path, req.FolderName)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Folder created successfully",
		"data": gin.H{
			"folder_name": req.FolderName,
			"path":        folderPath,
		},
	})
}

// DeleteFile deletes a file or folder
func (fc *FileController) DeleteFile(c *gin.Context) {
	var req struct {
		Path string `json:"path" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "File deleted successfully",
		"data":    gin.H{"path": req.Path},
	})
}

// RenameFile renames a file or folder
func (fc *FileController) RenameFile(c *gin.Context) {
	var req struct {
		OldPath string `json:"old_path" binding:"required"`
		NewName string `json:"new_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "File renamed successfully",
		"data": gin.H{
			"old_path": req.OldPath,
			"new_name": req.NewName,
		},
	})
}

// GetFileContent returns file content for editing
func (fc *FileController) GetFileContent(c *gin.Context) {
	filePath := c.Query("path")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File path is required"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "File content retrieved successfully",
		"data": gin.H{
			"path":    filePath,
			"content": "<!DOCTYPE html>\n<html>\n<head>\n    <title>Welcome</title>\n</head>\n<body>\n    <h1>Hello World!</h1>\n</body>\n</html>",
			"encoding": "utf-8",
		},
	})
}

// SaveFileContent saves file content
func (fc *FileController) SaveFileContent(c *gin.Context) {
	var req struct {
		Path    string `json:"path" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "File saved successfully",
		"data": gin.H{
			"path": req.Path,
			"size": len(req.Content),
		},
	})
}
