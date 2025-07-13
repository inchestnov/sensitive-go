# sensitive-go

A lightweight Go library for masking and restoring sensitive fields in structures.

# Features
* Detach fields with `sensitive:true` tag in structures
* Attach back sensitive fields
* Generic restriction for type-safety

# Installation
```shell
go get github.com/inchestnov/sensitive-go
```

# Example
Consider structure `User` with field `Password`. 
```go
type User struct {
	ID       int64
	Username string
	Password string
	Email    string
	// ...
}
```

With `sensitive-go` you can easily «detach» all sensitive fields and works use structure for non-safe operations, for example logging or storing in database.
All you need is add label `sensitive:true` at field:

```go
type User struct {
	ID       int64
	Username string
	Password string `sensitive:"true"`
	Email    string
	// ...
}
```

```go
u := User{
    ID:       1,
    Username: "inchestnov",
    Password: "dfc95ff00a9f9fef2840b5f7c2481d71",
    Email:    "inchestnov@gmail.com",
}

insensitive, _, err := sensitive.Detach(u)
if err != nil {
    // ...
}
password := insensitive.Password // ""
```

You can also use returned sensitive data for restoring sensitive information:
```go
insensitive, sensitiveData, err := sensitive.Detach(u)
if err != nil {
    // ...
}

restored, err := sensitive.Attach(detached, sensitiveData)
if err != nil {
    // ...
}

password := restored.Password // dfc95ff00a9f9fef2840b5f7c2481d71
```
