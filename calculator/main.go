package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Calculator!");


	var num1 int;
	var num2 int;
	var operator string;
	var result int;

	fmt.Println("Enter first number: ");
	fmt.Scanln(&num1);

	fmt.Println("Enter second number: ");
	fmt.Scanln(&num2);

	fmt.Println("Enter operator: ");
	fmt.Scanln(&operator);

	if(operator == "+") {
		result = num1 + num2;
	} else if(operator == "-") {
		result = num1 - num2;
	} else if(operator == "*") {
		result = num1 * num2;
	} else if(operator == "/") {
		result = num1 / num2;
	}else {
		fmt.Println("Invalid operator");
		return;
	}
	fmt.Println("Result: ", result);

}