# Comandos

**Executar testes com cobertura**
```bash
 go test -coverprofile=coverage.out ./testing
```

**Gerar relatorio de cobertura**
```bash
 go tool cover -html=coverage.out 
```

**Executar testes de mutação**
```bash
 go test -fuzz=. -fuzztime=5s
```

# SqlC
****
```bash
 migrate create -ext=sql -dir=sql/migrations -seq init
```
**Run migrations**
```bash
 migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose up
```

**Down migrations**
```bash
 migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose down
```