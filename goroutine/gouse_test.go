package goroutine_test

import (
	"GolangDemo/goroutine"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCakeMake(t *testing.T) {
	Convey("goroutine", t, func() {
		Convey("单go方法", func() {
			fmt.Print("\n")
			goroutine.Try()
		})
	})
}
