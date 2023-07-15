package controllers

import (
	"net/http"

	"github.com/muhardiansyah15/go-rest-api/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "Kemeja",
			"stok":         1000,
		},
		{
			"id":           2,
			"nama_product": "Celana",
			"stok":         10000,
		},
		{
			"id":           1,
			"nama_product": "Sepatu",
			"stok":         500,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)

}
