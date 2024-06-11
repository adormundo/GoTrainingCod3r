package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

// Titulo retorna um canal que emite o título da página para cada URL fornecida
func Titulo(url string) <-chan string {
	c := make(chan string)
	go func() {
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
	}()
	return c
}

// oMaisRapido retorna o título da URL que responder primeiro, ou uma mensagem de timeout
func oMaisRapido(url1, url2, url3 string) string {
	c1 := Titulo(url1)
	c2 := Titulo(url2)
	c3 := Titulo(url3)

	// estrutura de controle específica para concorrência (select)
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.cod3r.com.br",
		"https://www.g1.com.br",
		"https://www.youtube.com",
	)
	fmt.Println(campeao)
}
