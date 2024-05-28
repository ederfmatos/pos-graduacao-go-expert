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