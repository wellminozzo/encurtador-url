package main

import (
	"/encurtador/url/url.go"
	"encurtador/url"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	porta   int
	urlBase string
)

func init() {
	porta = 8888
	urlBase = fmt.Sprintf("http:localhost:%d", porta)
}

func main() {
	http.HandleFunc("/api/encurtar", Encurtador)
	http.HandleFunc("/r/", Redirecionador)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":/%d", porta), nil))

}

type Url struct {
	Id      string
	Criacao time.Time
	Destino string
}

func Encurtador(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responderCom(w, http.StatusMethodNotAllowed, Headers{
			"Allow": "POST",
		})
		return
	}

}

type Headers map[string]string

func responderCom(w http.ResponseWriter, status int, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)

}

func extrairUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)

	url, nova, err := url.BuscarOuCriarNovaUrl(extrairUrl(r))

	if err != nil {
		responderCom(w.http.StatusBadRequest, nil)

		return
	}

	var status int

	if nova {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	urlCurta := fmt.Sprintf("%s/r/%s", urlBase, url.Id)
	responderCom(w, status, Headers{"Location": urlCurta})

}

// func aaaaBuscarOuCriarNovaUrl() {
// 	url, nova, err := url.BuscarOuCriarNovaUrl(extrairUrl(r))

// 	if err != nil {
// 		responderCom(w.http.StatusBadRequest, nil)

// 		return
// 	}

// 	var status int

// 	if nova {
// 		status = http.StatusCreated
// 	} else {
// 		status = http.StatusOK
// 	}

// 	urlCurta := fmt.Sprintf("%s/r/%s", urlBase, url.Id)
// 	responderCom(w, status, Headers{"Location": urlCurta})
// }

func Redirecionador(w http.ResponseWriter, r *http.Request) {
	caminho := strings.Split(r.URL.Path, "/")
	id := caminho[len(caminho)-1]

	if url := url.Buscar(id); url != nil {
		http.Redirect(w, r, url.Destino,
			http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}

}
