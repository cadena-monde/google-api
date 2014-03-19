package main

import (
	drive "code.google.com/p/google-api-go-client/drive/v2"
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	registrarServico("download", drive.DriveScope, download)
}

func download(client *http.Client, argv []string) {

	service, _ := drive.New(client)

	query := fmt.Sprintf("title = '%s'", argv[0])

	lista, err := service.Files.List().Q(query).Do()
	if err != nil {
		fmt.Println(err.Error())
	}

	var itemMaisNovo = lista.Items[0]
	for i := range lista.Items {
		if lista.Items[i].ModifiedDate > itemMaisNovo.ModifiedDate {
			itemMaisNovo = lista.Items[i]
		}
	}

	req, err := http.NewRequest("GET", itemMaisNovo.DownloadUrl, nil)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}
	resp, err := OAuthClient(Config).Transport.RoundTrip(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}

	f, err := os.Create(argv[1])
	defer f.Close()
	io.Copy(f, resp.Body)
}
