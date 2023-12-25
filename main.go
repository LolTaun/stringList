package main

import (
	"fmt"
	"stringList/stringlist"
)

func main() {
	var err error
	sl := stringlist.New("Hello, World!")
	sl.Append('!')
	sl.Prepend('H')
	err = sl.Insert('e', 1)
	if err != nil {
		fmt.Println(err, "sl.Insert('e', 1)")
	}
	err = sl.Insert('l', 2)
	if err != nil {
		fmt.Println(err, "sl.Insert('l', 2)")
	}
	err = sl.Insert('l', 3)
	if err != nil {
		fmt.Println(err, "sl.Insert('l', 3)")
	}
	err = sl.Insert('o', 4)
	if err != nil {
		fmt.Println(err, "sl.Insert('o', 4)")
	}
	err = sl.Insert(',', 5)
	if err != nil {
		fmt.Println(err, "sl.Insert('o', 5)")
	}
	err = sl.Insert(' ', 6)
	if err != nil {
		fmt.Println(err, "sl.Insert('o', 6)")
	}
	fmt.Println(sl.String(), "length:", sl.Length)
	err = sl.Remove(20)
	if err != nil {
		fmt.Println(err, "sl.Remove(20)")
	}
	fmt.Println(sl.String(), "length:", sl.Length)
	
	c, err := sl.At(4)
	if err != nil {
		fmt.Println(err, "sl.At(4)")
	}
	fmt.Println("sl.At(4) =", string(c))

	sl = stringlist.New("Hello, World!")

	sl2 := stringlist.New("Hello, ")
	fmt.Println(sl2.String(), "length:", sl2.Length)
	sl3 := stringlist.New("World!")
	fmt.Println(sl3.String(), "length:", sl3.Length)

	sl4 := sl2.Concat(sl3)
	fmt.Println(sl4.String(), "length:", sl4.Length)

	if sl4.Equals(sl) {
		fmt.Println("sl4.Equals(sl) = true")
	} else {
		fmt.Println("sl4.Equals(sl) = false")
	}

	fmt.Println(sl4.IndexOf('W'), "sl4.IndexOf('W')")

	sl5, err := sl.Substring(2, 5)
	if err != nil {
		fmt.Println(err, "sl.Substring(2, 5)")
	}
	fmt.Println(sl5.String(), "length:", sl5.Length)

	sl = sl.Replace('o', '0', 2)
	fmt.Println(sl.String(), "length:", sl.Length)
}
