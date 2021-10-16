# Obter Maximo usando Bi Directional streaming Grpc
Breve estudo sobre funcionamento do Bi Directional streaming
## Iniciar aplicação
- Execute : ./server.sh para iniciar o servidor
- Execute : ./client.sh para iniciar a consulta
#   Configurando ambiente
- 
    - go mod download
    - pamac install protobuf || sudo pacman -S protobuf
    - more information about : https://grpc.io/docs/languages/go/quickstart/
    - protoc src/maximum_proto/maximum.proto --go_out=. --go-grpc_out=. ou ./gen.sh para gerar os codigos
    - chmod +x arquivos.sh
    - em caso de problemas com compilador set as seguintes variaveis:
        - export GO_PATH=~/go
        - export PATH=$PATH:/$GO_PATH/bin