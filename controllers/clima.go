package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Jonatha-Almeida/clima/models"
)

var tempClima = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsClima := models.ListagemDosClimas()
	tempClima.ExecuteTemplate(w, "Index", todosOsClima)
}

func New(w http.ResponseWriter, r *http.Request) {
	tempClima.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cidade := r.FormValue("cidade")
		municipio := r.FormValue("municipio")
		chuvaEsperada := r.FormValue("chuvam")
		chuvaQueCaiu := r.FormValue("chuvac")

		chuvaEsperadaConvertendoparaInt, err := strconv.Atoi(chuvaEsperada)
		if err != nil {
			log.Println("Erro na conversão do Chuva Esperada:", err)
		}

		chuvaQueCaiuConvertendoparaInt, err := strconv.Atoi(chuvaQueCaiu)
		if err != nil {
			log.Println("Erro na conversão do Chuva Que Caiu:", err)
		}

		models.CriaNovoClima(cidade, municipio, chuvaEsperadaConvertendoparaInt, chuvaQueCaiuConvertendoparaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoClima := r.URL.Query().Get("id")
	models.DeletaClima(idDoClima)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoClima := r.URL.Query().Get("id")
	clima := models.EditaClima(idDoClima)
	tempClima.ExecuteTemplate(w, "Edit", clima)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		cidade := r.FormValue("cidade")
		municipio := r.FormValue("municipio")
		chuvaEsperada := r.FormValue("chuvam")
		chuvaQueCaiu := r.FormValue("chuvac")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		chuvaEsperadaConvertendoParaInt, err := strconv.Atoi(chuvaEsperada)
		if err != nil {
			log.Println("Erro na convesão do chuva espearda para int:", err)
		}

		chuvaQueCaiuConvertendoParaInt, err := strconv.Atoi(chuvaQueCaiu)
		if err != nil {
			log.Println("Erro na convesão da chuva que caiu para int:", err)
		}

		models.AtualizaClima(idConvertidaParaInt, cidade, municipio, chuvaEsperadaConvertendoParaInt, chuvaQueCaiuConvertendoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
