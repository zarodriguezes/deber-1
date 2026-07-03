package docente_test

import (
	"testing"

	"docenteevaluacion/docente"
)

// TestAccesoPublicoDocente demuestra que desde fuera del paquete `docente`
// solo se puede interactuar con Docente a través de su API pública.
func TestAccesoPublicoDocente(t *testing.T) {
	d := docente.NuevoDocente("D002", "maria gonzalez", "maria@uide.edu.ec", "Redes", "Titular")

	if d.GetID() != "D002" {
		t.Errorf("GetID() = %s; se esperaba D002", d.GetID())
	}

	// El nombre llega normalizado porque el constructor público
	// invoca internamente al método privado normalizarNombre().
	if d.GetNombre() != "Maria Gonzalez" {
		t.Errorf("GetNombre() = %s; se esperaba 'Maria Gonzalez'", d.GetNombre())
	}

	if d.GetEmail() != "maria@uide.edu.ec" {
		t.Errorf("GetEmail() = %s; se esperaba 'maria@uide.edu.ec'", d.GetEmail())
	}

	// La validación del email solo es accesible a través
	// del método público EsEmailValido().
	if !d.EsEmailValido() {
		t.Error("se esperaba que el email fuera válido")
	}

	// El registro de evaluaciones solo es posible a través
	// del método público AgregarEvaluacion().
	if err := d.AgregarEvaluacion("EVAL-EXT-01"); err != nil {
		t.Fatalf("no se esperaba error al agregar evaluación: %v", err)
	}

	evaluaciones := d.GetEvaluaciones()
	if len(evaluaciones) != 1 || evaluaciones[0] != "EVAL-EXT-01" {
		t.Errorf("GetEvaluaciones() = %v; se esperaba ['EVAL-EXT-01']", evaluaciones)
	}

	// ===================================================================
	// Las siguientes líneas demuestran encapsulación: si se descomentan,
	// el código NO COMPILA, porque desde el paquete docente_test (externo)
	// no hay acceso a identificadores privados (en minúscula) del
	// paquete docente.
	// ===================================================================
	//
	// d.validarEmail()                  // ERROR: d.validarEmail undefined (no exportado)
	// d.normalizarNombre()              // ERROR: d.normalizarNombre undefined (no exportado)
	// d.agregarEvaluacionInterna("X")   // ERROR: d.agregarEvaluacionInterna undefined (no exportado)
	// _ = d.evaluaciones                // ERROR: d.evaluaciones undefined (no exportado)
}
