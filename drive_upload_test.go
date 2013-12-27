package main

import (
	"fmt"
	"path"
	"runtime"
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
	if runtime.GOOS == "windows" {
		args[0] = fmt.Sprintf("c:\\temp\\%v", Arquivo1)
		args[1] = fmt.Sprintf("C:\\temp\\Nova Pasta\\%v", Arquivo2)
	} else {
		args[0] = fmt.Sprintf("/temp/%v", Arquivo1)
		args[1] = fmt.Sprintf("/temp/Nova Pasta/%v", Arquivo2)
	}
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

	var lista = []struct {
		Arquivo string
		Caminho string
	}{
		{Arquivo1, args[0]},
		{Arquivo2, args[1]},
	}

	var arquivos = obterListaArquivos(args)

	for _, item := range lista {
		caminho, ok := arquivos[item.Arquivo]
		if !ok {
			t.Errorf("Arquivo não retornado: %v arquivos: %v", item.Arquivo, arquivos)
		} else if caminho != item.Caminho {
			t.Errorf("Caminho inválido. Esperado %v retornado %v", item.Caminho, caminho)
		}
	}
}

func TestObterPastas_2PastasSeparadasPorBarraInvertida_PastasRetornadas(t *testing.T) {
	retorno := obterPastas(path.Join("Monde\\13.1.2.3"))
	verificarRetornoObterPastas(t, retorno)
}

func TestObterPastas_2PastasSeparadasPorBarra_PastasRetornadas(t *testing.T) {
	retorno := obterPastas("Monde/13.1.2.3")
	verificarRetornoObterPastas(t, retorno)
}

func verificarRetornoObterPastas(t *testing.T, retorno []string) {
	if len(retorno) != 2 {
		t.Error("Deveria retornar 2 pastas")
	} else if retorno[0] != "Monde" {
		t.Errorf("Esperado Monde retornado %v")
	} else if retorno[1] != "13.1.2.3" {
		t.Errorf("Esperado 13.1.2.3 retornado %v")
	}
}
