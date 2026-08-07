# Taller 23 - Goroutines

## Información General

- **Asignatura:** Programación Orientada a Objetos
- **Taller:** 23 - Goroutines
- **Lenguaje:** Go (Golang)

---

# Objetivo

Aplicar el uso de Goroutines, WaitGroup y Mutex para desarrollar un programa concurrente que procese múltiples órdenes de manera simultánea, evitando condiciones de carrera.

---

# Descripción

El programa simula el procesamiento de 20 órdenes utilizando tres goroutines.

Cada orden contiene:

- ID
- Estado (Status)
- Un Mutex para proteger el acceso concurrente.

Durante la ejecución, cada goroutine actualiza el estado de las órdenes y registra el número total de actualizaciones realizadas.

---

# Tecnologías utilizadas

- Go (Golang)
- Goroutines
- sync.WaitGroup
- sync.Mutex
- Visual Studio Code
- Git
- GitHub

---

# Estructura del proyecto

```
taller23-goroutines/
│
├── go.mod
├── main.go
└── README.md
```

---

# Ejecución del proyecto

Abrir una terminal y ejecutar:

```bash
go run main.go
```

---

# Desarrollo por iteraciones

## Iteración 1

### Actividades realizadas

- Creación del proyecto.
- Configuración del módulo Go.
- Definición de la estructura `Order`.
- Declaración de variables globales.


### Resultado

Se creó correctamente la estructura base del proyecto.

---

## Iteración 2

### Actividades realizadas

- Implementación de la función `generateOrders()`.
- Generación de 20 órdenes.
- Uso de `sync.WaitGroup`.
- Creación de tres goroutines para procesar las órdenes.


### Resultado

Las órdenes comenzaron a procesarse de forma concurrente utilizando goroutines.

---

## Iteración 3

### Actividades realizadas

- Implementación de la función `updateOrderStatus()`.
- Protección de cada orden mediante `Mutex`.
- Protección del contador global de actualizaciones.
- Presentación del estado final de las órdenes.



### Resultado

El programa ejecutó correctamente las tres goroutines, actualizando las órdenes sin presentar condiciones de carrera y mostrando el total de actualizaciones realizadas.

Actualizando orden 1 con estado: Procesando
Actualizando orden 1 con estado: Entregado
Actualizando orden 1 con estado: Despachando
Actualizando orden 2 con estado: Procesando
Actualizando orden 3 con estado: Entregado
Actualizando orden 4 con estado: Entregado
Actualizando orden 2 con estado: Procesando
Actualizando orden 2 con estado: Entregado
Actualizando orden 3 con estado: Despachando
Actualizando orden 5 con estado: Despachando
Actualizando orden 4 con estado: Procesando
Actualizando orden 6 con estado: Despachando
Actualizando orden 5 con estado: Despachando
Actualizando orden 7 con estado: Despachando
Actualizando orden 3 con estado: Despachando
Actualizando orden 8 con estado: Procesando
Actualizando orden 6 con estado: Despachando
Actualizando orden 4 con estado: Entregado
Actualizando orden 7 con estado: Entregado
Actualizando orden 9 con estado: Procesando
Actualizando orden 10 con estado: Despachando
Actualizando orden 11 con estado: Despachando
Actualizando orden 5 con estado: Procesando
Actualizando orden 6 con estado: Despachando
Actualizando orden 12 con estado: Despachando
Actualizando orden 8 con estado: Entregado
Actualizando orden 9 con estado: Procesando
Actualizando orden 7 con estado: Entregado
Actualizando orden 10 con estado: Despachando
Actualizando orden 13 con estado: Procesando
Actualizando orden 8 con estado: Despachando
Actualizando orden 11 con estado: Entregado
Actualizando orden 9 con estado: Entregado
Actualizando orden 14 con estado: Entregado
Actualizando orden 15 con estado: Procesando
Actualizando orden 10 con estado: Procesando
Actualizando orden 12 con estado: Entregado
Actualizando orden 11 con estado: Despachando
Actualizando orden 16 con estado: Despachando
Actualizando orden 12 con estado: Entregado
Actualizando orden 17 con estado: Entregado
Actualizando orden 18 con estado: Entregado
Actualizando orden 13 con estado: Despachando
Actualizando orden 19 con estado: Entregado
Actualizando orden 14 con estado: Despachando
Actualizando orden 20 con estado: Entregado
Actualizando orden 13 con estado: Procesando
Actualizando orden 15 con estado: Procesando
Actualizando orden 16 con estado: Entregado
Actualizando orden 14 con estado: Entregado
Actualizando orden 15 con estado: Despachando
Actualizando orden 17 con estado: Despachando
Actualizando orden 16 con estado: Procesando
Actualizando orden 17 con estado: Entregado
Actualizando orden 18 con estado: Entregado
Actualizando orden 18 con estado: Despachando
Actualizando orden 19 con estado: Procesando
Actualizando orden 19 con estado: Entregado
Actualizando orden 20 con estado: Procesando
Actualizando orden 20 con estado: Entregado

Estado final de las órdenes:
Orden 1 -> Despachando
Orden 2 -> Entregado
Orden 3 -> Despachando
Orden 4 -> Entregado
Orden 5 -> Procesando
Orden 6 -> Despachando
Orden 7 -> Entregado
Orden 8 -> Despachando
Orden 9 -> Entregado
Orden 10 -> Procesando
Orden 11 -> Despachando
Orden 12 -> Entregado
Orden 13 -> Procesando
Orden 14 -> Entregado
Orden 15 -> Despachando
Orden 16 -> Procesando
Orden 17 -> Entregado
Orden 18 -> Despachando
Orden 19 -> Entregado
Orden 20 -> Entregado

Todas las operaciones completadas.
Total Actualizaciones: 60
---

# Conceptos aplicados

## Goroutines

Permiten ejecutar funciones de manera concurrente.

## WaitGroup

Espera a que todas las goroutines finalicen antes de terminar el programa.

## Mutex

Protege los recursos compartidos para evitar conflictos cuando varias goroutines intentan acceder al mismo dato.

---

# Conclusiones

- Se aplicó el uso de goroutines para ejecutar tareas concurrentes.
- Se utilizó WaitGroup para sincronizar la finalización de las goroutines.
- Se implementó Mutex para proteger los datos compartidos.
- El taller permitió comprender la importancia de la sincronización en aplicaciones concurrentes desarrolladas en Go.