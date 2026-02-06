# Go Syntax Cheatsheet (for C#/.NET & JavaScript Developers)

## Variables & Constants

| C# | JavaScript | Go |
|----|------------|-----|
| `var x = 10;` | `let x = 10;` | `x := 10` |
| `int x = 10;` | `const x = 10;` | `var x int = 10` |
| `const int X = 10;` | `const X = 10;` | `const X = 10` |
| `string? x = null;` | `let x = null;` | `var x *string = nil` |

```go
// Short declaration (inferred type) - most common
name := "Chef"
count := 42

// Explicit type
var name string = "Chef"
var count int = 42

// Zero-value initialization
var name string   // ""
var count int     // 0
var ok bool       // false

// Multiple
x, y := 1, 2
var a, b, c int

// Constants
const Pi = 3.14159
const (
    StatusOK    = 200
    StatusError = 500
)
```

## Basic Types

| C# | JavaScript | Go |
|----|------------|-----|
| `int` | `number` | `int`, `int32`, `int64` |
| `float`, `double` | `number` | `float32`, `float64` |
| `bool` | `boolean` | `bool` |
| `string` | `string` | `string` |
| `byte` | - | `byte` (alias for `uint8`) |
| `object` | `any` | `any` (or `interface{}`) |

## Strings

| C# | JavaScript | Go |
|----|------------|-----|
| `$"Hello {name}"` | `` `Hello ${name}` `` | `fmt.Sprintf("Hello %s", name)` |
| `str.Length` | `str.length` | `len(str)` |
| `str.Contains("x")` | `str.includes("x")` | `strings.Contains(str, "x")` |
| `str.Split(",")` | `str.split(",")` | `strings.Split(str, ",")` |
| `str.ToUpper()` | `str.toUpperCase()` | `strings.ToUpper(str)` |
| `String.Join(",", arr)` | `arr.join(",")` | `strings.Join(arr, ",")` |

```go
import "strings"
import "fmt"

name := "World"
msg := fmt.Sprintf("Hello %s!", name)     // "Hello World!"
msg := "Hello " + name + "!"              // concatenation

// Multi-line (raw string)
query := `
    SELECT *
    FROM users
    WHERE active = true
`
```

## Functions

| C# | JavaScript | Go |
|----|------------|-----|
| `int Add(int a, int b)` | `function add(a, b)` | `func Add(a, b int) int` |
| `void Log(string s)` | `function log(s)` | `func Log(s string)` |
| `(int, string) Get()` | `return [1, "a"]` | `func Get() (int, string)` |
| `params int[] nums` | `...nums` | `nums ...int` |

```go
// Basic function
func Add(a, b int) int {
    return a + b
}

// Multiple return values (common pattern)
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func Split(sum int) (x, y int) {
    x = sum / 2
    y = sum - x
    return  // returns x, y implicitly
}

// Variadic (like params/rest)
func Sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
Sum(1, 2, 3, 4)

// Function as value (like delegates/callbacks)
var operation func(int, int) int
operation = Add
result := operation(2, 3)

// Anonymous function / lambda
double := func(x int) int { return x * 2 }

// Closure
func Counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Control Flow

### If/Else
```go
// Standard
if x > 0 {
    // positive
} else if x < 0 {
    // negative
} else {
    // zero
}

// With initialization (scoped variable) - very common in Go
if err := doSomething(); err != nil {
    return err
}

// C#/JS equivalent would be:
// var err = doSomething(); if (err != null) { ... }
```

### Switch
```go
// No break needed - no fallthrough by default
switch day {
case "Mon", "Tue", "Wed", "Thu", "Fri":
    fmt.Println("Weekday")
case "Sat", "Sun":
    fmt.Println("Weekend")
default:
    fmt.Println("Invalid")
}

// Switch with no condition (cleaner if/else chains)
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
default:
    grade = "C"
}

// Type switch (like C# pattern matching)
switch v := value.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown")
}
```

### Loops
```go
// Go only has 'for' - no while/do-while

// Standard for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style
for count > 0 {
    count--
}

// Infinite loop
for {
    // break to exit
}

// Range (like foreach / for...of)
for index, value := range slice {
    fmt.Println(index, value)
}

// Ignore index
for _, value := range slice {
    fmt.Println(value)
}

// Ignore value (just iterate)
for i := range slice {
    fmt.Println(i)
}

// Range over map
for key, value := range myMap {
    fmt.Println(key, value)
}
```

## Collections

### Arrays (fixed size - rarely used directly)
```go
var arr [5]int                    // [0,0,0,0,0]
arr := [3]int{1, 2, 3}            // [1,2,3]
arr := [...]int{1, 2, 3, 4, 5}    // size inferred
```

### Slices (dynamic - like List<T> or Array)

| C# | JavaScript | Go |
|----|------------|-----|
| `new List<int>()` | `[]` | `[]int{}` or `make([]int, 0)` |
| `list.Add(x)` | `arr.push(x)` | `slice = append(slice, x)` |
| `list.Count` | `arr.length` | `len(slice)` |
| `list[0]` | `arr[0]` | `slice[0]` |
| `list.Skip(1).Take(3)` | `arr.slice(1, 4)` | `slice[1:4]` |

```go
// Create
nums := []int{1, 2, 3}
nums := make([]int, 0)        // empty with len=0
nums := make([]int, 5)        // [0,0,0,0,0]
nums := make([]int, 0, 100)   // empty, capacity 100

// Append (returns new slice!)
nums = append(nums, 4)
nums = append(nums, 5, 6, 7)
nums = append(nums, other...)  // spread

// Slicing [start:end] - end is exclusive
first3 := nums[:3]    // [0:3]
last3 := nums[2:]     // [2:len]
middle := nums[1:4]
copy := nums[:]       // shallow copy reference

// Actually copy
copy := make([]int, len(nums))
copy(copy, nums)
```

### Maps (like Dictionary<K,V> or Object/Map)

| C# | JavaScript | Go |
|----|------------|-----|
| `new Dictionary<string,int>()` | `{}` or `new Map()` | `make(map[string]int)` |
| `dict["key"] = 1` | `obj.key = 1` | `m["key"] = 1` |
| `dict.TryGetValue(k, out v)` | `obj.key ?? default` | `v, ok := m["key"]` |
| `dict.Remove("key")` | `delete obj.key` | `delete(m, "key")` |
| `dict.ContainsKey("key")` | `"key" in obj` | `_, ok := m["key"]` |

```go
// Create
m := map[string]int{}
m := make(map[string]int)
m := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,  // trailing comma required
}

// Access
value := m["key"]           // returns zero-value if missing
value, ok := m["key"]       // ok=false if missing (comma-ok idiom)

if val, ok := m["key"]; ok {
    fmt.Println("Found:", val)
}

// Set & Delete
m["key"] = 42
delete(m, "key")

// Iterate (order is random!)
for key, value := range m {
    fmt.Println(key, value)
}
```

## Structs (like Classes)

| C# | Go |
|----|-----|
| `class` | `struct` (no classes) |
| `new Person()` | `Person{}` or `&Person{}` |
| `this.` | implicit (no `this`) |
| constructors | factory functions by convention |
| inheritance | embedding (composition) |

```go
// Definition
type Person struct {
    Name string    // Exported (public) - uppercase
    age  int       // unexported (private) - lowercase
}

// Factory function (constructor pattern)
func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        age:  age,
    }
}

// Instantiation
p1 := Person{Name: "Alice", age: 30}
p2 := Person{"Bob", 25}              // positional (fragile)
p3 := &Person{Name: "Carol"}         // pointer
p4 := new(Person)                    // pointer, zero values

// Methods (receiver = implicit 'this')
func (p Person) Greet() string {
    return "Hello, " + p.Name
}

// Pointer receiver (can modify struct)
func (p *Person) Birthday() {
    p.age++
}

// Usage
person := NewPerson("Alice", 30)
person.Birthday()
fmt.Println(person.Greet())
```

### Embedding (Composition over Inheritance)
```go
// C# equivalent: class Employee : Person { }
type Employee struct {
    Person          // embedded - "inherits" fields/methods
    Title  string
}

emp := Employee{
    Person: Person{Name: "Bob", age: 30},
    Title:  "Developer",
}
emp.Name           // promoted from Person
emp.Greet()        // promoted method
emp.Person.Name    // explicit access
```

## Interfaces

| C# | Go |
|----|-----|
| `class Dog : IAnimal` | implicit (no declaration) |
| `interface IAnimal { }` | `type Animal interface { }` |
| explicit implementation | automatic if methods match |

```go
// Define interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Compose interfaces
type ReadWriter interface {
    Reader
    Writer
}

// Implement implicitly - no "implements" keyword!
type MyBuffer struct {
    data []byte
}

func (b *MyBuffer) Read(p []byte) (int, error) {
    // implementation
    return 0, nil
}
// MyBuffer now implements Reader automatically

// Empty interface = any type (like object/any)
func Print(v interface{}) {  // or: func Print(v any)
    fmt.Println(v)
}

// Type assertion
var i interface{} = "hello"
s := i.(string)              // panics if wrong type
s, ok := i.(string)          // safe - ok=false if wrong
```

## Error Handling

| C# | JavaScript | Go |
|----|------------|-----|
| `throw new Exception()` | `throw new Error()` | `return errors.New()` |
| `try { } catch { }` | `try { } catch { }` | `if err != nil { }` |
| `Exception` | `Error` | `error` (interface) |

```go
import "errors"
import "fmt"

// Return error (not throw!)
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Handle error (not try/catch!)
result, err := Divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return  // or handle appropriately
}
fmt.Println(result)

// Wrap errors (like InnerException)
if err != nil {
    return fmt.Errorf("divide failed: %w", err)
}

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Panic/Recover (like throw/catch - use sparingly!)
func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    panic("something went wrong")
}
```

## Defer (like finally/using)

```go
// Executes when function returns (LIFO order)
func ReadFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()  // guaranteed cleanup

    // ... use file
    return nil
}

// Multiple defers (stack - LIFO)
defer fmt.Println("1")
defer fmt.Println("2")
defer fmt.Println("3")
// Output: 3, 2, 1
```

## Pointers

| C# | Go |
|----|-----|
| `ref` parameters | `*Type` pointer parameter |
| `out` parameters | return multiple values |
| reference types (class) | use `*Type` explicitly |
| value types (struct) | default behavior |

```go
// Get pointer
x := 10
p := &x         // p is *int (pointer to int)

// Dereference
fmt.Println(*p) // 10
*p = 20         // x is now 20

// In functions
func Double(x *int) {
    *x *= 2
}

num := 5
Double(&num)    // num is now 10

// nil is null
var p *int      // nil
if p != nil {
    fmt.Println(*p)
}
```

## Concurrency (Goroutines & Channels)

| C# | JavaScript | Go |
|----|------------|-----|
| `Task.Run(() => ...)` | `setTimeout(...)` | `go func()` |
| `async/await` | `async/await` | channels or sync pkg |
| `Task<T>` | `Promise<T>` | `chan T` |

```go
// Start goroutine (lightweight thread)
go doSomething()

go func() {
    fmt.Println("In goroutine")
}()

// Channels (typed pipes for communication)
ch := make(chan int)        // unbuffered
ch := make(chan int, 10)    // buffered

// Send & receive
ch <- 42        // send
value := <-ch   // receive (blocks until data)

// Example: async operation
func fetchData() chan string {
    ch := make(chan string)
    go func() {
        // simulate work
        time.Sleep(time.Second)
        ch <- "data"
    }()
    return ch
}

result := <-fetchData()  // blocks until ready

// Select (like switch for channels)
select {
case msg := <-ch1:
    fmt.Println("From ch1:", msg)
case msg := <-ch2:
    fmt.Println("From ch2:", msg)
case <-time.After(time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No data ready")
}

// WaitGroup (like Task.WhenAll / Promise.all)
import "sync"

var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println(n)
    }(i)
}
wg.Wait()  // blocks until all done
```

## Generics

```go
// Generic function
func Map[T, U any](items []T, fn func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = fn(item)
    }
    return result
}

// Usage
doubled := Map([]int{1, 2, 3}, func(x int) int { return x * 2 })

// Generic type
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

// Constraints
type Number interface {
    int | int32 | int64 | float32 | float64
}

func Sum[T Number](nums []T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}
```

## Visibility (Public/Private)

| C# | Go |
|----|-----|
| `public` | `Uppercase` first letter |
| `private` | `lowercase` first letter |
| `internal` | lowercase (package-scoped) |
| `protected` | N/A (no inheritance) |

```go
type User struct {
    Name  string  // exported (public)
    email string  // unexported (private to package)
}

func ProcessUser() {}   // exported
func validateUser() {}  // unexported
```

## Common Imports

```go
import (
    "fmt"           // Printf, Println, Sprintf
    "strings"       // Join, Split, Contains, etc.
    "strconv"       // Atoi, Itoa, ParseInt
    "errors"        // New, Is, As
    "os"            // File I/O, env vars
    "io"            // Reader, Writer interfaces
    "time"          // Time, Duration, Sleep
    "sync"          // WaitGroup, Mutex
    "context"       // Context (cancellation/timeouts)
    "encoding/json" // Marshal, Unmarshal
    "net/http"      // HTTP client/server
)
```

## JSON

```go
import "encoding/json"

type User struct {
    Name  string `json:"name"`            // custom key
    Email string `json:"email,omitempty"` // omit if empty
    Age   int    `json:"-"`               // ignore field
}

// Serialize (like JsonSerializer.Serialize)
user := User{Name: "Alice", Email: "a@b.com"}
data, err := json.Marshal(user)
// {"name":"Alice","email":"a@b.com"}

// Deserialize (like JsonSerializer.Deserialize)
var user User
err := json.Unmarshal([]byte(jsonString), &user)

// Pretty print
data, _ := json.MarshalIndent(user, "", "  ")
```

## Quick Reference

| Concept | C# | JavaScript | Go |
|---------|----|-----------|----|
| Null | `null` | `null`/`undefined` | `nil` |
| Print | `Console.WriteLine()` | `console.log()` | `fmt.Println()` |
| Format | `$"x={x}"` | `` `x=${x}` `` | `fmt.Sprintf("x=%d", x)` |
| Lambda | `x => x * 2` | `x => x * 2` | `func(x int) int { return x * 2 }` |
| This | `this` | `this` | receiver (implicit) |
| New | `new T()` | `new T()` | `T{}` or `&T{}` |
| Typeof | `typeof(T)` | `typeof x` | `reflect.TypeOf(x)` |
| Cast | `(Type)x` | N/A | `x.(Type)` |
| Async | `async/await` | `async/await` | goroutines + channels |
