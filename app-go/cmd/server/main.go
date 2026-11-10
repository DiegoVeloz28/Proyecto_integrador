package main

import (
	"app-go/internal/models" // 10. REUTILIZACIÓN DE PAQUETES
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Base de datos temporal en memoria
var baseDeDatos []*models.Imagen

func main() {
	http.HandleFunc("/registrar", registrarHandler)
	http.HandleFunc("/listar", listarHandler)

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func registrarHandler(w http.ResponseWriter, r *http.Request) {
	// Leemos los datos enviados por el cliente
	nombre := r.FormValue("nombre")
	formato := r.FormValue("formato")
	peso, _ := strconv.ParseFloat(r.FormValue("peso"), 64)

	// Usamos nuestro CONSTRUCTOR POO
	img, err := models.NuevaImagen(nombre, formato, peso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Devuelve error si falla validación
		return
	}

	// Guardamos en memoria
	baseDeDatos = append(baseDeDatos, img)
	fmt.Fprintf(w, "Éxito: Imagen registrada con el ID secreto %s", img.ObtenerID())
}

func listarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(baseDeDatos)
}
