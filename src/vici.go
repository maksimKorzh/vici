package main

import "os"
import "fmt"
import "github.com/nsf/termbox-go"

/* edit -- main routine for text editor */
func main() {
  setbuf()
  if len(os.Args) > 1 {
    savefile = os.Args[1]
    doread(0, savefile)
  } else {
    savefile = "out.txt"
  }
  if lastln == 0 { puttxt("") }
  err := termbox.Init()
  if err != nil { fmt.Println(err); os.Exit(1) }
  curln = lastln
  dirty = false
  hl = 1
  autoindent = 1
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
