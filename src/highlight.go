package main

import "strings"
import "github.com/nsf/termbox-go"

/* hlline -- highlight current line */
func hlline(col, row int) {
  if row == curln - offrw {
    cols, rows := termbox.Size(); rows--
    if row > rows { return }
    for col = 0; col < cols; col++ {
      if len(buf) > 1 {
        cell := termbox.GetCell(col, row-1)
        termbox.SetCell(col, row-1, cell.Ch, DCOL, BCOL)
      } else {
        msg(col, row, DCOL, BCOL, "1" + strings.Repeat(" ", cols-1))
        break
      }
    }
  }
}
