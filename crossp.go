package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: crossp <A> <B> <C>")
		fmt.Println("Compute a 'produit en croix'")
		fmt.Println("CROSS PRODUCT")
		fmt.Println(" A -> B")
		fmt.Println(" C -> X")
		return
	}
	a := extract(os.Args[1])
	b := extract(os.Args[2])
	c := extract(os.Args[3])

	x := (b * c) / a
	fmt.Println(x)

}

func extract(input string) float64 {
	f, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
