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
			fmt.Print("capricornio")
		} else if dia >= 21 && dia <= 31 {
			fmt.Print("acuario")
		}
	case 2:
		if dia >= 1 && dia <= 18 {
			fmt.Print("acuario")
		} else if dia >= 19 && dia <= 29 {
			fmt.Print("piscis")
		}
	case 3:
		if dia >= 1 && dia <= 20 {
			fmt.Print("piscis")
		} else if dia >= 21 && dia <= 31 {
			fmt.Print("aries")
		}
	case 4:
		if dia >= 1 && dia <= 20 {
			fmt.Print("aries")
		} else if dia >= 21 && dia <= 30 {
			fmt.Print("tauro")
		}
	case 5:
		if dia >= 1 && dia <= 20 {
			fmt.Print("tauro")
		} else if dia >= 21 && dia <= 31 {
			fmt.Print("geminis")
		}
	case 6:
		if dia >= 1 && dia <= 21 {
			fmt.Print("geminis")
		} else if dia >= 22 && dia <= 30 {
			fmt.Print("cancer")
		}
	case 7:
		if dia >= 1 && dia <= 22 {
			fmt.Print("cancer")
		} else if dia >= 23 && dia <= 31 {
			fmt.Print("leo")
		}
	case 8:
		if dia >= 1 && dia <= 22 {
			fmt.Print("leo")
		} else if dia >= 23 && dia <= 31 {
			fmt.Print("virgo")
		}
	case 9:
		if dia >= 1 && dia <= 22 {
			fmt.Print("virgo")
		} else if dia >= 23 && dia <= 30 {
			fmt.Print("libra")
		}
	case 10:
		if dia >= 1 && dia <= 22 {
			fmt.Print("libra")
		} else if dia >= 23 && dia <= 31 {
			fmt.Print("escorpio")
		}
	case 11:
		if dia >= 1 && dia <= 22 {
			fmt.Print("escorpio")
		} else if dia >= 23 && dia <= 30 {
			fmt.Print("sagitario")
		}
	case 12:
		if dia >= 1 && dia <= 21 {
			fmt.Print("sagitario")
		} else if dia >= 22 && dia <= 31 {
			fmt.Print("capricornio")
		}
	}
}
