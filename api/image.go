package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/devnandito/imageapi/models"
)

var img models.Image

// HandlerApiShowImage list images
func HandleApiShowImage(w http.ResponseWriter, r *http.Request) {
	objects, err := img.ShowImage()
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(&objects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// HandleApiCreateClient create a new image
func HandleApiCreateImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&img)
	
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	jsonData, err := img.ToJson(img)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cli := models.Image{}
	textBytes := []byte(jsonData)
	er := img.ToText(textBytes, &cli)

	if er != nil {
		panic(er)
	}

	data := models.Image{
		Title: img.Title,
		Description: img.Description,
		Url: cli.Url,
	}

	response, err := img.CreateImage(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	w.Header().Set("Content-type", "application/json")
	w.Write(jsonData)
}

// HandleUploadImage upload an image
func HandleUploadImage(w http.ResponseWriter, r *http.Request) {
	// 32 MB is the default used by FormFile() function
	if err := r.ParseMultipartForm(10 * 1024 * 1024); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := r.MultipartForm.File["file"]

	var errNew string
	var http_status int

	for _, fileHeader := range files {
			// Open the file
			file, err := fileHeader.Open()
			if err != nil {
					errNew = err.Error()
					http_status = http.StatusInternalServerError
					break
			}

			defer file.Close()

			buff := make([]byte, 512)
			_, err = file.Read(buff)
			if err != nil {
					errNew = err.Error()
					http_status = http.StatusInternalServerError
					break
			}

			// checking the content type
			// so we don't allow files other than images
			filetype := http.DetectContentType(buff)
			if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
					errNew = "The provided file format is not allowed. Please upload a JPEG,JPG or PNG image"
					http_status = http.StatusBadRequest
					break
			}

			_, err = file.Seek(0, io.SeekStart)
			if err != nil {
					errNew = err.Error()
					http_status = http.StatusInternalServerError
					break
			}

			err = os.MkdirAll("./uploads", os.ModePerm)
			if err != nil {
					errNew = err.Error()
					http_status = http.StatusInternalServerError
					break
			}

			f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
			if err != nil {
					errNew = err.Error()
					http_status = http.StatusBadRequest
					break
			}

			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
					errNew = err.Error()
					http_status = http.StatusBadRequest
					break
			}
	}
	message := "file uploaded successfully"
	messageType := "S"

	if errNew != "" {
			message = errNew
			messageType = "E"
	}

	if http_status == 0 {
			http_status = http.StatusOK
	}

	resp := map[string]interface{}{
			"messageType": messageType,
			"message":     message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http_status)
	json.NewEncoder(w).Encode(resp)
}

// HandleUpload
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := r.FormFile("myfile")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Println("File Info")
	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	fmt.Println("File Type: ", handler.Header)
}