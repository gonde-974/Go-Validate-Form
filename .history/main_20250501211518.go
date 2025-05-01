package main

import (
	"fmt" //Packages we wil used:

	"strings" //For printing

	"unicode" // For character validation
)

// What we validate

// 1.Name : Must have at least 2 characters
// 2.Email: Must Contains "@" and "."
// 3.Password
//. at least 8 caracters
//. at least one number
//. at least one lower leter and Upper leter

//Name validation

func isValidName(name string) bool {
	return len(name) >= 2
}


//Email validation

func isValidEmail(email string) bool{
	return strings.Contains(email,"@") && strings.Contains(email,".")
}

//Password validation

func isValidPassword (password string) bool{
	var hasUpper , hasLower , hasDigit bool

	if len(password) < 8 {
		return false
	}

	for _, char := range password{

		switch {
		case unicode.IsUpper(char):
			hasUpper = true

		case unicode.IsLower(char):
			hasLower = true

		case unicode.IsDigit(char):
			hasDigit = true

		}

		
	}
	return hasUpper && hasLower && hasDigit
}

//main function === form simulation
func main (){

	name := "Ana"
	email := "ana@examle.com"
	password := "A1234567a"

	fmt.Println("============= Valid Form============")

	if !isValidName(name){
		fmt.Println("Name must have at least 2 character")
	}else if !isValidEmail(email){
		fmt.Println("Invalid email")
	}else if !isValidPassword(password){
	    fmt.Println("Password must be at least 8 characters long, include uppercase, lowercase letters, and a number")
	}else{
		 fmt.Println("Validate vas successful")
	}

}







