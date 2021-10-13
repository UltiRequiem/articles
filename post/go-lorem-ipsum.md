+++
title = "Generate Lorem Ipsum on Go"
date = "2021-10-12"
tags = ["go"]
draft = false
author = "UltiRequiem"
+++

Lorem ipsum is a placeholder text commonly used to demonstrate the visual form
of a document or a typeface without relying on meaningful content.

Here is a snippet to generate Lorem Ipsum quickly and easily.

```go
import (
	"fmt"
	"math/rand"
	"time"
)

var DATA [30]string = [30]string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "Mauris", "faucibus", "lectus", "eget", "cursus", "tempus", "ligula", "orci", "mattis", "massa", "nec", "eleifend", "lorem", "ipsum", "congue", "erat", "Pellentesque", "suscipit", "semper", "sapien", "sed", "luctus"}

func LoremWords(quantity int) string {
	rand.Seed(time.Now().UnixNano())

	lorem := ""

	for i := 0; i < quantity; i++ {
		lorem += DATA[rand.Intn(30)] + " "
	}

	return lorem
}
```

`LoremWords` will return something like: `"congue amet ligula nec lorem"`.

If you want it to look friendlier:

```go
func Sentence(words int) string {
	text := LoremWords(words)
	return strings.TrimSpace(strings.ToUpper(text[0:1])+text[1:]) + "."
}
```

`Sentence` will return something like: `"Consectetur eget tempus ipsum sapien."`

If you prefer an external package to take care of this,
you can use [Lorelai](https://github.com/UltiRequiem/lorelai).

```go
package main

import (
	"fmt"
	"github.com/UltiRequiem/lorelai/pkg"
)

func main() {
        fmt.Println(lorelai.Paragraph())
        fmt.Println(lorelai.Sentence())
        fmt.Println(lorelai.LoremWords(40))
}
```

It also brings some extra utilities:

```go
func sayHi() {
	fmt.Printf("My, my name is: %s.\n", lorelai.Word())
	fmt.Println("My email address is %s\n", lorelai.Email())
	fmt.Println("My website is: %s\n", lorelai.URL())
}
```

You can see Lorelai's docs [here](https://github.com/UltiRequiem/lorelai#documentation).
