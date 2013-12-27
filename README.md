# Recursos

## Upload de arquivos para o Google Drive

Permite fazer upload de múltiplos arquivos para uma pasta, criando a pasta e enviando todos os arquivos especificados nos parâmetros.

### Exemplo de uso:

    Google_API.exe upload 13.2.2.300 C:\build\MondeCliente.exe "C:\Pasta com espaço\MondeServidor.exe"

O comando acima irá fazer upload dos arquivos MondeCliente.exe e MondeServidor.exe para uma pasta 13.2.2.300 no Google Drive.

### Limitações e recursos a serem implementados

- O ID da pasta "parent" onde as pastas são criadas para fazer o upload está hardcoded, precisa ser configurável ou o ideal é receber nos parametros: upload Pasta1\subpasta e o aplicativo criar a estrutura no drive caso não exista
- Ao enviar o arquivo, caso a pasta já exista, o aplicativo irá criar uma pasta com o mesmo nome (o Google Drive permite isso), o ideal era detectar e fazer upload para a mesma pasta.

# Desenvolvimento

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

Para fazer a autenticação é necessário gerar um ClientId e um ClientSecret no console da API: [https://code.google.com/apis/console/ ](https://code.google.com/apis/console/)

#### Criar arquivos de configuração

Após gerar os Ids da API coloque cada id respectivamente nos arquivos ClientId.dat e ClientSecret.dat na mesma pasta do executável.

#### Autenticação

Após fazer a configuração inicial, execute o aplicativo e ele irá redirecionar para o Browser e salvar o Token de autorização em disco.


