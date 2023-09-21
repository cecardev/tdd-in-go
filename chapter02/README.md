## Questions
In Go, what is the difference between a module and a package?
Package: Is a collection of source files in same directory that are compiled together
Module: Is a collection of related Go packages that are versioned together

What is the additional test package? What are some of the advantages of using it?
A framework for writing unit test for Go code
Adventages:
Prevents brittle tests
Separates test and core package dependencies
allows developers to integrate with their own packages

What are the requirements for the test signature?
Tests are exported functions whose name beigns with "Test"
Test can have a suffix that specifies what is covering
Tests must take in a single parameter of the *testing.T type
Tests must not havea a return type

What are subtests and how do you create them?
Substests in go allow you to group related tests together and report their results separately
You can create  using the t.Run() method of the *testing.T parameter.


What is a benchmark? How do you write one?
Is a function that measures the performance of a piece of code.
You can create with the signature "func BenchmarkXXX(b *testing.B), where "XXX" is the name of the benchmark function. The *testing.B is used to control the number of iterations and to report benchmark results.


