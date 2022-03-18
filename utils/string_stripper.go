package utils

// NameStripper
///*Just a function that takes string strips it off of last 7 char position.
func NameStripper(s string) string {

	stringLength := len(s) - 7
	s = s[:stringLength]
	return s

}

// YearStripper
///*Another function to string year off of string */
func YearStripper(s string) string {
	yearStr := s[:4]
	return yearStr
}
