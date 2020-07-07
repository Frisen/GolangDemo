package ip

import "net/http"

//Entity 解析请求参数借口
type Entity interface {
	UnmarshalerHTTP(*http.Request) error
}

//GetEntity 获取解析后的实体
func GetEntity(r *http.Request, v Entity) error {
	if err := v.UnmarshalerHTTP(r); err != nil {
		return err
	}
	return nil
}
