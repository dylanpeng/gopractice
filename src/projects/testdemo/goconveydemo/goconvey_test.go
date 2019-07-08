package goconveydemo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

func TestIsEqual(t *testing.T){
	Convey("1 == 1", t, func() {
		So(IsEqual(1, 1), ShouldBeTrue)
	})
}

func TestIsEqualWithErr(t *testing.T){
	Convey("IsEqualWithErr", t, func(){
		Convey("2 > 1, over", func(){
			ok, err := IsEqualWithErr(2, 1)
			So(ok, ShouldBeFalse)
			So(err, ShouldNotBeNil)
		})

		Convey("1 < 2, under", func(){
			ok, err := IsEqualWithErr(1, 2)
			So(ok, ShouldBeFalse)
			So(err, ShouldNotBeNil)
		})

		Convey("1 = 1, equal", func(){
			ok, err := IsEqualWithErr(1, 1)
			So(ok, ShouldBeTrue)
			So(err, ShouldBeNil)
		})
	})
}
