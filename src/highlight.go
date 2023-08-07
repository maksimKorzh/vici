package main

import "strings"
import "github.com/nsf/termbox-go"

/* hlline -- highlight current line */
func hlline(row int) {
  cols, rows := termbox.Size(); rows--
  if row > rows { return }
  for col := 0; col < cols; col++ {
    if len(buf) > 1 && curln > 0 {
      cell := termbox.GetCell(col, row-1)
      if col == tabcl - offcl+lnwidth {
        termbox.SetCell(col, row-1, cell.Ch, DCOL, RCOL)
      } else {
        termbox.SetCell(col, row-1, cell.Ch, DCOL, BCOL)
      }
    } else {
      lnoff := lnwidth - 2
      msg(col, row, DCOL, BCOL, strings.Repeat(" ", lnoff) + "1" + strings.Repeat(" ", cols-1))
      break
    }
  }
}
