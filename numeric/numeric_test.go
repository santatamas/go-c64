package numeric

import (
	"testing"
)

func TestCIA_SetKey_SetsMatrixProperly_SingleKey(t *testing.T) {
	// Arrange
	testByte := byte(20) // 0001 0100
	/*for i := 0; i < 8; i++ {
		fmt.Println(i)
		fmt.Println(GetBit(testByte, byte(i)))
	}*/

	// Act
	result := GetBit(testByte, byte(2))

	// Assert

	if !result {
		t.Error("Getbit returned false, expected true")
	}
}
