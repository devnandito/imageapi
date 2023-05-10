package models

import (
	"encoding/json"

	"github.com/devnandito/imageapi/lib"
	"gorm.io/gorm"
)

// Module access public
type Image struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
}

// ToJson return to r.body to json
func (i *Image) ToJson(img Image) ([]byte, error) {
	return json.Marshal(img)
}

// ToText return r.body to text
func (m *Image) ToText(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

//CreateImage created a new image
func (i Image) CreateImage(img *Image) (Image, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Image{})
	response := db.Create(&img)
	data := Image{
		Title: img.Title,
		Description: img.Description,
		Url: img.Url,
	}

	return data, response.Error
}

// GetImage get one image
func (i Image) GetOneImage(id int) (Image, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Find(&i, id)
	return i, response.Error
}

// UpdateImage update an image
func (i Image) UpdateImage(id int, img *Image) (Image, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Model(&i).Where("id = ?", id).Updates(Image{Title: img.Title, Description: img.Description, Url: img.Url})
	return i, response.Error
}

// DeleteImage delete an image
func (i Image) DeleteImage(id int) error {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	response := db.Delete(&i, id)
	return response.Error
}

// ShowImage list images
func (i Image) ShowImage() ([]Image, error) {
	conn := lib.NewConfig()
	db := conn.DsnStringGorm()
	db.AutoMigrate(&Image{})
  var objects []Image
	response := db.Find(&objects)
	return objects, response.Error
}