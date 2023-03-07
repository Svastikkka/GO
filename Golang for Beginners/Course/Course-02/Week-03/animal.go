package main
import (
	"bufio"
	"os"
	"fmt"
	"strings"
)
func main() {
	type Animal struct {
		food string
		locomotion string
		noise string
	}
	Animal{food: "grass", locomotion: "walk", noise: "moo"}
	Animal{food: "worms", locomotion: "fly", noise: "peep"}
	Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
	input_split := strings.Split(input.Text(), " ")
	fmt.Print(input_split)

}