package main

import "fmt"

// printMessage imprime uma mensagem fixa "F1"
func printMessage() {
	fmt.Println("F1")
}

// printFormattedMessages imprime duas strings formatadas
func printFormattedMessages(message1, message2 string) {
	fmt.Printf("F2 %s %s\n", message1, message2)
}

// getMessage retorna uma string fixa "F3"
func getMessage() string {
	return "F3"
}

// formatMessages retorna uma string formatada a partir de duas strings de entrada
func formatMessages(message1, message2 string) string {
	return fmt.Sprintf("F4: %s %s", message1, message2)
}

// getTwoMessages retorna duas strings fixas
func getTwoMessages() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	printMessage()
	printFormattedMessages("Deus", "Fiel")

	message3 := getMessage()
	formattedMessage := formatMessages("Deus", "Fiel")

	return1, return2 := getTwoMessages()

	fmt.Println(message3, formattedMessage, return1, return2)
}
