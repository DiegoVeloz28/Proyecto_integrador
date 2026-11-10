package models

import "testing"

func TestNuevaImagen(t *testing.T) {
	img, err := NuevaImagen("entrega_sur.jpg", "jpg", 2.5)
	if err != nil {
		t.Fatalf("no se esperaba error y se obtuvo: %v", err)
	}
	if img.ObtenerID() == "" {
		t.Fatal("el ID generado no puede estar vacío")
	}
	if img.Meta.PesoMB != 2.5 {
		t.Fatalf("peso esperado 2.5, obtenido %v", img.Meta.PesoMB)
	}
}

func TestValidacionesError(t *testing.T) {
	if _, err := NuevaImagen("", "jpg", 1.0); err == nil {
		t.Fatal("se esperaba error por nombre vacío")
	}
	if _, err := NuevaImagen("foto.jpg", "jpg", -3); err == nil {
		t.Fatal("se esperaba error por peso menor o igual a cero")
	}
	if _, err := NuevaImagen("documento.pdf", "pdf", 1.2); err == nil {
		t.Fatal("se esperaba error por formato no permitido")
	}
}

func TestCambiarNombre(t *testing.T) {
	img, _ := NuevaImagen("original.png", "png", 1.5)
	img.CambiarNombre("actualizada.png")
	if img.Nombre != "actualizada.png" {
		t.Fatalf("se esperaba 'actualizada.png', se obtuvo '%s'", img.Nombre)
	}
}
