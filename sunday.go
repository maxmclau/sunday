package sunday

import (
  "github.com/codeskyblue/go-sh"
  "strconv"
  "fmt"
)

func clear() {
  sh.Command("tput", "clear").Run()
}

func move(x int, y int) {
  sh.Command("tput", "cup",  strconv.Itoa(y), strconv.Itoa(x)).Run()
}

func Header(x int, y int) {
  clear();

  move(x, y);

  fmt.Println("Hello")
}
