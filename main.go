package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	//"os"
)

var model string = "gemma2:2b"

func main(){
	//data, err := os.ReadFile("/dev/stdin")
	//if err != nil {
	//	panic(err)
	//}

	req, err := http.NewRequest("POST", "http://127.0.0.1:11434/api/chat", nil)
	// Handle error
	if err != nil {
		panic(err)
	}
	// Init Request
	req.Header.Set("User-Agent", "ollama/1.0")
	req.Header.Set("Accept", "application/x-ndjson")
	req.Header.Set("Content-Type", "application/json")
	req_json := make(map[string]interface{})
	req_json["model"] = model
	//req_json["system"] = "Tu proposito es recibir una orden de traduccion de texto, y devolver el resultado en el lenguaje que se requiera sin adornos en la respuesta. No hagas introduccion, ni concluciones, ni notas, ni notas importantes, ni explicaciones, solo se consiso con lo que se te pregunta, da explicaciones cuando el usuario las pida, sino, no las des.\n\n Ejemplo de lo que NO puedes hacer: \"Le monde\" se traduce al español como **\"El mundo\"**\n EJEMPLO DE COMO PUEDE SER: \"El mundo\""
	req_json["options"] = map[string]interface{}{}
	req_json["messages"] = make([]map[string]interface{}, 0)
	req_json["messages"] = append(req_json["messages"].([]map[string]interface{}), map[string]interface{}{
		"role":    "user",
		"content": "No hagas introduccion, ni concluciones, ni notas, ni notas importantes, ni explicaciones, solo se consiso con lo que se te pregunta, da explicaciones cuando el usuario las pida, sino, no las des",
	},
)
	req_json["messages"] = append(req_json["messages"].([]map[string]interface{}), map[string]interface{}{
		"role": "assistant",
		"content": "Entendido",
	})
	req_json["messages"] = append(req_json["messages"].([]map[string]interface{}), map[string]interface{}{
		"role": "user",
		"content": "Traduce `la vida es bella` al fr",
	})


	jsonData, err := json.Marshal(req_json)
	if err != nil {
		panic(err)
	}
	req.Body = io.NopCloser(bytes.NewReader(jsonData))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Imprimir la respuesta
	println(string(body))
}

