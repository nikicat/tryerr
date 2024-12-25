# TryErr

TryErr is a tiny Go module designed to simplify error handling in Go. It provides utilities and helper functions to streamline error management in your codebase by wrapping and handling errors elegantly.

## Features

- Simplified error wrapping and handling
- Utility functions for custom error messages
- Supports returning multiple values with error handling

## Installation

You can install TryErr using `go get`:

```bash
go get github.com/nikicat/tryerr
```

## Usage

### Example 1: Basic Error Handling with `Try`

```go
package main

import (
	"fmt"
	"github.com/nikicat/tryerr"
)

func main() {
	result, err := functionWithSubcalls()
	if err != nil {
		fmt.Println("error: ", err)
		retudrn
	}
	fmt.Println(result)
}

func functionWithSubcalls() (result int, err error) {
	defer tryerr.Catch(&err)

	result = tryerr.Try(divide(100, 4)).Err("failed to divide by 4")
	result, remainder = tryerr.Try2(divrem(result, 3)).Err("failed to divide by 3")
	result = tryerr.Try(safe_sum(result, remainder)).Err("failed to sum")
	return
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func divrem(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}
	return a / b, a % b, nil
}

func safe_sum(a, b int) (int, error) {
	res := a + b
	if sign(a) == sign(b) && sign(r) != sign(a) {
		return 0, fmt.Errorf("integer overflow")
	}
	return r, nil
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}
```

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, feel free to open an issue or contact the repository owner.

Feel free to further customize this README according to your project's specific details and needs.
