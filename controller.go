package main

import (
	"context"
	"database/sql"
)

var contx = context.Background()

func GetRespuestaProceso(num_di string) (Estudiante, string) {
	estudiante := Estudiante{}

	db, err := CreateConnection()
	if err != nil {
		return estudiante, "error_conexion"
	}
	defer db.Close()

	query := `exec Sp_	g @NumDI`
	row := db.QueryRowContext(contx, query, sql.Named("NumDI", num_di))

	estudiante.NumDI = num_di
	err = row.Scan(&estudiante.Rpt)
	if err == sql.ErrNoRows {
		return estudiante, "consulta_vacia"
	}
	if err != nil {
		return estudiante, err.Error()
	}

	return estudiante, "consultado"
}
