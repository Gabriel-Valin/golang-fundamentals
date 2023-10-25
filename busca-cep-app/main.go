package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ddd         string `json:"ddd"`
}

var enderecos []Endereco

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("enderecos.html").ParseFiles("enderecos.html"))
		err := t.Execute(w, enderecos)

		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("/cep", BuscaCEPHandler)
	http.ListenAndServe(":8081", nil)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cep" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		cep := r.FormValue("cep")

		if cep == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		endereco, err := BuscaCep(cep)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(endereco)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}
}

func BuscaCep(cep string) (*Endereco, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var endereco Endereco
	err = json.Unmarshal(body, &endereco)
	enderecos = append(enderecos, endereco)
	if err != nil {
		return nil, err
	}
	return &endereco, nil
}
