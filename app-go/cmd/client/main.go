package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	for {
		fmt.Println("\n--- SISTEMA DE IMÁGENES (CLIENTE) ---")
		fmt.Println("1. Registrar imagen nueva")
		fmt.Println("2. Listar imágenes")
		fmt.Println("3. Salir")
		fmt.Print("Elija una opción: ")

		var opcion string
		fmt.Scanln(&opcion)

		if opcion == "1" {
			registrar()
		} else if opcion == "2" {
			listar()
		} else if opcion == "3" {
			break
		}
	}
}

func registrar() {
	var nombre, formato, peso string
	fmt.Print("Nombre del archivo (ej. foto.jpg): ")
	fmt.Scanln(&nombre)
	fmt.Print("Formato (jpg/png): ")
	fmt.Scanln(&formato)
	fmt.Print("Peso en MB (ej. 2.5): ")
	fmt.Scanln(&peso)

	// Enviamos los datos al servidor (Comunicación Cliente-Servidor)
	datos := url.Values{"nombre": {nombre}, "formato": {formato}, "peso": {peso}}
	resp, err := http.PostForm("http://localhost:8080/registrar", datos)
	if err != nil {
		fmt.Println("Error conectando al servidor")
		return
	}
	defer resp.Body.Close()

	respuesta, _ := io.ReadAll(resp.Body)
	fmt.Println("Servidor dice:", string(respuesta))
}

func listar() {
	resp, err := http.Get("http://localhost:8080/listar")
	if err != nil {
		fmt.Println("Error conectando al servidor")
		return
	}
	defer resp.Body.Close()
	respuesta, _ := io.ReadAll(resp.Body)
	fmt.Println("Datos en el servidor:\n", string(respuesta))
}
