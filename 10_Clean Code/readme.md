10. Clean Code

- Clean Code : Clean code adalah sebuah script komputer yang mudah dipahami oleh pembuat program secara baik dan optimal

Aturan :
1. Use clear and concise variable names :
Variable names should be clear and concise, indicating what the variable represents. Avoid using single-letter variable names or names that are too generic. A good rule of thumb is to use variable names that are descriptive but not too long. For example:
//bad
var x int
//good
var numEmployees int

2. Use meaningful function and method names
Function and method names should be descriptive and accurately describe what the function or method does. Use verbs to describe the action the function or method performs. For example:

// bad
func x() {}
// good
func calculateAverage(nums []float64) float64 {}

3. Follow the Single Responsibility Principle
Each function or method should have only one responsibility. If a function or method is doing more than one thing, it becomes harder to read and understand. It also makes it more difficult to test and maintain. For example:

// bad
func calculateAverageAndSum(nums []float64) (float64, float64) {}
// good
func calculateAverage(nums []float64) float64 {}
func calculateSum(nums []float64) float64 {}

4. Use comments sparingly
Comments should be used to explain why something is done, not how it is done. Code should be self-explanatory, and comments should only be used when necessary. Avoid using comments to explain what the code is doing. For example:

// bad
// loop through the array
for i := 0; i < len(arr); i++ {}
// good
for index := 0; index < len(array); index++ {}

5. Avoid using global variables
Global variables can make code harder to test and maintain, as they can be modified from anywhere in the code. Instead, use local variables or pass variables as parameters to functions or methods. For example:

// bad
var count int
func incrementCount() {
    count++
}
// good
func incrementCount(count int) int {
    return count + 1
}

6. Handle errors properly
Always handle errors in your code. Ignoring errors can lead to unexpected behavior and make your code harder to debug. Use Go’s built-in error handling mechanism to handle errors in your code. For example:

// bad
f, err := os.Open("filename.txt")
if err == nil {
    // do something with f
}
// good
f, err := os.Open("filename.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
// do something with f

7. Use interfaces to decouple dependencies
Using interfaces can help decouple dependencies in your code. This makes it easier to test and maintain your code. For example:

// bad
func calculateAverage(nums []float64) float64 {
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return sum / float64(len(nums))
}
// good
type Calculator interface {
    Calculate(nums []float64) float64
}
type AverageCalculator struct {}
func (c *AverageCalculator) Calculate(nums []float64) float64 {
    sum := 0.0
    for _, num := range nums {
      sum += num
    }
   return sum / float64(len(nums))
}

8. Use proper formatting
Proper formatting makes your code more readable and easier to understand. Use Go’s built-in formatting tools to ensure that your code is properly formatted. For example:

// bad
func calculateAverage(nums []float64) float64 {
sum := 0.0
for _, num := range nums {
sum += num
}
return sum / float64(len(nums))
}
// good
func calculateAverage(nums []float64) float64 {
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return sum / float64(len(nums))
}

9. Keep functions and methods short
Functions and methods should be short and concise. This makes them easier to read and understand. A good rule of thumb is to keep functions and methods to less than 50 lines of code. If a function or method is longer than that, consider breaking it up into smaller, more manageable functions or methods. For example:

// bad
func calculateAverageAndSum(nums []float64) (float64, float64) {
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    average := sum / float64(len(nums))
    return average, sum
}
// good
func calculateAverage(nums []float64) float64 {
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return sum / float64(len(nums))
}
func calculateSum(nums []float64) float64 {
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return sum
}
10. Write testable code
Write code that is easy to test. This makes it easier to identify and fix bugs in your code. Use the “testing” package provided by Go to write tests for your code. For example:

// bad
func calculateAverage(nums []float64) float64 {
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return sum / float64(len(nums))
}
// good
func calculateAverage(nums []float64) float64 {
    if len(nums) == 0 {
        return 0.0
    }
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    return sum / float64(len(nums))
}
func TestCalculateAverage(t *testing.T) {
    nums := []float64{1.0, 2.0, 3.0, 4.0}
    expected := 2.5
    result := calculateAverage(nums)
    if result != expected {
        t.Errorf("Expected %f but got %f", expected, result)
    }
}
Conclusion
Writing clean code is essential for creating high-quality, maintainable software. Follow these 10 tips to write clean code in Go that is easy to read, understand, and maintain. By doing so, you will improve the quality of your code, reduce bugs, and make your code more testable.