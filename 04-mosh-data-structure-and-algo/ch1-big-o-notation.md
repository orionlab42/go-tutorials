Ch1 - The Big O Notation
-----------------------

# 1.1 Definition
* Big O Notation is a mathematical notation that describes the limiting behavior of a function when the
argument tends towards a particular value or infinity.(Wikipedia)
* We use it to describe the performance of an algorithm (to see how it will perform as the input grows 
large, scalability), basically the runtime complexity

# 1.2 Types

* CONSTANT GROWTH RATE: O( 1 ) - means that the method/algo runs in a constant time
* LINEAR GROWTH RATE: O( n ) At log3->log6 method the size of the input matters, the cost of this algo grows linearly and in direct 
* correlations with the size of the input

```go
func log1(int[] numbers) {
	// O(1)
	fmt.Println(numbers[0]); // no mather how big the array is, we are just printing the first item
}

func log2(int[] numbers) {
	// O(2), but we can simplify this by writing down O(1), meaning constant time
	fmt.Println(numbers[0]); 
	fmt.Println(numbers[0]); 
}

func log3(int[] numbers) {
	// O(n), where n represent the size of the input
	for _, num := range numbers {
        fmt.Println(num);
    }
}

func log4(int[] numbers) {
	// O( 2+n ), but we can drop the constant to O(n)
	fmt.Println(numbers[0]);
    for _, num := range numbers {
            fmt.Println(num);
        }
    fmt.Println(numbers[0]);
}

func log5(int[] numbers) {
	// O( 2n ), but we can simplify to O(n), because 2n represents still a linear growth
    for _, num := range numbers {
            fmt.Println(num);
        }
    for _, num := range numbers {
            fmt.Println(num);
        }
}

func log6(int[] numbers, string[] names) {
	// O( n+m ), but we can simplify to O(n), because n+m represents still a linear growth
    for _, num := range numbers {
            fmt.Println(num);
        }
    for _, name := range names {
            fmt.Println(name);
        }
}
```

* QUADRATIC GROWTH RATE: O( n^2 ) At log7 the function, because there is a nested loop, has a quadratic growth. 
* An algo which has the O(n^2) is slower than the one that has O(n)
* log9 with O(n^3) gets far more slow than an algo with O(n)

```go
func log7(int[] numbers) {
	// O( n*n ) or  O(n^2) has a quadratic growth
    for _, first := range numbers { // O(n)
        for _, second := range numbers { // O(n)
            fmt.Println(first + ", " + second);
        }
	}
}

func log8(int[] numbers) {
	// O( n^2 + n ), this can be simplified to O( n^2 )
    for _, first := range numbers { // O(n)
        for _, second := range numbers { // O(n)
            fmt.Println(first + ", " + second);
        }
	}
    for _, num := range numbers { // O(n)
        fmt.Println(num);
    }
}

func log9(int[] numbers) {
	// O( n^3 )
    for _, first := range numbers { // O(n)
        for _, second := range numbers { // O(n)
            for _, third := range numbers { // O(n)
                fmt.Println(first + ", " + second + ", " + third);
            }
        }
	}
}
```

* LOGARITHMIC GROWTH RATE: O( log n ) An algorithm that runs in a logarithmic time is more efficient and more scalable,
  than an algo that runs in linear or quadratic time
* The logarithmic curve slows down as the input size grows
* Example of logarithmic algo is the binary search in case of ordered arrays, which compared
  to linear search is much more optimal, how it works is that you start in the middle and inspect
  only the half which contains the searched item and so on, every time getting the half of the
  remaining half, narrowing down the search by half
* EXPONENTIAL GROWTH RATE: O( 2^n ) The opposite of a logarithmic curve is the exponential curve O( 2^n ), so an algo that runs in 
  exponential time is not scalable at all, it will slow down very soon

# 1.3 Space complexity
* Most of the time is a tradeoff between saving time or saving space
* The big O notation is also used to measure how much space an algorithm requires

```go
func log10(string[] names) {
	// O(1) space
    for _, name := range names { // O(n)
         fmt.Println("HI " + name);
    }
}
func log11(string[] names) {
	// O(n) space, in direct proportion with the size of the input, because of the array what we create
    var copy [len(names)]string;
    for _, name := range names { // O(n)
         fmt.Println("HI " + name);
    }
}
```

* We don't count the input size that is always n, we count how much 
extra space we need for the output