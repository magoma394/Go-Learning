package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main() {
	fmt.Println("Enter a string:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))
	if strings.HasPrefix(input, "i") &&
    strings.HasSuffix(input, "n") &&
    strings.Contains(input, "a") {
		fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}		
}