package main

import (
	drive "code.google.com/p/google-api-go-client/drive/v2"
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	registrarServico("baixar_arquivo_mais_recente", drive.DriveScope, baixar_arquivo_mais_recente)
}

func baixar_arquivo_mais_recente(client *http.Client, argv []string) {
	if len(argv) < 2 {
		fmt.Println(os.Stderr, "Uso: baixar_arquivo_mais_recente [arquivo.extensao] [Destino\\NomeArquivo.extensao] ...")
		return
	}

	service, _ := drive.New(client)

	query := fmt.Sprintf("title = '%s'", argv[0])

	lista, err := service.Files.List().Q(query).Do()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var itemMaisNovo = lista.Items[0]
	for i := range lista.Items {
		if lista.Items[i].ModifiedDate > itemMaisNovo.ModifiedDate {
			itemMaisNovo = lista.Items[i]
		}
	}

	BaixarArquivo(itemMaisNovo, argv[1])
}

func BaixarArquivo(item *drive.File, destino string) (string, error) {
	requisicaoArquivo, err := http.Get(item.DownloadUrl)
	if err != nil {
		fmt.Printf("Erro ao fazer a requisição: %v\n", err)
		return "", nil
	}
	respostaAquisicao, err := OAuthClient(Config).Transport.RoundTrip(requisicaoArquivo.Request)
	defer respostaAquisicao.Body.Close()
	if err != nil {
		fmt.Printf("Erro ao fazer o download: %v\n", err)
		return "", nil
	}

	arquivoLocal, err := os.Create(destino)
	defer arquivoLocal.Close()
	io.Copy(arquivoLocal, respostaAquisicao.Body)
	return "", nil
}
