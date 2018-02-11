package httputils

import (
	"testing"

	"net/http"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetHttpPool(t *testing.T) {
	Convey("get a http pool", t, func() {
		hpool := GetHttpPool(100, 5000)
		So(hpool, ShouldNotBeNil)
	})
}

func TestHttpConPool_Request(t *testing.T) {
	hpool := GetHttpPool(100, 100)
	_, err := hpool.Request("http://www.baidu.com", http.MethodGet, "", map[string]string{})
	Convey("test http get request", t, func() {
		So(err, ShouldBeNil)
	})
}
