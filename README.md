[![Build Status](https://travis-ci.org/cadena-monde/google-api.png?branch=master)](https://travis-ci.org/cadena-monde/google-api)

# Recursos

## Upload de arquivos para o Google Drive

Permite fazer upload de múltiplos arquivos criando a estrutura de pastas e enviando todos os arquivos especificados nos parâmetros.

### Exemplo de uso:

    google-api.exe upload Pasta\Subpasta C:\Arquivo1.exe "C:\Pasta com espaço\Arquivo2.exe"

O comando acima irá fazer upload dos arquivos Arquivo1.exe e Arquivo2.exe para uma Pasta\Subpasta, mantendo a estrutura de diretórios no Google Drive.

## Download de determinado arquivo mais recente

Permite fazer o download de um determinado arquivo no Google Drive com a versão mais recente desse arquivo.

### Exemplo de uso:

    google-api.exe baixar_arquivo_mais_recente Arquivo.7z C:\Download\Arquivo.7z

O comando acima irá fazer o download do Arquivo.7z mais recente dentro do Google Drive para a pasta local C:\Download com o nome de Arquivo.7z.

## Dependências

### google-api-go-client

[https://code.google.com/p/google-api-go-client/wiki/GettingStarted](https://code.google.com/p/google-api-go-client/wiki/GettingStarted)

### Drive v2:

    go get code.google.com/p/google-api-go-client/drive/v2

### goauth2

[https://code.google.com/p/goauth2/](https://code.google.com/p/goauth2/)

    go get code.google.com/p/goauth2/oauth

## Autenticação

#### Gerar ClientId e ClientSecret

Para fazer a autenticação é necessário gerar um ClientId e um ClientSecret no console da API: [https://code.google.com/apis/console/ ](https://code.google.com/apis/console/).
Lembrando que, ao gerar a autenticação, é necessário escolher a opção "Installed application", pois somente dessa forma o Client Secret será gerado também.

#### Criar arquivos de configuração

Após gerar os Ids da API coloque cada id respectivamente nos arquivos ClientId.dat e ClientSecret.dat na mesma pasta do executável.

#### Autenticação

Após fazer a configuração inicial, execute o aplicativo e ele irá redirecionar para o Browser e salvar o Token de autorização em disco.


