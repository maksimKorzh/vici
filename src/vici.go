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
      termbox.SetCursor(curcl - offcl+lnwidth, curln - offrw-1)
    } else {
      termbox.SetCursor(curcl - offcl+lnwidth, 0)
    }
    termbox.Flush()
    readkey()
  }

/*  var cursave, i int
  var status stcode
  for {
    lin = getline()
    i = 0;
    cursave = curln;
    if getlist(lin, &i, &status) == OK {
      status = docmd(lin, &i, &status)
    }
    if status == ERR {
      fmt.Println("?")
      curln = min(cursave, lastln)
    } else if status == ENDDATA {
      break
    }
  }*/
}
