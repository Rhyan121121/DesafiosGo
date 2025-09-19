package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func PauseScreen() {
	fmt.Println("Press Enter to continue...")
	// Cria um novo leitor para a entrada padrão (teclado).
	reader := bufio.NewReader(os.Stdin)
	// ReadString('\n') lê a entrada do usuário até encontrar uma quebra de linha.
	// O valor retornado é ignorado, já que só precisamos que a leitura aconteça.
	reader.ReadString('\n')
}

func ConvertionTemperatureCtoF(t float64) float64 {
	return (t * 9 / 5) + 32
}

func ConvertionTemperatureCtoK(t float64) float64 {
	return t + 273.15
}

func ConvertionTemperatureFtoC(t float64) float64 {
	return (t - 32) * 5 / 9
}

func ConvertionTemperatureFtoK(t float64) float64 {
	return (t-32)*5/9 + 273.15
}

func ConvertionTemperatureKtoC(t float64) float64 {
	return t - 273.15
}

func ConvertionTemperatureKtoF(t float64) float64 {
	return (t-273.15)*9/5 + 32
}

func main() {
	var temperatureWater float64
	var Scalewater int
	for {
		ClearScreen()
		fmt.Println("Enter the scale do you want to use to convert the others (1-Celsius, 2-Fahrenheit, 3-Kelvin, 4- Leave): ")
		fmt.Scan(&Scalewater)
		if Scalewater == 4 {
			fmt.Println("You choose to leave")
			break
		}
		if Scalewater != 1 && Scalewater != 2 && Scalewater != 3 {
			fmt.Println("You choose a invalid Scale of water")
			break
		}

		fmt.Print("Enter the temperature of water: ")
		fmt.Scan(&temperatureWater)

		if Scalewater == 1 {
			resultK := ConvertionTemperatureCtoK(temperatureWater)
			resultF := ConvertionTemperatureCtoF(temperatureWater)
			fmt.Printf("\nThe temperature of water is:ºC%g\nThis temperature in Fahrenheit is:ºF%g\nThis temperature in Kelvin is :ºK%g\n", temperatureWater, resultF, resultK)
			PauseScreen()
		} else if Scalewater == 2 {
			resultK := ConvertionTemperatureFtoK(temperatureWater)
			resultC := ConvertionTemperatureFtoC(temperatureWater)
			fmt.Printf("\nThe temperature of water is:ºF%g\nThis temperature in Celsius is: ºC%g\nThis temperature in Kelvin is :ºK%g\n", temperatureWater, resultC, resultK)
			PauseScreen()
		} else {
			resultF := ConvertionTemperatureKtoF(temperatureWater)
			resultC := ConvertionTemperatureKtoC(temperatureWater)
			fmt.Printf("\nThe temperature of water is:ºK%g\nThis temperature in Celsius is: ºC%g\nThis temperature in Fahrenheit is : ºF%g\n", temperatureWater, resultC, resultF)
			PauseScreen()
		}
	}
}
