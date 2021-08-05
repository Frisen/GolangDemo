package datastruct_test

import (
	datastruct "GolangDemo/data_struct"
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_CharMatch(t *testing.T) {
	fmt.Print("Enter text: \n")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('i')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	fmt.Println(input)

	charMap := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	count := len([]rune(input))
	stack := datastruct.NewStack()
	for i := 0; i < count; i++ {
		if stack.Length() > 0 {
			if v, ok := charMap[stack.Peek().(byte)]; ok && v == input[i] {
				stack.Pop()
				continue
			}
		}
		stack.Push(input[i])
	}
	if stack.Length() > 0 {
		fmt.Println("not matched!")
	} else {
		fmt.Println("matched!")
	}
}
