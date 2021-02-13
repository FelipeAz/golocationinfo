# golocationinfo
golocationinfo is a simple program that I've developed to practice the content I learned on a course at Udemy. This application returns the information using the POST CODE. If you enter a valid CODE, the API returns the City, State, Street and Neighborhood that it belongs.

# Installation
Use `go tool` to install this package on your go Project.

`go get -u github.com/FelipeAz/golocationinfo`

# Usage
You must have this package installed on your project. You can do that by using the line bellow.

```
import (
	"fmt"

	fibonacci "github.com/FelipeAz/golocationinfo"
)
```

# Example
```
package main

import locationhelper "github.com/FelipeAz/golocationinfo"

func main() {
	locationhelper.GetLocationInfo("yourCep")
}
```

## License
[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat)](LICENSE)
