#!/bin/bash
echo "🔄 Running pre-push hook..."
# Verifica si el código está correctamente formateado
go fmt ./...
# Realiza chequeos estáticos en el código
golint ./...
# Realiza análisis estático de posibles errores
go vet ./...
# Ejecuta las pruebas
go test ./...
if [ $? -ne 0 ]; then
  echo "❌ Las pruebas fallaron. Por favor, revisa los errores antes de hacer push."
  exit 1
fi