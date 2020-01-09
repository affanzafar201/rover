package main

import "testing"

func TestMoveOneStep(t *testing.T) {
	//      W                  |
	//   S      N       -------|--------x
	//		E	               |
	//                         y
	tests := []struct {
		curX          int64
		curY          int64
		curOrient     string
		endX          int64
		endY          int64
		lenX          int64
		lenY          int64
		expectedError bool
	}{
		{1, 2, "N", 1, 3, 5, 5, false},
		{1, 2, "E", 0, 2, 5, 5, false},
		{1, 2, "S", 1, 1, 5, 5, false},
		{1, 2, "W", 2, 2, 5, 5, false},
		{0, 0, "S", 0, -1, 5, 5, true},
	}
	directionStrToInt := map[string]int{"N": 0, "E": 1, "S": 2, "W": 3}
	for _, test := range tests {

		err := moveOneStep(&test.curX, &test.curY, directionStrToInt[test.curOrient], test.lenX, test.lenY)

		if (err != nil) != test.expectedError {
			t.Errorf("moveOneStep failed, expected error: %v and got error: %v", test.expectedError, (err != nil))
			return
		}
		if err != nil {
			t.Logf("moveOneStep success, expected and got error: %v and  error: %v", test.expectedError, err.Error())
		} else if test.curX != test.endX && test.curY != test.endY {
			t.Errorf("moveOneStep failed, expected curX: %v and curY: %v, got curX: %v and curY: %v", test.endX, test.endY, test.curX, test.curY)

		} else {
			t.Logf("moveOneStep success, expected curX: %v and curY: %v, got curX: %v and curY: %v", test.endX, test.endY, test.curX, test.curY)
		}
	}
}

func TestMoveRover(t *testing.T) {
	//      W                  |
	//   S      N       -------|--------x
	//		E	               |
	//                         y
	tests := []struct {
		curX          int64
		curY          int64
		curOrient     string
		movements     string
		endX          int64
		endY          int64
		finalOrient   string
		lenX          int64
		lenY          int64
		expectedError bool
	}{
		{1, 2, "N", "LMLMLMLMM", 1, 3, "N", 5, 5, false},
		{3, 3, "E", "MMRMMRMRRM", 0, 2, "E", 5, 5, false},
		{0, 0, "E", "MMMM", -1, 0, "E", 5, 5, true},
		{3, 5, "N", "M", 3, 6, "N", 5, 5, true},
	}
	for _, test := range tests {

		finalPosX, finalPosY, finalDir, err := moveRover(test.curX, test.curY, test.curOrient, test.movements, test.lenX, test.lenY)

		if (err != nil) != test.expectedError {
			t.Logf("%v %v", finalPosX, finalPosY)
			t.Errorf("moveRover failed, expected error: %v and got error: %v ", test.expectedError, (err != nil))
			return
		}
		if err != nil {
			t.Logf("moveRover success, expected and got error: %v and  error: %v", test.expectedError, err.Error())
		} else if finalPosX != test.endX && finalPosY != test.endY && finalDir != test.finalOrient {
			t.Errorf("moveRover failed, expected finalPosX: %v, finalPosY: %v and finalDir: %v, got finalPosX: %v, finalPosY: %v and finalDir: %v", test.endX, test.endY, test.finalOrient, finalPosX, finalPosY, finalDir)
		} else {
			t.Logf("moveRover sunccess, expected finalPosX: %v, finalPosY: %v and finalDir: %v, got finalPosX: %v, finalPosY: %v and finalDir: %v", test.endX, test.endY, test.finalOrient, finalPosX, finalPosY, finalDir)
		}
	}
}
