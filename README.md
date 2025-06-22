![Badge Concluído](http://img.shields.io/static/v1?label=STATUS&message=CONCLUÍDO&color=GREEN&style=for-the-badge)
![Badge Kubernetes](http://img.shields.io/static/v1?label=ARQUITETURA_HEXAGONAL&message=go1.16.15&color=blue&style=for-the-badge)
![Badge FullCycle](http://img.shields.io/static/v1?label=FULLCYCLE&message=3.0&color=orange&style=for-the-badge)
![Badge PDI](http://img.shields.io/static/v1?label=PDI&message=LOGCOMEX&color=purple&style=for-the-badge)


### Subindo Docker
```
docker compose up -d
```

### Ingressando no container
```
docker exec -it appproduct bash
```

### Instalando dependências
1. Ingressar no container
2. Executar ```go mod tidy```

### Executando testes unitários (dentro do container)
```
go test // for all tests
go test -run TestProductService_Get // or run a specific test
```

### Cobra CLI
```
cobra-cli init
cobra-cli add command
```

### SQLITE
Para conectar em uma base
```
sqlite3 "nome do banco"
```

Para sair do terminal
```
.exit ou .quit
```
