package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetSamplesAndDuration(t *testing.T) {

	result, duration := GenerateSamplesAsFloat("test.mp3")
	Convey("Given I try to parse \"test.mp3\"", t, func() {

		Convey("samples equal those predefined", func() {
			So(result, ShouldResemble, []float64{0, 0, 0, 0, 0, 0.1381378173828125, 0.1412353515625, 0.1432647705078125, 0.147125244140625, 0.150604248046875, 0.154022216796875, 0.156005859375, 0.1573944091796875, 0.1590118408203125, 0.1606292724609375, 0.1620635986328125, 0.1624908447265625, 0.163787841796875, 0.16546630859375, 0.1673126220703125, 0.167236328125, 0.168914794921875, 0.16888427734375, 0.1660308837890625, 0.164886474609375, 0.1645965576171875, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		})

		Convey("duration is equal to", func() {
			So(duration, ShouldEqual, 5.36786578e-315)
		})

	})

}

func TestGetStringSamplesAndDuration(t *testing.T) {
	result, duration := GenerateSamplesAsString("test.mp3", 5)
	Convey("Given I try to parse \"test.mp3\"", t, func() {

		Convey("samples equal those predefined", func() {
			So(result, ShouldResemble, []string{"0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.13814", "0.14124", "0.14326", "0.14713", "0.15060", "0.15402", "0.15601", "0.15739", "0.15901", "0.16063", "0.16206", "0.16249", "0.16379", "0.16547", "0.16731", "0.16724", "0.16891", "0.16888", "0.16603", "0.16489", "0.16460", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000", "0.00000"})
		})

		Convey("duration is equal to", func() {
			So(duration, ShouldEqual, 5.36786578e-315)
		})

	})

}
