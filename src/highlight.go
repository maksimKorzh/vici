package main

import "github.com/nsf/termbox-go"

/* hlline -- highlight current line */
func hlline(col, row int) {
  if row == curln - offrw {
    cols, rows := termbox.Size(); rows--
    if row >= rows { return }
    for col = 0; col < cols; col++ {
      cell := termbox.GetCell(col, row)
      termbox.SetCell(col, row, cell.Ch, DCOL, BCOL)
    }
  }
}
