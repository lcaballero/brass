package bank


import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	. "brass/fortesting"
)

const target = "src/brass/.test_files/target"

func SomeRoll() *Roll {
	bank := NewBank("items-2")
	roll,_ := bank.Register(RollDef{
		Name: "roll-3",
		Type: "unknown-type",
	})
	return roll
}

func isEven(e interface{}) bool {
	n,_ := e.(int)
	return n % 2 == 0
}

func isNegative(e interface{}) bool {
	n,_ := e.(int)
	return n < 0
}

func TestNewBank(t *testing.T) {

	Convey("Newly registered roll should have designated capacity.", t, func() {
	})

	Convey("Newly registered roll should have default capacity.", t, func() {
	})

	Convey("Should return empty array when no elements match predicate.", t, func() {
		roll := SomeRoll()
		roll.Add(6,1,3,5,7,42)
		items := roll.Where(isNegative)

		So(items, ShouldNotBeNil)
		So(items.Length(), ShouldEqual, 0)
	})

	Convey("Should find each of the even elements.", t, func() {
		roll := SomeRoll().Add(6,1,3,5,7,42).Where(isEven)
		So(roll.Length(), ShouldEqual, 2)
	})

	Convey("Should *NOT* find any item given the definition of the Predicate.", t, func() {
		roll := SomeRoll()
		roll.Add(1,3,5,7)
		_,ok := roll.Find(func(e interface{}) bool {
			n,_ := e.(int)
			return n == 11
		})

		So(ok, ShouldBeFalse)
	})

	Convey("Should find the item for which the Predicate returns true.", t, func() {
		roll := SomeRoll().Add(1,3,5,7)
		item,ok := roll.Find(func(e interface{}) bool {
			n,_ := e.(int)
			return n == 5
		})

		So(ok, ShouldBeTrue)
		So(item, ShouldEqual, 5)
	})

	Convey("Mapping new function produces all evens.", t, func() {
		roll := SomeRoll().Add(1,3,5,7).Map(
			func(i int, e interface{}) interface{} {
				n,_ := e.(int)
				return n+1
			})

		allEven := roll.All(isEven)

		So(allEven, ShouldBeTrue)
	})

	Convey("Bank with no items added should be empty.", t, func() {
		roll := SomeRoll()
		So(roll.IsEmpty(), ShouldBeTrue)
	})

	Convey("Bank with a few items added should not be empty.", t, func() {
		roll := SomeRoll().Add(2, 4, 6).Insert(1, 42)

		So(roll.IsEmpty(), ShouldBeFalse)
	})

	Convey("All() items in the slice should be even.", t, func() {
		roll := SomeRoll().Add(2, 4, 6).Insert(1, 42)

		var fn Pred = func(item interface{}) bool {
			n,ok := item.(int)
			return ok && (n % 2 == 0)
		}

		So(roll.All(fn), ShouldBeTrue)
	})

	Convey("Clear() should empty out the array.", t, func() {
		roll := SomeRoll().Add(1, 2, 3).Insert(1, 42)

		So(roll.Length(), ShouldEqual, 4)
		roll = roll.Clear()
		So(roll, ShouldNotBeNil)
		So(roll.Length(), ShouldEqual, 0)
	})

	Convey("Should be able to update item at an index", t, func() {
		roll := SomeRoll().Add(1, 2, 3).Insert(1, 42)

		So(roll.Length(), ShouldEqual, 4)
	})

	Convey("Should be able to add items to new registered roll", t, func() {
		bank := NewBank("items-1")
		roll,_ := bank.Register(RollDef{
			Name: "roll-1",
			Type: "unknown",
			InitialSize:100,
		})

		So(roll.Length(), ShouldEqual, 0)
		So(roll.Add(1, 2, 3).Length(), ShouldEqual, 3)

		bank.SaveToFile(Join(target, "t3/t3.json"))
	})

	Convey("Should read in newly minted and saved Bank instance", t, func() {
		bank := NewBank("t2.json")
		file := Join(target, "t2/t2.json")
		err := bank.SaveToFile(file)
		Dump(err, file)

		bank,err = LoadBank(file)
		Dump(err, file)

		So(bank.Name(), ShouldEqual, "t2.json")
	})

	Convey("Should newly minted Bank instance to file", t, func() {
		bank := NewBank("test-users.json")
		file := Join(target, "t1/t1.json")
		err := bank.SaveToFile(file)
		Dump(err, file)

		So(err, ShouldBeNil)

		b,err := fileExists(file)
		Dump(err, file)

		So(err, ShouldBeNil)
		So(b, ShouldBeTrue)
	})

	Convey("Should not throw exception for NewBank", t, func() {
		now := time.Now()
		bankName := "users"
		bank := NewBank(bankName)
		So(bank, ShouldNotBeNil)
		So(bank.Name(), ShouldEqual, bankName)
		So(bank.CreatedOn(), ShouldNotBeNil)
		So(bank.UpdatedOn(), ShouldNotBeNil)
		So(bank.UpdatedOn().Unix(), ShouldBeLessThanOrEqualTo, now.Unix())
	})
}
