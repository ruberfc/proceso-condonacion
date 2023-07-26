package main

import (
	"log"
	"os"
)

func LogRegistro() {
	// Crear un archivo de registro
	file, err := os.Create("Consulta.log")
	if err != nil {
		log.Fatal("No se pudo crear el archivo de registro:", err)
	}
	defer file.Close()

	// Configurar el logger para escribir en el archivo
	log.SetOutput(file)

	for i := 1; i <= 100; i++ {
		log.Println("Registro de log: ", i)
	}

	//log.Println("------")

	// Mensajes de registro que serán escritos en el archivo
	// log.Println("Este mensaje de registro será escrito en el archivo.")
	// log.Println("Otro mensaje de registro.")
}
