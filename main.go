// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Uso: Google_API.exe [serviço] [argumentos do serviço]\n\nServiços disponíveis:\n\n")
	for n, _ := range servicos {
		fmt.Fprintf(os.Stderr, "  * %s\n", n)
	}
	os.Exit(2)
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}

	name := flag.Arg(0)
	demo, ok := servicos[name]
	if !ok {
		usage()
	}

	Config.Scope = escopos[name]
	Config.ClientId = valueOrFileContents(*clientId, *clientIdFile)
	Config.ClientSecret = valueOrFileContents(*secret, *secretFile)

	client := OAuthClient(Config)
	demo(client, flag.Args()[1:])
}

var (
	servicos = make(map[string]func(*http.Client, []string))
	escopos  = make(map[string]string)
)

func registrarServico(nome, escopo string, metodo func(c *http.Client, argv []string)) {
	if servicos[nome] != nil {
		panic(nome + " já registrado.")
	}
	servicos[nome] = metodo
	escopos[nome] = escopo
}
