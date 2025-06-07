package main

import "os"
import "fmt"
import "strconv"
import "github.com/nsf/termbox-go"
import "github.com/mattn/go-runewidth"

/* min -- get min number */
func min(num1, num2 int) int {
  if num1 < num2 { return num1 }
  return num2
}

/* ctoi -- convert string at sri] to integer, increment i */
func ctoi(s string, i *int) int {
  var n, sign int
  skipbl(s, i)
  if s[*i] == '-' {
    sign = -1
  } else {
    sign = 1
  }
  if s[*i] == '+' || s[*i] == '-' { *i++ }
  for *i < len(s) && s[*i] >= '0' && s[*i] <= '9' {
    n = n * 10 + int(s[*i]) - '0'
    *i++
  }
  return sign * n
}

/* dbdie -- print debug info and die */
func dbdie(args ...interface{}) {
  termbox.Close()
  fmt.Println(args...)
  os.Exit(0)
}

/* getst -- get status code description */
func getst(status stcode) string {
  descr := ""
  switch status {
    case ENDDATA: descr = "ENDDATA"
    case ERR: descr = "ERR"
    case OK: descr = "OK"
  }
  return descr
}

/* itos - convert integer to string if possible */
func itos(v interface{}) string {
  switch val := v.(type) {
  case int:
    return strconv.Itoa(val)
  default:
    return fmt.Sprintf("%v", val) // fallback, generic
  }
}

/* msg -- print message on screen (visual mode) */
func msg(x, y int, fg, bg termbox.Attribute, msg string) {
  for _, c := range msg {
    if c == '\t' {
      for i := 0; i < TABS; i++ {
        termbox.SetCell(x, y, ' ', fg, bg)
        x += runewidth.RuneWidth(' ')
      }
    } else {
      termbox.SetCell(x, y, c, fg, bg)
      x += runewidth.RuneWidth(c)
    }
  }
}

/* debug -- print debug message (visual mode) */
func debug(message interface{}) {
  msg(0, rows-1, COL3, COL2, itos(message))
  termbox.Flush()
  getev()
}
