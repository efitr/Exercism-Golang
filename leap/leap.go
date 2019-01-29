// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	// Meaningful rule:
	//    - on every year that is evenly divisible by 4
	//    - except every year that is evenly divisible by 100
	//    - unless the year is also evenly divisible by 400
	
	//First if year not divisable by 4, false
		if year % 4 == 0{
			if  % 100 {
				return false 
			}
			return true
		}
	return false
}
