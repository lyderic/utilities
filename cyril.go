package main

import (
	"fmt"
	"os"
)

var dic map[rune]rune

func init() {
	dic = make(map[rune]rune)
	dic['a'] = 'а'
	dic['b'] = 'б'
	dic['c'] = 'ц'
	dic['d'] = 'д'
	dic['e'] = 'е'
	dic['f'] = 'ф'
	dic['g'] = 'г'
	dic['h'] = 'ч'
	dic['i'] = 'и'
	dic['j'] = 'ж'
	dic['k'] = 'к'
	dic['l'] = 'л'
	dic['m'] = 'м'
	dic['n'] = 'н'
	dic['o'] = 'о'
	dic['p'] = 'п'
	dic['q'] = 'ю'
	dic['r'] = 'р'
	dic['s'] = 'с'
	dic['t'] = 'т'
	dic['u'] = 'у'
	dic['v'] = 'в'
	dic['w'] = 'ш'
	dic['x'] = 'х'
	dic['y'] = 'ы'
	dic['z'] = 'з'
	dic['é'] = 'э'
	dic['è'] = 'ѐ'
	dic['à'] = 'я'
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cyril <word(s)>")
		fmt.Println("print word(s) in cyrillic")
		return
	}

	for _, lat := range os.Args[1:] {
		var cyr []rune
		for _, l := range lat {
			c := dic[l]
			if c == 0 {
				c = l
			}
			cyr = append(cyr, c)
		}
		fmt.Print(string(cyr), " ")
	}
	fmt.Println()
}
