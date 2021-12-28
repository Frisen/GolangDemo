package stringstudy_test

import (
	"GolangDemo/stringstudy"
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	// i := stringstudy.LengthOfLongestSubstring("abcbde")
	i := stringstudy.NewStrLength("abcbde")

	fmt.Println(i)
}
