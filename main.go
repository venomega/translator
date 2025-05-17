package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	//"os"
)

var model string = "gemma2:2b"
var text string = "la vida es bella"
var lang string = "fr"
var host string = "127.0.0.1:11434"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type Response struct {
	Model  string   `json:"model"`
	CreateAt string `json:"create_at"`
	Message Message `json:"message"`
}

func main(){
	//looking for OLLAMA_HOST env var
	value, exist := os.LookupEnv("OLLAMA_HOST")
	if exist {
		host = value
	}

	//Handling args
	if len(os.Args) == 3 {
		if os.Args[1] == "-" {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
			text = string(data)
		}else{
			text = os.Args[1]
		}
		lang = os.Args[2]
		goto start
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "-" {
			// Read from stdin
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
			text = string(data)
		}else{
			// Read from args
			text = os.Args[1]
		}
	}else{
		//Error
		fmt.Println("Usage: translator <text> <lang>")
	}

	start:

	req, err := http.NewRequest("POST", "http://" + host + "/api/chat", nil)
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
		"content": "Traduce `" + text + "` al " + lang,
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

	if resp.StatusCode == 500{
		println("Internal Server Error")
		os.Exit(1)
	}

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var resp_json Response

	data := bytes.Split(body, []byte("}\n"))

	for _, i := range data {
		if len(i) != 0 {
			if i[len(i) - 1] != byte('}') {
				i = append(i, byte('}'))
			}
		}else{
			break
		}

		err = json.Unmarshal(i, &resp_json)
		if err != nil {
			break
		}

		//print result
		for _, x := range strings.Split(string(resp_json.Message.Content), "\\n"){
			fmt.Printf("%s", x)
		}


	}

}

