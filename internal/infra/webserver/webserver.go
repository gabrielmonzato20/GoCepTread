package webserver

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gabrielmonzato20/GoCepTread/internal/dto"
	"github.com/gabrielmonzato20/GoCepTread/internal/entity"
)

type Handler struct {
	urlServer1 string
	urlServer2 string
}

func NewHandler(url1 string, url2 string) *Handler {
	return &Handler{
		urlServer1: url1,
		urlServer2: url2,
	}
}
func (h *Handler) CallFistServer(cpf string) *entity.ResponseEntity {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	cpfBuilder := cpf[0:5] + "-" + cpf[5:8]
	req, err := http.NewRequestWithContext(ctx, "GET", strings.Replace(h.urlServer1, "{}", cpfBuilder, -1), nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data dto.CepServer1
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	response := entity.NewResponseEntity(strings.Replace(h.urlServer1, "{}", cpfBuilder, -1), result)
	return response
}

func (h *Handler) CallSecondServer(cpf string) *entity.ResponseEntity {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", strings.Replace(h.urlServer2, "{}", cpf, -1), nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data dto.CepServer2
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	response := entity.NewResponseEntity(strings.Replace(h.urlServer2, "{}", cpf, -1), result)

	return response
}
