package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var cpt int = 0
var data string = "0"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func increment(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("stock/cpt-local.txt", os.O_CREATE|os.O_WRONLY, 0600)
	defer file.Close() // on ferme automatiquement le fichier après l'avoir manipulé
	check(err)

	data = readIncrement(file.Name())
	cpt, _ = strconv.Atoi(data)
	cpt = cpt + 1

	fmt.Fprintf(w, "Bonjour, vous avez accédé %d fois à cette page.", cpt)

	writeIncrement(strconv.Itoa(cpt), file)
}

func writeIncrement(c string, file *os.File) {
	if _, err := file.WriteString(c); err != nil {
		panic(err)
	} // écrire dans le fichier
}

func readIncrement(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return string(data)
}

func main() {
	http.HandleFunc("/api", increment)
	log.Fatal(http.ListenAndServe(":80", nil))
}
