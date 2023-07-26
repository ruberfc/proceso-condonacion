package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {

	// Crear un archivo de registro
	file, err := os.Create("Proceso.log")
	if err != nil {
		fmt.Println("No se pudo crear el archivo de registro:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Configurar el logger para escribir en el archivo
	log.SetOutput(file)

	//excelFileName := "Estudiantes_proceso.xlsx" // Reemplaza esto con la ruta real de tu archivo de Excel
	excelFileName := "REPORTE DE IMPEDIDOS DEL 2015 AL 2022 - ALCTUALIZADO AL 06062023 - PARA CONDONAR.xlsx"

	// Abrir el archivo de Excel
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		os.Exit(1)
	}

	// Obtener la primera hoja del archivo
	sheet := xlFile.Sheets[0]

	// Leer y mostrar los datos de una sola columna
	// columnIndex := 0 // Índice de la columna que deseas leer (0 para la primera columna, 1 para la segunda, y así sucesivamente)

	colSede := 0
	colCodigo := 1
	colCicloAcademico := 7
	colYear := 8
	colPeriodo := 9
	colNumCuota := 10
	colCuota := 11

	// Iterar sobre las filas y obtener el valor de la columna deseada
	for _, row := range sheet.Rows {
		// if columnaDeseada < len(row.Cells) {
		// 	cell := row.Cells[columnaDeseada]
		// 	valor := cell.String()
		// 	fmt.Println(valor)
		// }
		cellSede := row.Cells[colSede].String()
		cellCodigo := row.Cells[colCodigo].String()
		cellCicloAcademico := row.Cells[colCicloAcademico].String()
		cellYear := row.Cells[colYear].String()
		cellPeriodo := row.Cells[colPeriodo].String()
		cellNumCuota := row.Cells[colNumCuota].String()
		cellCuota := row.Cells[colCuota].String()

		log.Printf("Sede: %s, Codigo: %s , Clico Academico: %s, Año: %s, Perido: %s, Numero Cuota: %s, Cuota: %s \n", cellSede, cellCodigo, cellCicloAcademico, cellYear, cellPeriodo, cellNumCuota, cellCuota)
	}

	/*

		for _, sheet := range xlFile.Sheets {


			// totalRegistros := 0


			for _, row := range sheet.Rows {
				if columnIndex < len(row.Cells) {
					cell := row.Cells[columnIndex]
					text := cell.String()
					// fmt.Println(text)

					log.Println("Codigo de estudiante:", text)

					// est, proceso := GetRespuestaProceso(text)
					// if proceso == "error_conexion" {
					// 	fmt.Println("error_conexion")
					// } else if proceso == "consulta_vacia" {
					// 	fmt.Println("consulta_vacia")
					// } else if proceso == "consultado" {

					// } else {
					// 	fmt.Println("error_" + proceso)
					// }
				}

				// totalRegistros = totalRegistros + 1

			}

			//fmt.Println(totalRegistros)
		}
	*/
}
