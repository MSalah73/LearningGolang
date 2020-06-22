// To create a test suite - create a file whose name ends with  _test.
// The _test indicates that this file has unit tests which will be run using
// the go test command.
// Note: The Icon will change when appending the _test to the end of the file name.
package ValidationKit

import (
	"testing"
)

// To use this Suite run: go test -v - Inside the UnitTest directory

// Notice all the function start with Test word at the begining
func TestCheckUsernameSyntaxMinimumCharacterLength(t *testing.T) {

	result := CheckUernameSyntax("")

	if result != false {
		t.Errorf("Failed the minimum character check")
	}
}
func TestCheckUsernameSyntaxMaximumCharacterLength(t *testing.T) {

	result := CheckUernameSyntax("@AAABBBBCCCCDDDDE")

	if result != false {
		t.Errorf("Failed the maximun character check")
	}
}
func TestCheckUsernameSyntaxSymbols(t *testing.T) {

	result := CheckUernameSyntax("@!#$%^&*_-+=~")

	if result != false {
		t.Errorf("Failed the Special character check")
	}
}
func TestCheckUsernameSyntaxUnderscore(t *testing.T) {

	result := CheckUernameSyntax("I_HOPE_IT_WORKS")

	if result != true {
		t.Errorf("Failed the check to allow underscore characters")
	}
}
func TestCheckUsernameSyntaxAtSignInsideUsername(t *testing.T) {
	// The @ sign should only be at the begining of the username and it should fail if it appear elsewhere.
	result := CheckUernameSyntax("Going@It")

	if result != false {
		t.Errorf("Failed the @ sign check. The @ sign was found in another place besides the start of the string.")
	}
}
func TestCheckUsernameSyntaxRandomUsernames(t *testing.T) {
	for i := 0; i < 10008; i++ {
		username := GenerateRandomUsername()
		result := CheckUernameSyntax(username)
		if result != true {
			t.Errorf("The random username, %v, failed to pass the username check.", username)
			t.Fatal("Quitting")
		}
	}
}
