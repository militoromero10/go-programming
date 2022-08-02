# Common Commands

 - go version
 - go help
 - go run
 - godoc paquete funcion


# Data types
```
	Declaration: 
		var name type = value -> attribute
		name := value -> methods
	
	Constants Enumeration, iota is usefull working with enumerations
	const {
		First = iota
		Second
		Third
	}

	int, uint, uintptr

	uint8, uint16, uint32, uint64
	int8, int16, int32, int64

	byte = uint8
	rune = int32

	float32, float64
	complex64, complex128
    
```

# Slices
	- range, itera sobre un array devuelve index, element
	- len(), tamaño del array
	- make(), crea slices
	- append(), crea un slice apartir de otro slice y parametros
	- copy(), copia un slice en otro
	- [low:high], devuelve parte de un slice
	- delete, elimina una llave-valor en un map

# Special functions
	- closure, recursivo, retornar multiples valores
	- defer, ejecuta una funcion, justo despues de otra
	- panic
	- recover

# Structs and interfaces
	- type name struct
	- type name interface
	- type name dataType

# Concurrency
	- go funcion
	- make/chan, emite o recibe informacion con operador <-, usualmente los channels son sincronos
	- select/case/default, el switch de la concurrecia
	- buffer se define igual que un channel, con parametro adicional de capacidad, para procesos asincronos

# reread packages and testing

# Core Packages - chapter13
	- strings
	- io
	- os
	- container/list
	- sort
	- non-cryptographic
		+ hash/adler32
		+ hash/crc32
		+ hash/crc64
		+ hash/fnv
	- crypthographic
		+ crypto/sha1
	- server
	- encoding
	- net/rpc
	- sync
	- sync/atomic

# Next  steps
	- http://golang.org/src/pkg/
	- http://golang.org/src/pkg/io/ioutil/ioutil.go
	- Projec Euler, http://projecteuler.net 
	- http://groups.google.com/group/golang-nuts