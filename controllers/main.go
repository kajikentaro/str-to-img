package controllers

import (
	"image/png"
	"net/http"

	"github.com/kajikentaro/playground/str-to-img/services"
)

type GenImageController struct {
	service *services.GenImageService
}

func NewGenImageController(s *services.GenImageService) *GenImageController {
	return &GenImageController{service: s}
}

func (c *GenImageController) GenImageHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	text := query.Get("text")
	if text == "" {
		text = "クエリパラメーター 'text' にテキストを指定してください.  例: ?text=hoge"
	}
	image := c.service.GenImage(text)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	png.Encode(w, image)
}
