package main

import (
	"fmt"

	"github.com/billzajac/rltao"
)

func main() {
	id := rltao.GenerateId()
	tetragram := rltao.GetTetragram(id)
	passage := rltao.GetPassage(id)
	fmt.Printf("%02d\n%s\n%s\n%s\n", id, tetragram, passage.Title, passage.Body)
}
