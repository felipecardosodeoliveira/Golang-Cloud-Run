# Golang-Cloud-Run

Um sistema de consulta de temperatura por cep.

## Comu usar

### Pré-requisitos

- Go (>= 1.22) instalado
- Docker (opcional)

### Executando localmente

1. Clone o repositório:

    ```bash
    git clone git@github.com:felipecardosodeoliveira/Golang-Cloud-Run.git
    cd Golang-Cloud-Run

2. Compile e execute o programa:

    go run cmd/main.go 
    Server is running 8080 

3. Fazer a requisição que está no arquivo doc/api.http. O resultado será o seguinte:
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 06 Nov 2024 00:14:31 GMT
    Content-Length: 59

    {
    "temp_C": 20.2,
    "temp_F": 68.36,
    "temp_K": 293.34999999999997
    } 


### Executando com Docker

Você pode construir uma imagem docker e executar a aplicação.

1. Build da imagem:

    docker build -t golang_cloud_run .

2. Execute a imagem:

    docker run -p 8080:8080  golang_cloud_run 

3. Fazer a requisição que está no arquivo doc/api.http. O resultado será o seguinte:
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 06 Nov 2024 00:14:31 GMT
    Content-Length: 59

    {
    "temp_C": 20.2,
    "temp_F": 68.36,
    "temp_K": 293.34999999999997
    } 

### Link cloud

https://golang-cloud-run-5ahvc4sr5a-uc.a.run.app/cep/81490420