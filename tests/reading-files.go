package main
import (
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("./example.txt")
	fmt.Println(dat)
	fmt.Println(string(dat))
}