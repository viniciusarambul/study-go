# iniciar a aplicação com docker
- docker-compose up -d

# acessar container da aplicação
- docker-compose exec app bash

# rodar a aplicação em go
- go run cmd/main.go

# entrar no container kafka
- docker-compose exec kafka bash

# criar topicos no kafka
- kafka-topics --bootstrap-server=localhost:9092 --topic courses --create --partitions=3 --replication-factor=1

# entrar no topico para enviar mensagem na fila
- kafka-console-producer --bootstrap-server=localhost:9092 --topic=courses