package main

import "strings"
import "github.com/nsf/termbox-go"

/* pnum -- parse number */
func pnum(line string, i *int) string {
  num := ""
  for *i < len(line) && (line[*i] >= '0' && line[*i] <= '9') {
    num += string(line[*i])
    *i++
  }
  return num
}

/* pword -- parse word */
func pword(line string, i *int) string {
  word := ""
  for *i < len(line) && line[*i] != ' ' {
    if strings.Contains(":\"'+-*/<>=", string(line[*i])) { return word }
    word += string(line[*i])
    *i++
  }
  return word
}

/* pblank -- parse blank space */
func pblank(line string, i *int) string {
  bl := ""
  for *i < len(line) && line[*i] == ' ' {
    bl += string(line[*i])
    *i++
  }
  return bl
}

/* pstr -- parse string */
func pstr(line string, i *int) string {
  *i++
  str := "\""
  for *i < len(line) && line[*i] != '"' {
    str += string(line[*i])
    *i++
  }
  if *i < len(line) { str += string(line[*i]) }
  *i++
  return str
}

/* pchr -- parse chars */
func pchr(line string, i *int) string {
  *i++
  chr := "'"
  for *i < len(line) && line[*i] != '\'' {
    chr += string(line[*i])
    *i++
  }
  if *i < len(line) { chr += string(line[*i]) }
  *i++
  return chr
}

/* pcomm -- parse comments */
func pcomm(line string, i *int) string {
  comm := ""  
  for *i < len(line) {
    comm += string(line[*i])
    *i++
  }
  return comm
}


/* hlsyntax -- highlight syntax */
func hlsyntax(col, row int, line string) {
  i := 0
  if len(line) != len([]rune(line)) {
    msg(col, row, DCOL, DCOL, line)
    return
  }
  for i < len(line) {
    j := i
    if line[i] >= '0' && line[i] <= '9' {
      msg(col, row, YCOL|BOLD, DCOL, pnum(line, &i))
    } else if line[i] == ' ' {
      msg(col, row, DCOL, DCOL, pblank(line, &i))
    } else if line[i] == '"' {
      msg(col, row, YCOL, DCOL, pstr(line, &i))
    } else if line[i] == '\'' {
      msg(col, row, YCOL|BOLD, DCOL, pchr(line, &i))
    } else if line[i] == '/' || line[i] == '#' {
      msg(col, row, MCOL|BOLD, DCOL, string(line[i]))
      if line[i] == '#' {
        msg(col, row, MCOL|BOLD, DCOL, pcomm(line, &i))
      }
      i++
      if i < len(line) {
        if line[i] == '/' || line[i] == '*' {
          i--
          msg(col, row, MCOL|BOLD, DCOL, pcomm(line, &i))
        }
      } 
    } else if strings.Contains("+-*/=%<>:", string(line[i])) {
      msg(col, row, MCOL|BOLD, DCOL, string(line[i]))
      i++
    } else {
      tok := pword(line, &i)
      iskw := false
      for _, kw := range KEYWORDS {
        if tok == kw {
          msg(col, row, GCOL|BOLD, DCOL, tok)
          iskw = true
          break
        }
      }
      if iskw == false {
        msg(col, row, DCOL, DCOL, tok)
      }
    }
    col += i-j
  }
}

/* hlline -- highlight current line */
func hlline(row int) {
  cols, rows := termbox.Size(); rows--
  if row > rows { return }
  for col := 0; col < cols; col++ {
    if len(buf) > 1 && curln > 0 {
      cell := termbox.GetCell(col, row-1)
      if col == tabcl - offcl+lnwidth {
        termbox.SetCell(col, row-1, cell.Ch, WCOL, RCOL)
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
