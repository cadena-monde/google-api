package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	drive "code.google.com/p/google-api-go-client/drive/v2"
)

const (
	// TODO: Pasta parent ser configurável
	IdPastaMondeDriveDoBuilder = "0B3LEEXgVtpacc3VCejloUnRhTGs"
)

func init() {
	registrarServico("upload", drive.DriveScope, upload)
}

func upload(client *http.Client, argv []string) {
	if len(argv) < 2 {
		fmt.Println(os.Stderr, "Uso: upload [pasta_remota] [arquivo1] [arquivo2] ...")
		return
	}

	subPasta := argv[0]
	arquivos := obterListaArquivos(argv[1:])

	service, _ := drive.New(client)

	novaPastaId := criarSubpasta(service, IdPastaMondeDriveDoBuilder, subPasta)
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

func obterObjetoPasta(pastaId string) []*drive.ParentReference {
	listaPastas := make([]*drive.ParentReference, 1, 1)
	listaPastas[0] = &drive.ParentReference{Id: pastaId}
	return listaPastas
}

func criarSubpasta(service *drive.Service, parentId, nome string) string {
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
