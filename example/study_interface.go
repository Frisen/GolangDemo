package example

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

//RubyDate python时间格式
type RubyDate time.Time

//JSONDate 解析字符串
func JSONDate() {
	input := `{"created_at": "Thu May 31 00:00:01 +0000 2012"}`
	var dateMap map[string]RubyDate
	if err := json.Unmarshal([]byte(input), &dateMap); err != nil {
		panic(err)
	}
	v := dateMap["created_at"]
	t := time.Time(v)
	fmt.Println(t)
	fmt.Println(reflect.TypeOf(t))
}

//UnmarshalJSON 重写方法，可以转换Ruby格式的日期
func (t *RubyDate) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = RubyDate(v)
	return nil
}
