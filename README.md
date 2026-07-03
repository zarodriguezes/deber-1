# DocenteEvaluacion — Encapsulación en Go

Proyecto académico para la asignatura de Programación (UIDE), enfocado en
demostrar **encapsulación** en Go mediante campos y métodos privados,
pruebas internas y pruebas externas.

**Autores:** José Sebastián Tumbaco Aguirre, Miguel Ángel Menéndez Flores

## Estructura del proyecto

```
DocenteEvaluacion/
├── docente/
│   ├── docente.go              # Estructura Docente + métodos públicos y privados
│   ├── docente_test.go         # Pruebas internas (package docente)
│   └── docente_externo_test.go # Pruebas externas (package docente_test)
├── go.mod
└── README.md
```

## Métodos privados agregados

| Método | Descripción |
|---|---|
| `validarEmail()` | Verifica que el email contenga `@` y un `.` después de la arroba. |
| `normalizarNombre()` | Convierte el nombre a formato título (cada palabra con su primera letra en mayúscula). |
| `agregarEvaluacionInterna(idEvaluacion string)` | Agrega un ID de evaluación al slice privado `evaluaciones`, validando que no esté vacío ni duplicado. |

Estos tres métodos son **privados** (empiezan con minúscula), por lo que solo
son accesibles desde archivos que pertenezcan al paquete `docente`. Para que
sigan siendo útiles desde afuera, se agregaron wrappers públicos:
`EsEmailValido()`, `GetEmail()`, `AgregarEvaluacion()` y `GetEvaluaciones()`.
El constructor `NuevoDocente` también usa `normalizarNombre()` internamente,
así que todo nombre queda normalizado desde su creación.

## Comandos ejecutados y resultados

```bash
# Compilar el módulo completo
$ go build ./...
# (sin salida = compilación exitosa)

# Análisis estático
$ go vet ./...
# (sin salida = sin advertencias)

# Ejecutar todas las pruebas (internas + externas) en modo verbose
$ go test ./docente -v
PASS
ok      docenteevaluacion/docente      0.001s

# Ejecutar pruebas con cobertura
$ go test ./docente -v -cover
PASS
coverage: 100.0% of statements
ok      docenteevaluacion/docente      0.004s

# Detalle de cobertura por función
$ go test ./docente -coverprofile=coverage.out
$ go tool cover -func=coverage.out
docente.go:19:   NuevoDocente               100.0%
docente.go:35:   GetID                      100.0%
docente.go:39:   GetNombre                  100.0%
docente.go:43:   GetEmail                   100.0%
docente.go:49:   EsEmailValido              100.0%
docente.go:55:   AgregarEvaluacion          100.0%
docente.go:61:   GetEvaluaciones            100.0%
docente.go:71:   validarEmail               100.0%
docente.go:89:   normalizarNombre           100.0%
docente.go:105:  agregarEvaluacionInterna   100.0%
total:                                      100.0% (statements)
```

Captura real de `go test ./docente -v`: ver `captura_go_test.png` adjunto.

## Cobertura alcanzada

**100.0%** de las sentencias del paquete `docente` están cubiertas por las
pruebas (internas + externas combinadas).

## Análisis comparativo: pruebas internas vs. pruebas externas

**Pruebas internas (`docente_test.go`, `package docente`)**

Al estar dentro del mismo paquete, tienen acceso total: pueden instanciar un
`Docente` directamente con `&Docente{...}` (sin pasar por el constructor),
llamar a los métodos privados (`validarEmail`, `normalizarNombre`,
`agregarEvaluacionInterna`) y leer el campo privado `evaluaciones`
directamente. Esto permite probar cada unidad de lógica de forma muy
puntual y aislada, sin depender de que la API pública la exponga "tal cual".
Es el lugar correcto para probar reglas de negocio internas en detalle
(por ejemplo, todos los casos borde de `validarEmail`).

**Pruebas externas (`docente_externo_test.go`, `package docente_test`)**

Al vivir en un paquete distinto (aunque en la misma carpeta, Go permite un
paquete `_test` adicional para pruebas de caja negra), solo pueden usar lo
que el paquete `docente` exporta: `NuevoDocente`, `GetID`, `GetNombre`,
`GetEmail`, `EsEmailValido`, `AgregarEvaluacion`, `GetEvaluaciones`. Cualquier
intento de llamar a `validarEmail()`, `normalizarNombre()` o
`agregarEvaluacionInterna()`, o de leer `d.evaluaciones`, produce un **error
de compilación** (`undefined` / `unexported field or method`), no un error en
tiempo de ejecución. Esto es justamente la prueba de que la encapsulación
funciona: el compilador, no una convención, es quien impide el acceso.

En resumen: las pruebas internas verifican que la lógica interna sea
correcta; las pruebas externas verifican que el "contrato público" del
paquete sea suficiente y que los detalles internos realmente queden
ocultos.

## Reflexión personal

Este ejercicio sirvió para entender en la práctica algo que en teoría suena
simple ("Go encapsula con mayúsculas/minúsculas") pero que recién se nota de
verdad cuando uno intenta romperlo: al escribir `d.validarEmail()` desde
`docente_externo_test.go`, el compilador lo rechaza inmediatamente con un
mensaje claro (`d.validarEmail undefined`). Eso deja mucho más claro que en
otros lenguajes (como Java con `private`) el nivel de visibilidad en Go no
es una palabra clave aparte, sino una convención de nombres a nivel de
paquete, y que el paquete (carpeta) es la unidad real de encapsulación, no
la estructura.

La mayor dificultad fue decidir cómo exponer los métodos privados sin
romper la encapsulación: por ejemplo, `agregarEvaluacionInterna` necesitaba
algún punto de entrada público (`AgregarEvaluacion`) para poder probarlo
desde afuera, pero ese wrapper no debía simplemente "reexportar" la lógica
sin sentido, sino actuar como una puerta controlada. Lo mismo con
`normalizarNombre`, que se invoca automáticamente desde el constructor para
que el dato siempre quede consistente, en vez de obligar al usuario del
paquete a acordarse de llamar a un método de normalización aparte.

Como aprendizaje, quedó claro el valor de separar pruebas de caja blanca
(internas) y caja negra (externas): las internas dan confianza de que la
lógica interna no tiene casos borde rotos, y las externas dan confianza de
que cualquiera que use el paquete desde afuera tendrá una API suficiente y
seguro de que no puede meterse a romper invariantes internos por accidente
(por ejemplo, agregar evaluaciones duplicadas sin pasar por la validación).

**Sobre el uso de IA:** se usó Claude como apoyo para generar la primera
versión del código y de las pruebas, así como para automatizar la
instalación de Go, la ejecución de los comandos (`go build`, `go vet`,
`go test -cover`) y la generación de la captura de la terminal. Sin
embargo, cada método fue revisado y se entiende su funcionamiento completo:
por ejemplo, por qué `normalizarNombre` usa `strings.Fields` (para limpiar
espacios extra automáticamente) en vez de un simple `strings.Split(" ")`, o
por qué `validarEmail` busca el punto específicamente *después* de la
arroba y no en cualquier parte del string (para no aceptar algo como
`juan.perez@gmailcom`, que tiene un punto pero no en el dominio). El reto
de "qué pasa si comento esta línea, ¿compila o no?" se verificó realmente
ejecutando `go vet` con la línea descomentada, en vez de asumir el
resultado.
