#!/bin/bash
echo "🔄 Running pre-commit hook..."
# Formatea el código Go
go fmt .
# Verifica estilo de código
golint .
# Realiza chequeos estáticos en el código
go vet .
# Ejecuta las pruebas
go test .