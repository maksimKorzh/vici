package main

import "os"
import "fmt"
import "github.com/nsf/termbox-go"

/* edit -- main routine for text editor */
func main() {
  if len(os.Args) > 1 {
    setbuf()
    savefile = os.Args[1]
    status := doread(0, savefile)
    if status == ERR { fmt.Println("?") }
  } else {
    setbuf()
    savefile = "out.txt"
  }
  err := termbox.Init()
  if err != nil { fmt.Println(err); os.Exit(1) }
  curln = 1
  for {
    doshow(true)
    if len(buf) > 1 {
      termbox.SetCursor(tabcl - offcl+lnwidth, curln - offrw-1)
    } else {
      termbox.SetCursor(tabcl - offcl+lnwidth, 0)
    }
    termbox.Flush()
    readkey()
  }
}
