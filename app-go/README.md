# Prototipo Funcional SIGRIC (`app-go`) — Semana 10 (POO)

Sistema cliente-servidor en Go para el registro y listado de imágenes corporativas con validaciones de negocio, composición, encapsulamiento y manejo explícito de errores.

## Estructura del proyecto

```text
app-go/
├── go.mod
├── README.md
├── internal/
│   └── models/
│       ├── imagen.go       # Structs Imagen y Metadata, constructor NuevaImagen, ObtenerID y CambiarNombre
│       └── imagen_test.go  # Pruebas unitarias automatizadas
└── cmd/
    ├── server/
    │   └── main.go         # Servidor HTTP en :8080 (/registrar y /listar)
    └── client/
        └── main.go         # Cliente interactivo de consola
```

## Cómo ejecutar paso a paso

### 1. Ejecutar las pruebas unitarias
```bash
go test -v ./internal/models
```

### 2. Iniciar el servidor (Terminal 1)
```bash
go run ./cmd/server
```

### 3. Iniciar el cliente (Terminal 2)
```bash
go run ./cmd/client
```
