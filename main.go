package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

/**
Simple ejemplo de Merkle Tree.
Dado un archivo Json con un listado de texto: se devuelve el hash del árbol de Merkle
*/

func RootMerkle(lista []string) string {
	//Importante primero ordenar la lista:
	sort.Strings(lista)

	//Proceseo iteraivo: tomo de dos en dos para armar branches o ramas de hashes
	//si el numero de elementos es impar entonces el ultimo string
	//pasa a la siguiente iteracion sin ser hasheado
	for len(lista) > 1 {
		tmplist := []string{}
		for i := 0; i < len(lista)/2; i++ {
			izq := lista[i*2]
			der := lista[i*2+1]
			h := sha256.New()
			if _, err := h.Write([]byte(izq + der)); err != nil {
				log.Fatalln(err)
			}
			hash := fmt.Sprintf("%x", h.Sum(nil))
			tmplist = append(tmplist, hash)
		}
		if len(lista)%2 != 0 {
			tmplist = append(tmplist, lista[len(lista)-1])
		}
		lista = tmplist
	}
	return lista[0]
}

func main() {
	var filehashes [][]string

	data, err := os.ReadFile("lista.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(data, &filehashes)
	if err != nil {
		log.Fatalln(err)
	}

	//armo un arbol nuevo del listado de hashes del archivo JSON
	arbol := []string{}
	for _, hashes := range filehashes {
		arbol = append(arbol, RootMerkle(hashes))
	}

	fmt.Println("Root de Merkle:", RootMerkle(arbol))

}
