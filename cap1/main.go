package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	palavras := os.Args[1:]

	estatisticas := colherEstatisticas(palavras)

	imprimir(estatisticas)
}

func colherEstatisticas(palavras []string) map[string]int {

	estatisticas := make(map[string]int)

	for _, palavra := range palavras {

		inicial := strings.ToUpper(string(palavra[0]))
		fmt.Println(estatisticas)
		contador, encontrado := estatisticas[inicial]

		if encontrado {

			estatisticas[inicial] = contador + 1

		} else {

			estatisticas[inicial] = 1

		}

	}

	return estatisticas

}

func imprimir(estatisticas map[string]int) {

	fmt.Println("Contagem de palavras iniciadas em cada letra:")

	for inicial, contador := range estatisticas {

		fmt.Printf("%s = %d\n", inicial, contador)

	}

}

// var inputVar string
// 	var secretVar string
// 	var file string
// 	var input Input
// 	flag.StringVar(&inputVar, "message", "", "message to create a digest from")
// 	flag.StringVar(&secretVar, "secret", "", "secret for the digest")
// 	flag.StringVar(&file, "file", "", "secret for the digest")
// 	flag.Parse()

// 	input.Secret = secretVar
// 	input.Message = inputVar

// 	if file != "" {
// 		data, _ := ioutil.ReadFile(file)
// 		json.Unmarshal(data, &input)
// 		inputVar = input.Message
// 		secretVar = input.Secret
// 	}

// 	resp, _ := json.Marshal(Generate(input))

// 	fmt.Printf("Digest: %s", resp)

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		json.NewEncoder(w).Encode(resp)
// 		fmt.Fprintf(w, string(resp))
// 	})
// 	http.ListenAndServe(":8080", nil)
