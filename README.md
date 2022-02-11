# Subir containers docker
```bash 
docker-compose up -d
```
# Acessar container da aplicação
```bash 
docker-compose exec app bash
```
# Rodar a aplicação
```bash 
go run cmd/main.go
```
# Entrar no container kafka
```bash 
docker-compose exec kafka bash
```
# Criar topicos no kafka
```bash 
kafka-topics --bootstrap-server=localhost:9092 --topic courses --create --partitions=3 --replication-factor=1
```
# Entrar no topico para enviar mensagem na fila
```bash 
kafka-console-producer --bootstrap-server=localhost:9092 --topic=courses
```
# Acessar Kafdrop 
```bash 
http://localhost:19000/
```