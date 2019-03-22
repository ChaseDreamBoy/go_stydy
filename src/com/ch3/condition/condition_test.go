package condition

import "testing"

func TestMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1 == 1")
	}
}

func TestMultiSwitchTry(t *testing.T) {
	for i := 0; i < 6; i++ {
		switch i {
		case 0, 2:
			t.Log("event", i)
		case 1, 3:
			t.Log("Odd", i)
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 6; i++ {
		switch {
		case i%2 == 0:
			t.Log("event", i)
		case i%2 == 1:
			t.Log("Odd", i)
		default:
			t.Log("un know")
		}
	}
}
