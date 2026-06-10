package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"the-go-programming-language-practice/ch2/packages-and-files/converter/lengthconv"
	"the-go-programming-language-practice/ch2/packages-and-files/converter/tempconv"
	"the-go-programming-language-practice/ch2/packages-and-files/converter/weightconv"
)

func main() {
	var args = os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Режим ввода через консоль\n")
		stdInMod()
		return
	}

	for _, arg := range args {
		switch arg {
		case "--help", "-h":
			printUsage()
			return
		default:
			val, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: '%s' не число.\n", arg)
				continue
			}
			runConversion(val)

		}
	}

}

func printUsage() {
	fmt.Println("Утилита для конвертации величин\n")
	fmt.Println("Справка по использованию\n")

	fmt.Println("Если аргументы отсутствуют, то программа работает в режиме ввода через консоль.")

}

func runConversion(v float64) {
	fmt.Printf("\n--- Перевод для значения: %g ---\n\n", v)

	fmt.Println("Температура:")
	c := tempconv.Celsius(v)
	f := tempconv.Fahrenheit(v)

	fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
	fmt.Printf("%s = %s\n\n", f, tempconv.FToC(f))

	fmt.Println("Длина:")
	m := lengthconv.Meter(v)
	ft := lengthconv.Foot(v)

	fmt.Printf("%s = %s\n", m, lengthconv.MToFt(m))
	fmt.Printf("%s = %s\n\n", ft, lengthconv.FtToM(ft))

	fmt.Println("Вес:")
	p := weightconv.Pound(v)
	kg := weightconv.Kilogram(v)

	fmt.Printf("%s = %s\n", p, weightconv.PtoKg(p))
	fmt.Printf("%s = %s\n", kg, weightconv.KgToP(kg))
}

func stdInMod() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите числа для конвертации, для выхода введите 'exit'")
	for scanner.Scan() {
		text := scanner.Text()

		val, err := strconv.ParseFloat(text, 64)

		if err != nil {
			if text == "exit" {
				os.Exit(0)
			}

			fmt.Fprintf(os.Stderr, "Error: '%s' не число.\n", text)
			continue
		}

		runConversion(val)
	}
}
