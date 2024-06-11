package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sync"
)

// Título retorna um canal somente leitura que emite títulos de páginas web a partir das URLs fornecidas
func Título(urls ...string) <-chan string {
	c := make(chan string)
	var wg sync.WaitGroup // WaitGroup para sincronização das goroutines
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				c <- fmt.Sprintf("Erro ao acessar a URL: %v", url)
				return
			}
			defer resp.Body.Close()

			html, err := io.ReadAll(resp.Body)
			if err != nil {
				c <- fmt.Sprintf("Erro ao ler o corpo da resposta: %v", url)
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

	go func() {
		wg.Wait() // Aguarda todas as goroutines terminarem
		close(c)  // Fecha o canal após todas as goroutines terem completado
	}()

	return c
}

func main() {

	t1 := Título("https://www.cod3r.com.br", "https://www.google.com")
	t2 := Título("https://www.amazon.com", "https://www.youtube.com")

	printTítulos("Primeiros", t1, t2)
	printTítulos("Segundos", t1, t2)

}

func printTítulos(msg string, canais ...<-chan string) {
	fmt.Println(msg + ":")
	for _, ch := range canais {
		for titulo := range ch {
			fmt.Printf("  - %s\n", titulo)
		}
	}
}
