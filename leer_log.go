package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	// "strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type LogData struct {
	Timestamp      string
	Sede           string
	Codigo         string
	CicloAcademico int
	Anio           int
	Periodo        int
	NumeroCuota    int
	Cuota          float64
}

// func LeerLogAndCreateExcel

func main() {
	// Paso 1: Abrir el archivo de logs
	file, err := os.Open("Proceso.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Paso 4: Crear un nuevo archivo Excel
	xlsx := excelize.NewFile()

	// Definimos una estructura para el patrón de log que coincide con el formato proporcionado
	logPattern := regexp.MustCompile(`^(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}) Sede: (.*?), Codigo: (.*?) , Clico Academico: (\d+), Año: (\d+), Perido: (\d+), Numero Cuota: (\d+), Cuota: (\d+\.\d{2})$`)

	// Para almacenar los datos extraídos
	var logData []LogData

	// Paso 2 y 3: Procesar el archivo de logs y almacenar la información
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Paso 2: Procesar la información con la expresión regular
		matches := logPattern.FindStringSubmatch(line)
		if len(matches) == 9 {
			// Convertir valores numéricos en enteros o flotantes apropiados
			cicloAcademico, _ := strconv.Atoi(matches[4])
			anio, _ := strconv.Atoi(matches[5])
			periodo, _ := strconv.Atoi(matches[6])
			numeroCuota, _ := strconv.Atoi(matches[7])
			cuota, _ := strconv.ParseFloat(matches[8], 64)

			logData = append(logData, LogData{
				Timestamp:      matches[1],
				Sede:           matches[2],
				Codigo:         matches[3],
				CicloAcademico: cicloAcademico,
				Anio:           anio,
				Periodo:        periodo,
				NumeroCuota:    numeroCuota,
				Cuota:          cuota,
			})
		}
	}

	// Paso 4: Escribir los datos en el archivo Excel
	sheetName := "LogData"
	xlsx.NewSheet(sheetName)

	// Escribir encabezados en el archivo Excel
	headers := []string{"Timestamp", "Sede", "Codigo", "Ciclo Academico", "Anio", "Periodo", "Numero Cuota", "Cuota"}
	for col, header := range headers {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("%c1", 'A'+col), header)
	}

	// Escribir datos en el archivo Excel
	for i, data := range logData {
		rowNum := i + 2
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", rowNum), data.Timestamp)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", rowNum), data.Sede)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", rowNum), data.Codigo)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", rowNum), data.CicloAcademico)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", rowNum), data.Anio)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", rowNum), data.Periodo)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", rowNum), data.NumeroCuota)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", rowNum), data.Cuota)
	}

	// Guardar el archivo Excel
	if err := xlsx.SaveAs("output.xlsx"); err != nil {
		log.Fatal(err)
	}
}
