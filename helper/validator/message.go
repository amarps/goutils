package validator

import "fmt"

func GetMessage(field string, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprint("The ", field, " field is required")
	case "email":
		return fmt.Sprint("The ", field, " field must be a valid email address.")
	case "min":
		return fmt.Sprint("The ", field, " field is not met minimum length required.")
	case "max":
		return fmt.Sprint("The ", field, " field is exceeds the maximum required length.")

	default:
		return tag
	}
}
