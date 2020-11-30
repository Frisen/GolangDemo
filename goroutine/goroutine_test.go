package goroutine_test

import (
	"GolangDemo/goroutine"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRunWithMaxProd(t *testing.T) {
	Convey("goroutine", t, func() {
		Convey("测", func() {
			//goroutine.RunWithMaxProd()
			goroutine.MutiSenderSingleReceipt()
		})
	})
}
