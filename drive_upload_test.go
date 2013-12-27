package main

import (
	"fmt"
	"testing"
)

const (
	Arquivo1 = "Arquivo1.exe"
	Arquivo2 = "Arquivo2.exe"
)

var (
	args = make([]string, 2, 2)
)

func setup() {
	args[0] = fmt.Sprintf("c:\\temp\\%v", Arquivo1)
	args[1] = fmt.Sprintf("C:\\temp\\Nova Pasta\\%v", Arquivo2)
}

func TestObterListaArquivos_2Arquivos_2ItensRetornado(t *testing.T) {
	setup()

	var arquivos = obterListaArquivos(args)

	if len(arquivos) != 2 {
		t.Errorf("Quantidade arquivos inválida expected 2 was %v.", len(arquivos))
	}

}

func TestObterListaArquivos_2Arquivos_NomeECaminhoRetornado(t *testing.T) {
	setup()

	var arquivos = obterListaArquivos(args)

	caminho, ok := arquivos[Arquivo1]
	if !ok {
		t.Errorf("Arquivo %v não retornado na lista", Arquivo1)
	} else if caminho != args[0] {
		t.Errorf("Caminho inválido. Esperado %v retornado %v", args[0], caminho)
	}

	caminho, ok = arquivos[Arquivo2]
	if !ok {
		t.Errorf("Arquivo %v não retornado na lista", Arquivo2)
	} else if caminho != args[1] {
		t.Errorf("Caminho inválido. Esperado %v retornado %v", args[1], caminho)
	}
}
