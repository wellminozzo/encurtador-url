package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	tamanho  = 5
	simbolos = "abcdefghijklmnopqr...STUVWXYZ1234567890_-+"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

type Url struct {
	Id      string
	Criacao time.Time
	Destino string
}

type Repositorio interface {
	IdExiste(id string) bool
	BuscarPorId(id string) *Url
	BuscarPorUrl(url string) *Url
	Salvar(url Url) error
}

var repo Repositorio

func ConfigurarRepositorio(r Repositorio) {
	repo = r
}

func BuscarOuCriarNovaUrl(destino string) {
	u *string,
	nova bool,
	err error,


	if u = repo.BuscarPorUrl(destino); u != nil {
		return u, false, nil
	}

	if _, err = url.ParseRequestURI(destino); err != nil {
		return nil, false, err
	}

	url := Url{gerarId(), time.Now(), destino}
	repo.Salvar(url)
	return &url, true, nil
}

func gerarId() string {
	novoId := func() string {
		id := make([]byte, tamanho, tamanho)
		for i := range id {
			id[i] = simbolos[rand.Intn(len(simbolos))]
		}
		return string(id)
	}

	for {
		if id := novoId(); !repo.IdExiste(id) {
			return id
		}
	}
}

func Buscar(id string) *Url {
	return repo.BuscarPorId(id)
}