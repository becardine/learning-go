package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/becardine/learning-go/apis/internal/dto"
	"github.com/becardine/learning-go/apis/internal/entity"
	"github.com/becardine/learning-go/apis/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProduct(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// normalmente essa parte seria feita em um service
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
