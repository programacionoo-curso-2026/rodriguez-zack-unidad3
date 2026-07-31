package dao

import (
	"database/sql"
	"fmt"
	"log"
)

// DocenteDAO maneja las operaciones CRUD para Docente
type DocenteDAO struct {
	db *sql.DB
}

// NewDocenteDAO crea una nueva instancia de DocenteDAO
func NewDocenteDAO(db *sql.DB) *DocenteDAO {
	return &DocenteDAO{db: db}
}

// CreateTable crea la tabla de docentes si no existe
func (d *DocenteDAO) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS docentes (
		id TEXT PRIMARY KEY,
		nombre TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		departamento TEXT,
		cargo TEXT,
		anios_antiguedad INTEGER DEFAULT 0
	);`

	_, err := d.db.Exec(query)
	if err != nil {
		return fmt.Errorf("error al crear tabla docentes: %w", err)
	}

	log.Println("Tabla docentes creada/verificada exitosamente")
	return nil
}
