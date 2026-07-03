package docente

import (
	"errors"
	"strings"
)

type Docente struct {
	ID           string
	Nombre       string
	Email        string
	Departamento string
	Cargo        string
	// Campo privado
	evaluaciones []string
}

// Constructor público
func NuevoDocente(id, nombre, email, depto, cargo string) *Docente {
	d := &Docente{
		ID:           id,
		Nombre:       nombre,
		Email:        email,
		Departamento: depto,
		Cargo:        cargo,
		evaluaciones: []string{},
	}
	// Se normaliza el nombre automáticamente al construir el docente
	d.Nombre = d.normalizarNombre()
	return d
}

// ===================== Métodos públicos =====================

func (d *Docente) GetID() string {
	return d.ID
}

func (d *Docente) GetNombre() string {
	return d.Nombre
}

func (d *Docente) GetEmail() string {
	return d.Email
}

// EsEmailValido expone hacia afuera el resultado de validarEmail(),
// que es privado, sin exponer el método en sí.
func (d *Docente) EsEmailValido() bool {
	return d.validarEmail()
}

// AgregarEvaluacion es la puerta pública para registrar una evaluación.
// Internamente delega en el método privado agregarEvaluacionInterna.
func (d *Docente) AgregarEvaluacion(idEvaluacion string) error {
	return d.agregarEvaluacionInterna(idEvaluacion)
}

// GetEvaluaciones devuelve una COPIA de las evaluaciones registradas,
// para no exponer directamente el slice privado original.
func (d *Docente) GetEvaluaciones() []string {
	copia := make([]string, len(d.evaluaciones))
	copy(copia, d.evaluaciones)
	return copia
}

// ===================== Métodos privados =====================

// validarEmail verifica que el correo tenga un formato básico válido:
// debe contener "@" y, después de la "@", debe existir un ".".
func (d *Docente) validarEmail() bool {
	email := strings.TrimSpace(d.Email)

	posArroba := strings.Index(email, "@")
	if posArroba <= 0 || posArroba == len(email)-1 {
		return false
	}

	dominio := email[posArroba+1:]
	if !strings.Contains(dominio, ".") {
		return false
	}

	return true
}

// normalizarNombre convierte el nombre a formato título
// (primera letra de cada palabra en mayúscula, resto en minúscula).
func (d *Docente) normalizarNombre() string {
	palabras := strings.Fields(strings.ToLower(d.Nombre))

	for i, palabra := range palabras {
		runas := []rune(palabra)
		if len(runas) > 0 {
			runas[0] = []rune(strings.ToUpper(string(runas[0])))[0]
		}
		palabras[i] = string(runas)
	}

	return strings.Join(palabras, " ")
}

// agregarEvaluacionInterna agrega una evaluación a la lista privada,
// validando que no esté vacía ni duplicada.
func (d *Docente) agregarEvaluacionInterna(idEvaluacion string) error {
	idEvaluacion = strings.TrimSpace(idEvaluacion)

	if idEvaluacion == "" {
		return errors.New("el id de evaluación no puede estar vacío")
	}

	for _, e := range d.evaluaciones {
		if e == idEvaluacion {
			return errors.New("la evaluación ya fue registrada previamente")
		}
	}

	d.evaluaciones = append(d.evaluaciones, idEvaluacion)
	return nil
}
