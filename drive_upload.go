package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	drive "code.google.com/p/google-api-go-client/drive/v2"
)

func init() {
	registrarServico("upload", drive.DriveScope, upload)
}

func upload(client *http.Client, argv []string) {
	if len(argv) < 2 {
		fmt.Println(os.Stderr, "Uso: upload [pasta_remota] [arquivo1] [arquivo2] ...")
		return
	}

	pastas := obterPastas(argv[0])
	if len(pastas) == 0 {
		log.Fatal("Pasta remota inválida: %v", argv[0])
	}

	arquivos := obterListaArquivos(argv[1:])

	service, _ := drive.New(client)

	novaPastaId, err := CriarPastasRemotas(service, pastas)
	if err != nil {
		log.Fatalf("Erro ao criar pastas remotas: %v", err)
	}
	pastaBuild := obterObjetoPasta(novaPastaId)

	for nome, caminho := range arquivos {
		fmt.Println("Enviando arquivo: ", nome)
		arquivo, err := os.Open(caminho)
		if err != nil {
			log.Fatalf("Erro ao abrir o arquivo %q: %v", caminho, err)
		}

		arquivoDrive := &drive.File{Title: nome, Parents: pastaBuild}
		driveFile, err := service.Files.Insert(arquivoDrive).Media(arquivo).Do()
		if err != nil {
			log.Fatalf("Erro ao enviar arquivo: %v", err)
			return
		}

		fmt.Println("Arquivo enviado com sucesso: ", driveFile.Title)
		fmt.Println("URL: ", driveFile.DownloadUrl)
	}
}

func CriarPastasRemotas(service *drive.Service, pastas []string) (string, error) {
	if len(pastas) == 0 {
		return "", errors.New("Nenhuma pasta informada.")
	}

	IdPasta := "root"
	// var error err
	for _, pasta := range pastas {
		id, err := ObterIdPastaRemota(service, IdPasta, pasta)
		if err != nil {
			return "", err
		}
		IdPasta = id
	}
	return IdPasta, nil
}

func ObterIdPastaRemota(service *drive.Service, parentId, pasta string) (string, error) {
	query := fmt.Sprintf("title = '%s' and mimeType = 'application/vnd.google-apps.folder' and '%s' in parents", pasta, parentId)
	lista, err := service.Files.List().Q(query).Do()
	if err != nil {
		return "", err
	}

	if len(lista.Items) == 0 {
		return criarPasta(service, parentId, pasta), nil
	} else {
		return lista.Items[0].Id, nil
	}

}

func obterObjetoPasta(pastaId string) []*drive.ParentReference {
	listaPastas := make([]*drive.ParentReference, 1, 1)
	listaPastas[0] = &drive.ParentReference{Id: pastaId}
	return listaPastas
}

func criarPasta(service *drive.Service, parentId, nome string) string {
	pastaParent := obterObjetoPasta(parentId)
	novaPasta := &drive.File{Title: nome, Parents: pastaParent, MimeType: "application/vnd.google-apps.folder"}

	pastaCriada, err := service.Files.Insert(novaPasta).Do()
	if err != nil {
		log.Fatalf("Erro ao criar pasta: %v", err)
		return ""
	}
	return pastaCriada.Id
}

func obterListaArquivos(argv []string) map[string]string {
	var arquivos = make(map[string]string)
	for i := 0; i < len(argv); i++ {
		arquivos[filepath.Base(argv[i])] = argv[i]
	}
	return arquivos
}

func obterPastas(parametro string) []string {
	if strings.Contains(parametro, "\\") {
		return strings.Split(parametro, "\\")
	} else {
		return strings.Split(parametro, "/")
	}
}
