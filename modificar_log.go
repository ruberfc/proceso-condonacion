package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Abre el archivo de logs
	file, err := os.Open("Proceso.log")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crea un nuevo archivo para escribir las líneas modificadas
	outputFile, err := os.Create("archivo_sin_fecha_hora.log")
	if err != nil {
		fmt.Println("Error al crear el archivo de salida:", err)
		return
	}
	defer outputFile.Close()

	// Procesar el archivo de logs línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Quitar la fecha y hora de la línea original
		lineWithoutTimestamp := strings.TrimSpace(line[19:])

		// Escribir la línea modificada en el nuevo archivo
		fmt.Fprintln(outputFile, lineWithoutTimestamp)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	fmt.Println("Proceso completado. Se ha creado el archivo sin la fecha y hora.")
}
