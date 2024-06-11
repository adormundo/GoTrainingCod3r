package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Título retorna um canal somente leitura que emite títulos de páginas web a partir das URLs fornecidas
func Título(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			defer close(c) // Fechar o canal ao terminar todas as goroutines
			resp, err := http.Get(url)
			if err != nil {
				c <- fmt.Sprintf("Erro ao acessar a URL: %v", err)
				return
			}
			defer resp.Body.Close()

			html, err := io.ReadAll(resp.Body)
			if err != nil {
				c <- fmt.Sprintf("Erro ao ler o corpo da resposta: %v", err)
				return
			}

			r, _ := regexp.Compile("<title>(.*?)</title>")
			title := r.FindStringSubmatch(string(html))
			if len(title) != 0 {
				c <- title[1]
			} else {
				c <- fmt.Sprintf("Título não encontrado para a URL: %v", url)
			}
		}(url)
	}
	return c
}

// encaminhar lê de origem e escreve para destino
func encaminhar(origem <-chan string, destino chan<- string) {
	for msg := range origem {
		destino <- msg
	}
}

// juntar multiplexa (combina) mensagens de múltiplos canais em um único canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		encaminhar(entrada1, c)
		encaminhar(entrada2, c)
	}()
	return c
}

func main() {
	t1 := Título("https://www.cod3r.com.br", "https://www.google.com")
	t2 := Título("https://www.g1.com.br", "https://www.youtube.com")

	c := juntar(t1, t2)
	for i := 0; i < 4; i++ { // Iterar para receber todas as mensagens
		fmt.Println(<-c)
	}
}
