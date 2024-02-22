package main

import "testing"

/*
To made test in Go we can use the built in modules, one of them is "testing" and a simple example
of a unit test is show in the following code

1. first you need to create a file named <file_name>_test.go
2. create your function where you get one parameter and is a pointer to an struct testing.T
3. in the followinf code we show how create multiple test cases or a single test

after write all your test you can use the command: go test an go going throuht all files that
match with the pattern and execute their functions that start with Test
*/

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	}
	// table of test cases
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

/*
Coverage code: go implements a module thta let us validate which queantity of code
we test, the idea is to have a 100%

the comands are

1. go test -cover [it show in porcentage which code are test]
2. go test -coverprofile=coverage.out [generate a file with the cover]
3. go tool covert -html=coverage.out [show the output more readeable also you can use func instead of html]
*/
func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{20, 10, 20},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("The value returned not correspond with the max, got %d expect %d", max, item.n)
		}
	}
}

/*PROFILING
profiling let you know how is your code behavior in your machine
the way that we can use it is

1. run de comamand -> go test -cpuprofile=cpu.out [this going to generate a new file that contain de information]
2. go tool pprof cpu.out [this command let you know the results of your test]
	- inside the cmd in the command you can use top to see how is the programs behaviors
	- list <function_name> let you know the beavior for the especific func
	- you can see it in we or pdf using the command web and pdf (:P)
	- quit or Ctrl + D let you out from the pprof CLI

*/
func TestFibonacci(t *testing.T) {
	table := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range table {
		result := Fibonacci(item.a)
		if result != item.n {
			t.Errorf("The result of Fibonacci is incorrect, got %d expect %d", result, item.n)
		}
	}
}
