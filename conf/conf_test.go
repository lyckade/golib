package conf

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConf(t *testing.T) {

	Convey("It should load a confFile into a conf struct", t, func() {
		type conf struct {
			Folder string `json:"folder"`
		}
		var c conf
		LoadConf("test.json", &c)
		So(c.Folder, ShouldEqual, "path/to/folder")
	})

}
