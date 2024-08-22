package rangenum

import "fmt"

func DoMagic() {
	for range 10 {
		fmt.Println("abrakadabra")
	}
}

type Magician interface {
	DoMagic()
}
