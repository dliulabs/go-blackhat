# Net Monitor

```
package main

import (
	"log"
	"os"
)

type Monitor struct {
	*log.Logger
}

func main() {
	m := &Monitor{Logger: log.New(os.Stdout, "monitor: ", 0)}
	m.Output(2, "hello")
}
```