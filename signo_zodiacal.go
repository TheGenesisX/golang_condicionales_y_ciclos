package main

import "fmt"

func main() {
	var dia uint64
	var mes uint64

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	switch mes {
	case 1:
		if dia >= 1 && dia <= 20 {
			fmt.Print("Capricornio")
		} else if dia >= 21 && dia <= 31 {
			fmt.Print("Acuario")
		}
	case 2:
		if dia >= 1 && dia <= 18 {
			fmt.Print("Acuario")
		} else if dia >= 19 && dia <= 29 {
			fmt.Print("Piscis")
		}
	case 3:
		if dia >= 1 && dia <= 20 {
			fmt.Print("Piscis")
		} else if dia >= 21 && dia <= 31 {
			fmt.Print("Aries")
		}
	case 4:
		if dia >= 1 && dia <= 20 {
			fmt.Print("Aries")
		} else if dia >= 21 && dia <= 30 {
			fmt.Print("Tauro")
		}
	case 5:
		if dia >= 1 && dia <= 20 {
			fmt.Print("Tauro")
		} else if dia >= 21 && dia <= 31 {
			fmt.Print("Geminis")
		}
	case 6:
		if dia >= 1 && dia <= 21 {
			fmt.Print("Geminis")
		} else if dia >= 22 && dia <= 30 {
			fmt.Print("Cancer")
		}
	case 7:
		if dia >= 1 && dia <= 22 {
			fmt.Print("Cancer")
		} else if dia >= 23 && dia <= 31 {
			fmt.Print("Leo")
		}
	case 8:
		if dia >= 1 && dia <= 22 {
			fmt.Print("Leo")
		} else if dia >= 23 && dia <= 31 {
			fmt.Print("Virgo")
		}
	case 9:
		if dia >= 1 && dia <= 22 {
			fmt.Print("Virgo")
		} else if dia >= 23 && dia <= 30 {
			fmt.Print("Libra")
		}
	case 10:
		if dia >= 1 && dia <= 22 {
			fmt.Print("Libra")
		} else if dia >= 23 && dia <= 31 {
			fmt.Print("Escorpio")
		}
	case 11:
		if dia >= 1 && dia <= 22 {
			fmt.Print("Escorpio")
		} else if dia >= 23 && dia <= 30 {
			fmt.Print("Sagitario")
		}
	case 12:
		if dia >= 1 && dia <= 21 {
			fmt.Print("Sagitario")
		} else if dia >= 22 && dia <= 31 {
			fmt.Print("Capricornio")
		}
	}
}
