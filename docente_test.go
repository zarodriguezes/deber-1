package docente

import "testing"

// TestValidarEmailPrivado prueba el método privado validarEmail()
// con distintos formatos de email, válidos e inválidos.
func TestValidarEmailPrivado(t *testing.T) {
	casos := []struct {
		nombre string
		email  string
		valido bool
	}{
		{"email valido institucional", "juan.perez@uide.edu.ec", true},
		{"email valido gmail", "ana@gmail.com", true},
		{"sin arroba", "anagmail.com", false},
		{"sin punto despues de la arroba", "ana@gmailcom", false},
		{"vacio", "", false},
		{"arroba al final", "ana@", false},
		{"arroba al inicio", "@gmail.com", false},
		{"con espacios pero formato valido", "  ana@gmail.com  ", true},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			d := &Docente{Email: c.email}
			resultado := d.validarEmail()
			if resultado != c.valido {
				t.Errorf("validarEmail(%q) = %v; se esperaba %v", c.email, resultado, c.valido)
			}
		})
	}
}

// TestNormalizarNombrePrivado prueba el método privado normalizarNombre()
// con nombres en distintos formatos de mayúsculas/minúsculas.
func TestNormalizarNombrePrivado(t *testing.T) {
	casos := []struct {
		nombre   string
		entrada  string
		esperado string
	}{
		{"minusculas", "juan perez", "Juan Perez"},
		{"mayusculas", "JUAN PEREZ", "Juan Perez"},
		{"mixto", "jUaN pErEz", "Juan Perez"},
		{"una sola palabra", "maria", "Maria"},
		{"espacios extra", "  ana   lucia  ", "Ana Lucia"},
		{"tres palabras", "jose luis vera", "Jose Luis Vera"},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			d := &Docente{Nombre: c.entrada}
			resultado := d.normalizarNombre()
			if resultado != c.esperado {
				t.Errorf("normalizarNombre(%q) = %q; se esperaba %q", c.entrada, resultado, c.esperado)
			}
		})
	}
}

// TestAgregarEvaluacionInterna verifica que el método privado
// agregarEvaluacionInterna agregue, rechace duplicados y rechace vacíos.
func TestAgregarEvaluacionInterna(t *testing.T) {
	d := NuevoDocente("D001", "juan perez", "juan@uide.edu.ec", "Sistemas", "Titular")

	// Caso 1: agregar una evaluación válida
	err := d.agregarEvaluacionInterna("EVAL-2026-01")
	if err != nil {
		t.Fatalf("no se esperaba error al agregar evaluación: %v", err)
	}
	if len(d.evaluaciones) != 1 {
		t.Errorf("se esperaba 1 evaluación registrada, hay %d", len(d.evaluaciones))
	}
	if d.evaluaciones[0] != "EVAL-2026-01" {
		t.Errorf("evaluación registrada incorrecta: %s", d.evaluaciones[0])
	}

	// Caso 2: agregar una evaluación duplicada debe fallar
	err = d.agregarEvaluacionInterna("EVAL-2026-01")
	if err == nil {
		t.Error("se esperaba error al agregar una evaluación duplicada, pero no ocurrió")
	}

	// Caso 3: agregar una evaluación vacía debe fallar
	err = d.agregarEvaluacionInterna("   ")
	if err == nil {
		t.Error("se esperaba error al agregar una evaluación vacía, pero no ocurrió")
	}

	// El slice privado no debió crecer con los intentos fallidos
	if len(d.evaluaciones) != 1 {
		t.Errorf("se esperaba que el total siguiera siendo 1, hay %d", len(d.evaluaciones))
	}
}
