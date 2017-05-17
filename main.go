package main

import (
	"net/http"
	"fmt"
	"encoding/base64"
	"encoding/json"
)

func base_64_resize(cadena string) string {
	data,err := base64.StdEncoding.DecodeString(cadena)
	if err != nil {
		fmt.Println("Hay un error con data")
	}
	var longitud  = len(data)
	if longitud%2!=0{
		longitud--
	}

	var arreglo = make([]byte,longitud/2)


	for con := 0;con < longitud/2 ; con++  {
		arreglo[con] = byte((data[con*2]+data[(con*2)+1])/2)
	}

	nuevo := base64.StdEncoding.EncodeToString(arreglo)
	return nuevo
}

func gray_scale(cadena string)string  {

	return "..."
}

type bitmap struct {
	Base64 string	`json:"base_64"`
	Height int	`json:"height"`
	Wight int	`json:"wight"`
}

func main()  {
	bit := bitmap{"Qk1mAAAAAAAAADYAAAAoAAAABAAAAAQAAAABABgAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAA3HETjpzjlav013gn3JZN0uPa0OTD2Kdmn5ZOqtazt9iptZVH63kCsJVPuZBL4nkJ",10,10}

	http.HandleFunc("/ejercicio1", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"ejercicio1.html")
	})
	http.HandleFunc("/ejercicio2", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"ejercicio2.html")
	})
	http.HandleFunc("/ejercicio3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"<h1>Estoy en el ejercicio 3</h1>")
	})
	http.HandleFunc("/ejercicio4", func(w http.ResponseWriter, r *http.Request) {
		bitNuevo := bitmap{base_64_resize(bit.Base64),bit.Height/2,bit.Wight/2}
		json.NewEncoder(w).Encode(bit)
		json.NewEncoder(w).Encode(bitNuevo)
	})

	http.ListenAndServe(":8080",nil)
}