package main
import (
	"os"
	"bufio"
	"fmt"
	"encoding/json"
)

func main(){
	name := bufio.NewScanner(os.Stdin)
	name.Scan()
	address := bufio.NewScanner(os.Stdin)
	address.Scan()

	m := make(map[string]string)
	m["name"] = name.Text()
	m["address"] = address.Text()

	jsonData, err := json.Marshal(m)
	if err != nil{ fmt.Println("Something bad happen")}
	fmt.Printf("%s",jsonData)
}