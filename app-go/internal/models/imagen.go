package models

import (
	"errors"
	"strings"
	"time"
)

// 1. COMPOSICIÓN: Struct independiente para agrupar metadatos
type Metadata struct {
	Resolucion string
	PesoMB     float64
}

// 2. STRUCT PRINCIPAL
type Imagen struct {
	// 3. CAMPO NO EXPORTADO (Encapsulamiento, inicia con minúscula)
	id string

	// 3. CAMPOS EXPORTADOS (Inician con mayúscula)
	Nombre  string
	Formato string
	Meta    Metadata // 4. COMPOSICIÓN en uso
}

// 5. FUNCIÓN CONSTRUCTORA (Crea el objeto con estado válido)
func NuevaImagen(nombre, formato string, peso float64) (*Imagen, error) {
	// 6. VALIDACIÓN DE DATOS y 7. MANEJO EXPLÍCITO DE ERRORES
	if nombre == "" {
		return nil, errors.New("error: el nombre de la imagen no puede estar vacío")
	}
	if peso <= 0 {
		return nil, errors.New("error: el peso debe ser mayor a cero")
	}
	if !strings.HasSuffix(nombre, ".jpg") && !strings.HasSuffix(nombre, ".png") {
		return nil, errors.New("error: solo se permiten formatos .jpg o .png")
	}

	// Retorna el objeto creado
	return &Imagen{
		id:      time.Now().Format("150405"), // ID simple basado en la hora
		Nombre:  nombre,
		Formato: formato,
		Meta: Metadata{
			Resolucion: "1920x1080", // Dato simulado por simplicidad
			PesoMB:     peso,
		},
	}, nil
}

// 8. ENCAPSULAMIENTO Y MÉTODOS (Getter para acceder al campo privado 'id')
func (i *Imagen) ObtenerID() string {
	return i.id
}

// 9. RECEPTOR POR PUNTERO (Modifica el estado interno de la estructura)
func (i *Imagen) CambiarNombre(nuevoNombre string) {
	i.Nombre = nuevoNombre
}
