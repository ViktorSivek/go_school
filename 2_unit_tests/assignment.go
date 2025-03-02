package assignment

// Student
type Student interface {
	Name() string
}

// Course
type Course interface {
	Name() string
	EnrollStudent(s Student) error
}

// DataSource
type DataSource interface {
	ReadStudent(studentID int) (Student, error)
	ReadCourse(courseID int) (Course, error)
}

// Reverse returns a new slice with the elements in reverse order
func Reverse(s []string) []string {
	n := len(s)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = s[n-1-i]
	}
	return result
}

// Palindrome checks if a slice reads the same forward and backward
func Palindrome(s []string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// Anagram checks if two strings contain the same characters in a different order
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charCount := make(map[rune]int)

	for _, char := range s1 {
		charCount[char]++
	}

	for _, char := range s2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}

// RemoveDigits removes all digits from a string
func RemoveDigits(s string) string {
	result := ""
	for _, char := range s {
		if char < '0' || char > '9' {
			result += string(char)
		}
	}
	return result
}

// ReplaceDigits replaces all digits in a string with the replacement string
func ReplaceDigits(s string, r string) string {
	result := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			result += r
		} else {
			result += string(char)
		}
	}
	return result
}

// EnrollStudentToCourse enrolls a student into a course
func EnrollStudentToCourse(dataSource DataSource, studentID int, courseID int) error {
	student, err := dataSource.ReadStudent(studentID)
	if err != nil {
		return err
	}
	
	course, err := dataSource.ReadCourse(courseID)
	if err != nil {
		return err
	}
	
	return course.EnrollStudent(student)
}