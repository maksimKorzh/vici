package main

import "os"
import "fmt"

/* edit -- main routine for text editor */
func main() {
  if len(os.Args) > 1 {
    setbuf()
    savefile = os.Args[1]
    status := doread(0, savefile)
    if status == ERR { fmt.Println("?") }
  }
  var cursave, i int
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
  }
}
