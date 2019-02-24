package numeric

import (
	"testing"
)

func TestNumberic_GetBit_HappyPath(t *testing.T) {
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

func TestNumberic_isNegative_HappyPath(t *testing.T) {
	// Arrange
	testByte := byte(128) // 0001 0100
	/*for i := 0; i < 8; i++ {
		fmt.Println(i)
		fmt.Println(GetBit(testByte, byte(i)))
	}*/

	// Act
	result := IsNegative(testByte)

	// Assert

	if !result {
		t.Error("isNegative returned false, expected true")
	}
}
