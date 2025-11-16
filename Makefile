# Variables
BINARY_NAME=gup
BUILD_DIR=build
MAIN_FILE=main.go

# Colores para output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
NC=\033[0m # No Color

.PHONY: help build clean deps dev run test install

# Comando por defecto
all: deps build

# Mostrar ayuda
help:
	@echo "$(YELLOW)GUP - Makefile Commands$(NC)"
	@echo ""
	@echo "$(GREEN)Available commands:$(NC)"
	@echo "  build     - Compilar el binario"
	@echo "  clean     - Limpiar archivos generados"
	@echo "  deps      - Descargar e instalar dependencias"
	@echo "  dev       - Ejecutar en modo desarrollo"
	@echo "  run       - Ejecutar sin compilar"
	@echo "  test      - Ejecutar tests"
	@echo "  install   - Instalar el binario en el sistema"
	@echo "  help      - Mostrar esta ayuda"

# Descargar dependencias
deps:
	@echo "$(YELLOW)📦 Descargando dependencias...$(NC)"
	go mod tidy
	go mod download
	@echo "$(GREEN)✅ Dependencias instaladas correctamente$(NC)"

# Compilar el binario
build: deps
	@echo "$(YELLOW)🔨 Compilando $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "$(GREEN)✅ Binario compilado: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

# Limpiar archivos generados
clean:
	@echo "$(YELLOW)🧹 Limpiando archivos...$(NC)"
	rm -rf $(BUILD_DIR)
	go clean
	@echo "$(GREEN)✅ Archivos limpiados$(NC)"

# Ejecutar en modo desarrollo (con hot reload básico)
dev: deps
	@echo "$(YELLOW)🔄 Modo desarrollo - Ejecutando...$(NC)"
	go run $(MAIN_FILE)

# Ejecutar sin compilar
run: deps
	@echo "$(YELLOW)▶️  Ejecutando CliGO...$(NC)"
	go run $(MAIN_FILE) $(ARGS)

# Ejecutar tests
test:
	@echo "$(YELLOW)🧪 Ejecutando tests...$(NC)"
	go test -v ./...
	@echo "$(GREEN)✅ Tests completados$(NC)"

# Instalar en el sistema (requiere sudo)
install: build
	@echo "$(YELLOW)📥 Instalando $(BINARY_NAME) en /usr/local/bin...$(NC)"
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	sudo chmod +x /usr/local/bin/$(BINARY_NAME)
	@echo "$(GREEN)✅ $(BINARY_NAME) instalado correctamente$(NC)"
	@echo "$(GREEN)Ahora puedes usar 'gup' desde cualquier lugar$(NC)"

# Compilar para múltiples arquitecturas
build-all: deps
	@echo "$(YELLOW)🔨 Compilando para múltiples arquitecturas...$(NC)"
	@mkdir -p $(BUILD_DIR)
	
	# Linux AMD64
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_FILE)
	
	# Linux ARM64
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_FILE)
	
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_FILE)
	
	# macOS ARM64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_FILE)
	
	@echo "$(GREEN)✅ Binarios compilados para múltiples arquitecturas$(NC)"
	@ls -la $(BUILD_DIR)/

# Verificar el código
lint:
	@echo "$(YELLOW)🔍 Verificando código...$(NC)"
	go fmt ./...
	go vet ./...
	@echo "$(GREEN)✅ Código verificado$(NC)"

# Mostrar información del sistema
info:
	@echo "$(YELLOW)ℹ️  Información del sistema:$(NC)"
	@echo "Go version: $(shell go version)"
	@echo "OS: $(shell uname -s)"
	@echo "Architecture: $(shell uname -m)"
	@echo "Current directory: $(PWD)"