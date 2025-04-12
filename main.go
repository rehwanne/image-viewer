package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	uploadFolder = "./images"
	port         = ":8080"
)

func main() {
	// Upload-Ordner erstellen falls nicht vorhanden
	if err := os.MkdirAll(uploadFolder, os.ModePerm); err != nil {
		fmt.Printf("Fehler beim Erstellen des Upload-Ordners: %v\n", err)
		return
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Nur das aktuelle Bild servieren
	router.GET("/current-image", func(c *gin.Context) {
		imgPath := filepath.Join(uploadFolder, "current.jpg")
		c.File(imgPath)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/upload-image", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Kein Bild erhalten"})
			return
		}

		// Nur JPG/JPEG erlauben
		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".jpeg" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nur JPG/JPEG-Dateien erlaubt"})
			return
		}

		// Altes Bild löschen
		os.Remove(filepath.Join(uploadFolder, "current.jpg"))

		// Neues Bild speichern
		filePath := filepath.Join(uploadFolder, "current.jpg")
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Speichern der Datei"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Bild erfolgreich ersetzt",
			"path":    "/current-image",
		})
	})

	fmt.Printf("Server läuft auf http://localhost%s\n", port)
	router.Run(port)
}
