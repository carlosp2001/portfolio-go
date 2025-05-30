# Variables
SASS = sass
SASS_SRC = frontend/scss
SASS_DEST = frontend/css
GO_APP = backend/main.go
PORT = 8080

# Comandos
.PHONY: help dev build run sass clean

help:
	@echo "Comandos disponibles:"
	@echo "  make dev      - Ejecuta Sass en watch mode y el servidor Go"
	@echo "  make sass     - Compila archivos Sass una vez"
	@echo "  make build    - Compila el binario Go"
	@echo "  make run      - Ejecuta la app compilada"
	@echo "  make clean    - Limpia el binario"

sass:
	$(SASS) $(SASS_SRC):$(SASS_DEST)

dev:
	@echo "Iniciando entorno de desarrollo..."
	@$(SASS) --watch $(SASS_SRC):$(SASS_DEST) &
	go run $(GO_APP)

build:
	go build -o portfolio backend/main.go

run:
	./portfolio

clean:
	rm -f portfolio
