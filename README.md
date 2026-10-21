# gpass

Gerador de senhas para a linha de comando, escrito em Go. Usa `crypto/rand`
(gerador criptograficamente seguro) e não depende de bibliotecas externas.

## Instalação

Requer Go instalado.

```
go install github.com/luismartoranomartini/gpass@latest
```

O binário é instalado em `$(go env GOPATH)/bin` (padrão: `~/go/bin`).
Confira se esse diretório está no seu `PATH`.

## Uso

```
gpass [opções]
```

| Opção | Padrão | Descrição |
|---|---|---|
| `-n caracteres` | `16` | Tamanho de cada senha |
| `-c int` | `1` | Quantidade de senhas |
| `-s` | desligada | Inclui símbolos (`!@#$%^&*()-_=+[]{}<>?`) |
| `-h` | | Mostra a ajuda |

Sem `-s`, as senhas usam letras maiúsculas, minúsculas e dígitos.

## Exemplos

Uma senha de 16 caracteres:

```
$ gpass
5dNR3i9DrHOSwCXc
```

Três senhas de 10 caracteres, com símbolos:

```
$ gpass -c 3 -n 10 -s
[+N83?nTX]
7lqe1vdTZ+
8=66UW_Rss
```

Copiar para a área de transferência (Wayland, com `wl-clipboard`):

```
$ gpass -n 24 | wl-copy
```

## Códigos de saída

| Código | Significado |
|---|---|
| `0` | Sucesso (inclui `-h`) |
| `1` | Falha de execução |
| `2` | Uso inválido (flag desconhecida, valor inválido, argumento inesperado) |

As senhas vão para a saída padrão, uma por linha. Mensagens de erro e de
ajuda vão para a saída de erro, então o uso em pipes não é contaminado.

## Desenvolvimento

```
go build -o gpass .
go vet ./...
go test ./...
```

## Licença

Defina a licença do projeto (por exemplo, MIT) e adicione o arquivo `LICENSE`.
