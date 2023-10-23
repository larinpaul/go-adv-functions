//// 2023/10/23 // 13:18 //

//// Go Advanced Functions

//// First Class and Higher Order Functions

// A programming langauge is said to have "first-class functions" when functions in that language
// are treated like any other variable. For exampole, in such a language,
// a function can be passed as an argument to other functions, can be returned by
// another function and can be assigned as a value to a variable.

// A function that returns a function or accepts a function as input is called a
// Higher-Order Function.

// Go supports first-class and higher-order functions. Another way to think of this is
// that a function is jsut another type -- just like ints and strings and bools.

// // For example, to accept a function as a parameter:

// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func mul(x, y int) int {
// 	return x * y
// }

// // aggregate applies the given math function to the first 3 inputs
// func aggregate(a, b, c int, arithmetic func(int, int) int) int {
// 	return arithmetic(arithmetic(a, b), c)
// }

// func main() {
// 	fmt.Println(aggregate(2, 3, 4, add))
// 	// prints 9
// 	fmt.Println(aggregate(2, 3, 4, mul))
// 	// prints 24
// }

//// Assignment

// // Textio is launching a new email messaging product, "Mailio"!

// // Fix the compile-time bug in the getFormattedMessages function. The function
// // body is correct, but the function signature is not.

// package main

// import "fmt"

// func getFormattedMessages(messages []string, formatter func(string) string) []string {
// 	formattedMessages := []string{}
// 	for _, message := range messages {
// 		formattedMessages = append(formattedMessages, formatter(message))
// 	}
// 	return formattedMessages
// }

// func addSignature(message string) string {
// 	return message + " Kind regards."
// }

// func addGreeting(message string) string {
// 	return "Hello! " + message
// }

// func test(messages []string, formatter func(string) string) {
// 	defer fmt.Println("===================================")
// 	formattedMessages := getFormattedMessages(messages, formatter)
// 	if len(formattedMessages) != len(messages) {
// 		fmt.Println("The number of messages returned is incorrect.")
// 		return
// 	}
// 	for i, message := range messages {
// 		formatted := formattedMessages[i]
// 		fmt.Printf(" * %s -> %s\n", message, formatted)
// 	}
// }

// func main() {
// 	test([]string{
// 		"Thanks for getting back to me,",
// 		"Great to see you again.",
// 		"I would love to hand out this weekend.",
// 		"Got any hot stock tips?",
// 	}, addSignature)
// 	test([]string{
// 		"Thanks for getting back to me.",
// 		"Great to see you again.",
// 		"I would love to hang out this weekend.",
// 		"Gor any hot stock tips?",
// 	}, addGreeting)
// }

//// 13:36

//// Currying

// // Function currying is the practice of writing a function that takes a function (or
// // functions) as input, and returns a new function.

// // For example:

// package main

// import "fmt"

// func main() {
// 	squareFunc := selfMath(multiply)
// 	doubleFunc := selfMath(add)

// 	fmt.Println(squareFunc(5))
// 	// prints 25

// 	fmt.Println(doubleFunc(5))
// 	// prints 10
// }

// func multiply(x, y int) int {
// 	return x * y
// }

// func add(x, y int) int {
// 	return x + y
// }

// func selfMath(mathFunc func(int, int) int) func(int) int {
// 	return func(x int) int {
// 		return mathFunc(x, x)
// 	}
// }

// In the exampel above, the selfMath function takes in a function as its
// parameter, and returns a function that itself returns the value of running that
// input function on its parameter.

// //// Assignment

// // The Mailio API needs a very robust error-logging system so we can see when
// // things are going awry in the back-end system. We need a function that can create
// // a custom "logger" (a function that prints to the console) given a specific
// // formatter.

// // Complete the getLogger function. It should return a new function that prints
// // the formatted inputs using the given formatter function. The inputs should be
// // passed into the formatter function in order they are given to the logger
// // function.

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// // getLogger takes a function that formats two strings into
// // a single string and returns a function that formats two strings but prints
// // the result instead of returning it
// func getLogger(formatter func(string, string) string) func(string, string) {
// 	return func(first, second string) {
// 		fmt.Println(formatter(first, second))
// 	}
// }

// func test(first string, errors []error, formatter func(string, string) string) {
// 	defer fmt.Println("=====================================")
// 	logger := getLogger(formatter)
// 	fmt.Println("Logs:")
// 	for _, err := range errors {
// 		logger(first, err.Error())
// 	}
// }

// func colonDelimit(first, second string) string {
// 	return first + ": " + second
// }
// func commaDelimit(first, second string) string {
// 	return first + ", " + second
// }

// func main() {
// 	dbErrors := []error{
// 		errors.New("out of memory"),
// 		errors.New("cpu is pegged"),
// 		errors.New("networking issue"),
// 		errors.New("invalid syntax"),
// 	}
// 	test("Error on database server", dbErrors, colonDelimit)

// 	mailErrors := []error{
// 		errors.New("email too large"),
// 		errors.New("non alphanumeric symbols found"),
// 	}
// 	test("Error on mail server", mailErrors, commaDelimit)
// }

//// 13:52
//// Pause
//// 14:29

//// Defer

// The defer keyword is a fairly unique feature of Go. It allows a function to be
// executed automatically just before its enclosing function returns.

// The deferred call's arguments are evaluated immediately, but the function call is
// not executed until the surrounding funciton returns.

// Deferred functions are typically used to close database connections, file handlers
// and the like.

// For example:

// // CopyFile copies a file from srcNmae to dstName on the local filesystem.
// func CopyFile(dstName, srcName string) (written int64, err error) {

// 	// Open the source file
// 	src, err := os.Open(srcName)
// 	if err != nil {
// 		return
// 	}
// 	// Close the source file when the CopyFile function returns
// 	defer src.Close()

// 	// Create the destination file
// 	ddst, err := os.Create(dstName)
// 	if err != nil {
// 		return
// 	}
// 	// Close the destination file when the CopyFile function returns
// 	defer dst.Close()

// 	return io.Copy(dst, src)
// }

// In the above example, the src.Close() function is not called until after the
// CopyFile function was called but immediately before the CopyFile function
// returns.

// Defer is a great way to make sure that something happens at the end of a
// function, even if there are multiple return statements.

// //// Assignment

// // There is a bug in the logAndDelete function, fix it!

// // This function shoudl always delete the user from the user's map, which is a map
// // that stores the user's name as keys. It also returns a log string that indicates to
// // the caller some information about the user's deletion.

// // To avoid bugs like this in the future, instead of calling delete before each
// // return, just defer the delete once at the beginning of the function.

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// const (
// 	logDeleted  = "user deleted"
// 	logNotFound = "user not found"
// 	logAdmin    = "admin deleted"
// )

// // func logAndDelete(users map[string]user, name string) (log string) {
// // 	user, ok := users[name]
// // 	if !ok {
// // 		delete(users, name)
// // 		return logNotFound
// // 	}
// // 	if user.admin {
// // 		return logAdmin
// // 	}
// // 	delete(users, name)
// // 	return logDeleted
// // }

// func logAndDelete(users map[string]user, name string) (log string) {
// 	defer delete(users, name)

// 	user, ok := users[name]
// 	if !ok {
// 		return logNotFound
// 	}
// 	if user.admin {
// 		return logAdmin
// 	}
// 	return logDeleted
// }

// type user struct {
// 	name   string
// 	number int
// 	admin  bool
// }

// func test(users map[string]user, name string) {
// 	fmt.Printf("Attempting to delete %s...\n", name)
// 	defer fmt.Println("======================================")
// 	log := logAndDelete(users, name)
// 	fmt.Println("Log:", log)
// }

// func main() {
// 	users := map[string]user{
// 		"john": {
// 			name:   "john",
// 			number: 18965554631,
// 			admin:  true,
// 		},
// 		"elon": {
// 			name:   "elon",
// 			number: 19875556452,
// 			admin:  true,
// 		},
// 		"breanna": {
// 			name:   "breanna",
// 			number: 98575554231,
// 			admin:  false,
// 		},
// 		"kade": {
// 			name:   "kade",
// 			number: 10765557221,
// 			admin:  false,
// 		},
// 	}

// 	fmt.Println("Initial users:")
// 	usersSorted := []string{}
// 	for name := range users {
// 		usersSorted = append(usersSorted, name)
// 	}
// 	sort.Strings(usersSorted)
// 	for _, name := range usersSorted {
// 		fmt.Println(" -", name)
// 	}
// 	fmt.Println("=====================================")

// 	test(users, "john")
// 	test(users, "santa")
// 	test(users, "kade")

// 	fmt.Println("Final users:")
// 	usersSorted = []string{}
// 	for name := range users {
// 		usersSorted = append(usersSorted, name)
// 	}
// 	sort.Strings(usersSorted)
// 	for _, name := range usersSorted {
// 		fmt.Println(" -", name)
// 	}
// 	fmt.Println("======================================")
// }

//// 14:56

//// Closures

// // A clouse is a function that references variables from outside its own function
// // body. The function may access adn assign to the freferenced variables.

// // In this example, the concatter() function returns a functio ntaht has reference
// // to an eclosed doc value. Each successive call to harryPotterAggregator
// // mutates that same doc variable.

// package main

// import "fmt"

// func concatter() func(string) string {
// 	doc := ""
// 	return func(word string) string {
// 		doc += word + " "
// 		return doc
// 	}
// }

// func main() {
// 	harryPotterAggregator := concatter()
// 	harryPotterAggregator("Mr.")
// 	harryPotterAggregator("and")
// 	harryPotterAggregator("Mrs.")
// 	harryPotterAggregator("Dursley")
// 	harryPotterAggregator("of")
// 	harryPotterAggregator("number")
// 	harryPotterAggregator("four,")
// 	harryPotterAggregator("Privet")

// 	fmt.Println(harryPotterAggregator("Drive"))
// 	// Mr. and Mrs. Dursley of number four, Privet Drive
// }

//// Assignment

// // Keeping track of how many emails we send is mission-critical at Mailio. Complete
// // the adder() function.

// // It should return a function that adds its input (an int) to an eclosed sum
// // value, then return the new sum. In other words, it keeps a running total of the
// // sum variable within a closure.

// package main

// import "fmt"

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// type emailBill struct {
// 	costInPennies int
// }

// func test(bills []emailBill) {
// 	defer fmt.Println("============================")
// 	countAdder, costAdder := adder(), adder()
// 	for _, bill := range bills {
// 		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
// 	}
// }

// func main() {
// 	test([]emailBill{
// 		{45},
// 		{32},
// 		{43},
// 		{12},
// 		{34},
// 		{54},
// 	})

// 	test([]emailBill{
// 		{12},
// 		{12},
// 		{976},
// 		{12},
// 		{543},
// 	})

// 	test([]emailBill{
// 		{743},
// 		{13},
// 		{8},
// 	})
// }

//// 15:06
//// Pause
//// 16:02

//// Anonymous Function

// Anonymous function are true to form in that they have no name. We've been
// using them throughout this chapter, but we haven't really talked about them yet.

// Anonymous functions are true to form in that they have no name. We've been
// using them througout this chapter, but we haven't really talked about them yet.

// // Anonymous functions are useful when defining a function that will only be used
// // once or to create a quick closure.

// // doMath accepts a function that converts one int into another
// // and a slice of ints. It returns a slice of ints that have been
// // converted by the passed in functions.
// package main

// import "fmt"

// func doMath(f func(int) int, nums []int) []int {
// 	var results []int
// 	for _, n := range nums {
// 		results = append(results, f(n))
// 	}
// 	return results
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}

// 	// Here we define an anonymosu function that doubles an int
// 	// and pass it to doMath
// 	allNumsDoubles := doMath(func(x int) int {
// 		return x + x
// 	}, nums)

// 	fmt.Println(allNumsDoubles)
// 	// prints:
// 	// [2 4 6 8 10]

// }

//// Assignment

// Complete the printReports function.

// Call printCostReport once for each message. Pass in an anonymosu function as
// the costCalculator that returns an int equal to twice the length of the input
// message.

package main

import "fmt"

func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func(m string) int {
			return len(m) * 2
		}, message)
	}
}

func test(messages []string) {
	defer fmt.Println("================================")
	printReports(messages)
}

func main() {
	test([]string{
		"Here's Johny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no palce like home",
	})

	test([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
