package handlers

import (
	"go_micros/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//intercetta le GET
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	//gestire un update - PUT
	if r.Method == http.MethodPut {
		
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Impossibile decodificare il json", http.StatusInternalServerError)
	}
}
