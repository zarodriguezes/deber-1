# DocenteEvaluacion — Encapsulación en Go

Proyecto de evaluación de Programación en Go

Este proyecto fue desarrollado para la asignatura de Programación de la UIDE con el propósito de demostrar el uso del principio de encapsulación en Go mediante la implementación de atributos y métodos privados, además de pruebas unitarias internas y externas.

Autor: Zack Rodriguez

Organización del proyecto
DocenteEvaluacion/
├── docente/
│   ├── docente.go              # Definición de la estructura Docente y sus métodos
│   ├── docente_test.go         # Pruebas internas (package docente)
│   └── docente_externo_test.go # Pruebas externas (package docente_test)
├── go.mod
└── README.md
Métodos privados implementados
Método	Función
validarEmail()	Comprueba que el correo electrónico contenga una arroba (@) y un punto (.) ubicado después de ella.
normalizarNombre()	Convierte el nombre al formato título, colocando en mayúscula la primera letra de cada palabra.
agregarEvaluacionInterna(idEvaluacion string)	Incorpora un identificador de evaluación al slice privado evaluaciones, evitando valores vacíos o repetidos.

Estos métodos son de acceso privado, ya que sus nombres comienzan con letra minúscula, por lo que únicamente pueden utilizarse dentro del paquete docente. Para ofrecer funcionalidades equivalentes desde otros paquetes se implementaron los métodos públicos EsEmailValido(), GetEmail(), AgregarEvaluacion() y GetEvaluaciones().

Asimismo, el constructor NuevoDocente() invoca internamente a normalizarNombre(), garantizando que todos los nombres queden correctamente formateados desde el momento en que se crea un objeto.

Comandos ejecutados y resultados
# Compilar el módulo completo
$ go build ./...
# (sin salida = compilación exitosa)

# Ejecutar análisis estático
$ go vet ./...
# (sin advertencias)

# Ejecutar pruebas unitarias
$ go test ./docente -v
PASS
ok      docenteevaluacion/docente      0.001s

# Ejecutar pruebas con cobertura
$ go test ./docente -v -cover
PASS
coverage: 100.0% of statements
ok      docenteevaluacion/docente      0.004s

# Obtener cobertura por función
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

La evidencia de la ejecución del comando go test ./docente -v se encuentra en la imagen captura_go_test.png.

Cobertura obtenida

Las pruebas desarrolladas alcanzaron una cobertura del 100 % de las sentencias del paquete docente, considerando tanto las pruebas internas como las externas.

Comparación entre pruebas internas y externas
Pruebas internas (docente_test.go)

Al pertenecer al mismo paquete (package docente), estas pruebas tienen acceso completo a los elementos internos del código. Esto permite crear instancias de Docente directamente, invocar métodos privados como validarEmail(), normalizarNombre() y agregarEvaluacionInterna(), además de acceder al atributo privado evaluaciones.

Gracias a ello es posible validar de manera detallada la lógica interna del paquete y comprobar casos específicos que no forman parte de la interfaz pública.

Pruebas externas (docente_externo_test.go)

Estas pruebas utilizan el paquete docente_test, por lo que solo pueden interactuar con los elementos exportados por el paquete docente. Entre ellos se encuentran NuevoDocente(), GetID(), GetNombre(), GetEmail(), EsEmailValido(), AgregarEvaluacion() y GetEvaluaciones().

Si se intenta acceder a métodos privados o atributos no exportados, como validarEmail(), normalizarNombre(), agregarEvaluacionInterna() o evaluaciones, el compilador genera un error al momento de compilar el programa. Esto confirma que la encapsulación funciona correctamente y que los detalles internos permanecen protegidos.

En conclusión, las pruebas internas permiten verificar el funcionamiento de la implementación, mientras que las pruebas externas validan el comportamiento de la API pública desde la perspectiva de un usuario del paquete.

Reflexión personal

La realización de este proyecto permitió comprender de manera práctica cómo funciona la encapsulación en Go. Aunque inicialmente parece un concepto sencillo, la diferencia entre elementos exportados y no exportados se aprecia claramente cuando se intenta acceder a un método privado desde otro paquete y el compilador impide dicha operación.

A diferencia de otros lenguajes que utilizan palabras reservadas como private, Go basa el control de acceso en el uso de mayúsculas y minúsculas, siendo el paquete la verdadera unidad de encapsulación. Esta característica hace que el código sea sencillo y consistente.

Uno de los principales desafíos consistió en definir la forma adecuada de exponer determinadas funcionalidades sin romper el encapsulamiento. Por ejemplo, agregarEvaluacionInterna() debía mantenerse privado, pero era necesario ofrecer un método público (AgregarEvaluacion()) que actuara como una interfaz controlada para realizar la operación respetando las validaciones correspondientes. De manera similar, normalizarNombre() se ejecuta automáticamente desde el constructor para asegurar que toda instancia mantenga datos consistentes desde su creación.

Otra enseñanza importante fue comprender la diferencia entre las pruebas de caja blanca y las pruebas de caja negra. Las primeras permiten validar en profundidad la lógica interna del paquete, mientras que las segundas comprueban que la interfaz pública sea suficiente para cualquier usuario y que los elementos internos permanezcan protegidos.

Uso de inteligencia artificial

Como apoyo durante el desarrollo se utilizó Claude para generar una versión inicial del código y de las pruebas, además de facilitar la instalación de Go, la ejecución de herramientas como go build, go vet y go test, y la obtención de la captura de la terminal.

No obstante, todo el código fue revisado y comprendido antes de su entrega. Se verificó el funcionamiento de cada método y se analizaron las decisiones de implementación, como el uso de strings.Fields() para eliminar espacios innecesarios al normalizar nombres, o la validación de que el punto del correo electrónico aparezca después de la arroba para evitar aceptar direcciones inválidas. Asimismo, las pruebas se ejecutaron realmente para confirmar los resultados obtenidos, en lugar de asumir su comportamiento.
