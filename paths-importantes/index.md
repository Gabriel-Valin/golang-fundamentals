## Variaveis e PATHS importantes


`go env` vai mostrar todas as envs utilizadas pelo go.


### Variaveis importantes


`GOPATH`: Caminho de instalação do go.

`GOMODCACHE`: Cache de packages.


## Paths


`$HOME/go` Possui 3 diretorios

- src
    -   pode ser usado para criacao de projetos
- pkg
    -   packages, compilacoes, etc...
- bin
    -   todos os binarios de utilizacao do go, ex: go install, go get..


### GO MOD

O comando `go mod` permite com que voce possa criar projetos em qualquer diretorio do seu computador e tornando nao obrigatorio a criacao do projeto na pasta src da instalacao do go.