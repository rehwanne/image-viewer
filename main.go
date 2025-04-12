package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	currentImage     []byte
	currentImageMime string
	imageMutex      sync.RWMutex
)

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Max 10MB image
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the image", http.StatusInternalServerError)
		return
	}

	imageMutex.Lock()
	currentImage = imageData
	currentImageMime = handler.Header.Get("Content-Type")
	imageMutex.Unlock()

	fmt.Fprintf(w, "Image uploaded successfully!")
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	imageMutex.RLock()
	defer imageMutex.RUnlock()

	if currentImage == nil {
		http.Error(w, "No image available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", currentImageMime)
	w.Write(currentImage)
}

func main() {
	http.HandleFunc("/upload-image", uploadImageHandler)
	http.HandleFunc("/", imageHandler)

	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(port, nil)
}
