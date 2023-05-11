package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/devnandito/imageapi/models"
	"github.com/gorilla/mux"
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

	im := models.Image{}
	textBytes := []byte(jsonData)
	er := im.ToText(textBytes, &im)
	if er != nil {
		panic(er)
	}

	data := models.Image{
		Title: im.Title,
		Description: im.Description,
		Url: im.Url,
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
	r.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	fmt.Println("File Type: ", handler.Header.Get("Content-Type"))

	f := fmt.Sprintf("img-%d*%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))

	// Upload file
	tempFile, err := ioutil.TempFile("./assets/uploads", f)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		fmt.Println(err3)
	}
	tempFile.Write(fileBytes)
	fmt.Println("Done")
	
	// save to database
	title := r.PostFormValue("description")
	description := r.PostFormValue("description")
	slice := strings.Split(tempFile.Name(), ".")
	uri := slice[1]+"."+slice[2]
	data := models.Image{
		Title: title,
		Description: description,
		Url: uri,
	}

	response, err := img.CreateImage(&data)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Println("Data inserted", response)
}

// HandleApiGetOneImage get a one image
func HandleApiGetOneImage(w http.ResponseWriter, r *http.Request) {
	imageId := mux.Vars(r)["id"]
	pk, err :=  strconv.Atoi(imageId)
	if err != nil {
		panic(err)
	}

	objects, err := img.GetOneImage(pk)
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



// func HandleUploadImage(w http.ResponseWriter, r *http.Request) {
// 	// 32 MB is the default used by FormFile() function
// 	if err := r.ParseMultipartForm(10 * 1024 * 1024); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 	}

// 	// Get a reference to the fileHeaders.
// 	// They are accessible only after ParseMultipartForm is called
// 	files := r.MultipartForm.File["file"]

// 	var errNew string
// 	var http_status int

// 	for _, fileHeader := range files {
// 			// Open the file
// 			file, err := fileHeader.Open()
// 			if err != nil {
// 					errNew = err.Error()
// 					http_status = http.StatusInternalServerError
// 					break
// 			}

// 			defer file.Close()

// 			buff := make([]byte, 512)
// 			_, err = file.Read(buff)
// 			if err != nil {
// 					errNew = err.Error()
// 					http_status = http.StatusInternalServerError
// 					break
// 			}

// 			// checking the content type
// 			// so we don't allow files other than images
// 			filetype := http.DetectContentType(buff)
// 			if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
// 					errNew = "The provided file format is not allowed. Please upload a JPEG,JPG or PNG image"
// 					http_status = http.StatusBadRequest
// 					break
// 			}

// 			_, err = file.Seek(0, io.SeekStart)
// 			if err != nil {
// 					errNew = err.Error()
// 					http_status = http.StatusInternalServerError
// 					break
// 			}

// 			err = os.MkdirAll("./uploads", os.ModePerm)
// 			if err != nil {
// 					errNew = err.Error()
// 					http_status = http.StatusInternalServerError
// 					break
// 			}

// 			f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
// 			if err != nil {
// 					errNew = err.Error()
// 					http_status = http.StatusBadRequest
// 					break
// 			}

// 			defer f.Close()

// 			_, err = io.Copy(f, file)
// 			if err != nil {
// 					errNew = err.Error()
// 					http_status = http.StatusBadRequest
// 					break
// 			}
// 	}
// 	message := "file uploaded successfully"
// 	messageType := "S"

// 	if errNew != "" {
// 			message = errNew
// 			messageType = "E"
// 	}

// 	if http_status == 0 {
// 			http_status = http.StatusOK
// 	}

// 	resp := map[string]interface{}{
// 			"messageType": messageType,
// 			"message":     message,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http_status)
// 	json.NewEncoder(w).Encode(resp)
// }