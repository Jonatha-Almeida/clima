package models

import (
	"github.com/Jonatha-Almeida/clima/db"
)

type Clima struct {
	Id            int
	Cidade        string
	Municipio     string
	ChuvaEsperada int
	ChuvaQueCaiu  int
}

func ListagemDosClimas() []Clima {
	db := db.ConectaAoBancoDeDados()

	listandoOsClimas, err := db.Query("select * from climas")
	if err != nil {
		panic(err.Error())
	}

	c := Clima{}
	climas := []Clima{}

	for listandoOsClimas.Next() {
		var id, chuvaEsperada, chuvaQueCaiu int
		var cidade, municipio string

		err = listandoOsClimas.Scan(&id, &cidade, &municipio, &chuvaEsperada, &chuvaQueCaiu)
		if err != nil {
			panic(err.Error())
		}

		c.Id = id
		c.Cidade = cidade
		c.Municipio = municipio
		c.ChuvaEsperada = chuvaEsperada
		c.ChuvaQueCaiu = chuvaQueCaiu

		climas = append(climas, c)
	}
	defer db.Close()
	return climas
}
func CriaNovoClima(cidade, municipio string, chuvaEsperada, chuvaQueCaiu int) {
	db := db.ConectaAoBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into climas(cidade, municipio, chuvaEsperada, chuvaQueCaiu) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(cidade, municipio, chuvaEsperada, chuvaQueCaiu)
	defer db.Close()

}

func DeletaClima(id string) {
	db := db.ConectaAoBancoDeDados()

	deletarClima, err := db.Prepare("delete from climas where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarClima.Exec(id)
	defer db.Close()
}

func EditaClima(id string) Clima {
	db := db.ConectaAoBancoDeDados()

	editandoClimaNoBanco, err := db.Query("select * from climas where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	climaParaAtualizar := Clima{}

	for editandoClimaNoBanco.Next() {
		var id, chuvaEsperada, chuvaQueCaiu int
		var cidade, municipio string

		err = editandoClimaNoBanco.Scan(&id, &cidade, &municipio, &chuvaEsperada, &chuvaQueCaiu)
		if err != nil {
			panic(err.Error())
		}
		climaParaAtualizar.Id = id
		climaParaAtualizar.Cidade = cidade
		climaParaAtualizar.Municipio = municipio
		climaParaAtualizar.ChuvaEsperada = chuvaEsperada
		climaParaAtualizar.ChuvaQueCaiu = chuvaQueCaiu
	}
	defer db.Close()
	return climaParaAtualizar
}

func AtualizaClima(id int, cidade, municipio string, chuvaEsperada, chuvaQueCaiu int) {
	db := db.ConectaAoBancoDeDados()

	AtualizaClima, err := db.Prepare("update climas set cidade=$1, municipio=$2, chuvaEsperada=$3, chuvaQueCaiu=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaClima.Exec(cidade, municipio, chuvaEsperada, chuvaQueCaiu, id)
	defer db.Close()
}
