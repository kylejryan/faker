   _____       ______    _             
  / ____|     |  ____|  | |            
 | |  __  ___ | |__ __ _| | _____ _ __ 
 | | |_ |/ _ \|  __/ _` | |/ / _ \ '__|
 | |__| | (_) | | | (_| |   <  __/ |   
  \_____|\___/|_|  \__,_|_|\_\___|_|   
                                       
                                       
# Table of Contents
- Compatibility
- Installation
- Basic Usage
- Features
- Contributing
- License

## Compatibility
`faker` requires Go 1.20 or higher. Please ensure your environment meets this requirement to use the package without issues.

## Installation
Install the package using go get:

```bash
go get github.com/kylejryan/faker
```

## Basic Usage
Import the package into your Go project and start generating fake names:

```go
package main

import (
    "fmt"

    "github.com/kylejryan/faker"
)

func main() {
    firstName := faker.GenerateFirstName()
    lastName := faker.GenerateLastName()
    fullName := faker.GenerateFullName()

    fmt.Println("First Name:", firstName)
    fmt.Println("Last Name:", lastName)
    fmt.Println("Full Name:", fullName)
}
```

### Sample Output:

```mathematica
First Name: Emily
Last Name: Johnson
Full Name: Michael Smith
```

Each function call generates a random name:

```go
for i := 0; i < 5; i++ {
    fmt.Println(faker.GenerateFullName())
}
```

### Possible Output:

```mathematica
Olivia Brown
Liam Davis
Sophia Miller
Noah Wilson
Isabella Moore
```

## Features
- **Simple API**: Easy-to-use functions for generating first names, last names, and full names.
- **No External Dependencies**: Pure Go implementation without the need for external data files.
- **Embedded Data**: Name data is embedded within the package using Go's embed package.
- **High Performance**: Optimized for speed and efficiency.
- **Extensible**: Designed with future enhancements in mind, such as adding more data types.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request on GitHub. Before contributing, please read our [Contributing Guidelines](https://github.com/kylejryan/faker/CONTRIBUTING.md).

## License
This project is licensed under the MIT License.

