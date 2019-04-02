package union

import (
	"testing"
)

func TestUnion(t *testing.T) {
	u := NewUnion(10)

	u.initArr()

	u.union(0, 9)
	u.union(0, 8)
	u.union(8, 7)
	u.union(2, 7)

	if !u.connected(2, 9) {
		t.Errorf("2 and 9 not connected")
	}
}

func TestQuickUnion(t *testing.T) {
	u := NewUnion(20)

	u.initArr()
	u.quickUnion(0, 9)
	if !u.quickConnect(0, 9) {
		t.Errorf("2 and 8 not connected")
	}
	u.quickUnion(0, 8)
	if !u.quickConnect(9, 8) {
		t.Errorf("2 and 8 not connected")
	}
	u.quickUnion(8, 7)
	if !u.quickConnect(0, 8) {
		t.Errorf("2 and 8 not connected")
	}
	u.quickUnion(2, 7)
	if !u.quickConnect(7, 9) {
		t.Errorf("2 and 8 not connected")
	}
	u.quickUnion(9, 17)
	u.quickUnion(14, 13)
	u.quickUnion(12, 19)
	u.quickUnion(11, 9)
	u.quickUnion(12, 8)

	if !u.quickConnect(0, 19) {
		t.Errorf("0 and 19 not connected")
	}

}
