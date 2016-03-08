package sunday

import (
  "fmt"
  "time"
  "github.com/maxmclau/gput"
)

type Shape struct {
  X int
  Y int
  Width int
  Height int
  Fill bool
  FillColor int
  Pop bool
  PopColor int
}

func Begin() {
  gput.Clear()
}

func Exit() {
  gput.Cup(0, (gput.Lines() - 2))
  gput.Sgr0()
  gput.Dim()
  fmt.Println(" ")
}

func (s Shape) Render(speed time.Duration) {
  if s.Pop {
    fill(s.X + 2, s.Y + 1, s.Width, s.Height, s.PopColor, speed)
  }

  if s.Fill {
    fill(s.X, s.Y, s.Width, s.Height, s.FillColor, speed)
  }
}

func fill(x int, y int, width int, height int, color int, speed time.Duration) {
  gput.Setab(color)

  for h := 0; h < height; h++ {
    gput.Cup(x, (y + h))

    for w := 0; w < width; w++ {
      fmt.Printf(" ")

      time.Sleep(speed * time.Millisecond)
    }
  }
}
