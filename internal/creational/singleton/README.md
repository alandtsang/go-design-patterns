# Singleton

`Singleton` is a creational design pattern that lets you ensure that a class has only one instance, while providing a
global access point to this instance.

```go
var (
    once           sync.Once
    singleInstance *Single
)

// Single is a singleton structure.
type Single struct {
}

// New creates the single instance and return it when the first time the function is called.
// The second time the function is called, the single instance is returned.
func New() *Single {
    once.Do(func() {
        fmt.Println("Creating single instance")
        singleInstance = &Single{}
    })

    return singleInstance
}
```

## Usage

```go
instance := singleton.New()
```