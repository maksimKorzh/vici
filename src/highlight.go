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
    if strings.Contains(":\"'+-*/<>=.,(){}[];", string(line[*i])) { return word }
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
    if line[*i] == '\\' && (*i+1) < len(line) && line[*i+1] == '"'{
      *i += 2
      str += "\\\""
      if (*i-3) >= 0 && line[*i-3] == '\\' {
        return str
      } else {
        continue
      }
    }
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
    if line[*i] == '\\' && (*i+1) < len(line) && line[*i+1] == '\'' {
      *i += 2
      chr += "\\'"
      if (*i-3) >= 0 && line[*i-3] == '\\' {
        return chr
      } else {
        continue
      }
    }
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
    msg(col, row, COL1, COL1, line)
    return
  }
  for i < len(line) {
    j := i
    if line[i] >= '0' && line[i] <= '9' {
      msg(col, row, COL7|BOLD, COL1, pnum(line, &i))
    } else if line[i] == ' ' {
      msg(col, row, COL1, COL1, pblank(line, &i))
    } else if line[i] == '"' {
      msg(col, row, COL8, COL1, pstr(line, &i))
    } else if line[i] == '\'' {
      msg(col, row, COL8|BOLD, COL1, pchr(line, &i))
    } else if line[i] == '/' || line[i] == '#' {
      msg(col, row, COL9|BOLD, COL1, string(line[i]))
      if line[i] == '#' {
        msg(col, row, COL9|BOLD, COL1, pcomm(line, &i))
      }
      i++
      if i < len(line) {
        if line[i] == '/' || line[i] == '*' {
          i--
          msg(col, row, COL9|BOLD, COL1, pcomm(line, &i))
        }
      }
    } else if strings.Contains(".,(){}[];", string(line[i])) {
      msg(col, row, COL1, COL1, string(line[i]))
      i++
    } else if strings.Contains("+-*/=%<>:", string(line[i])) {
      msg(col, row, COL1, COL1, string(line[i]))
      i++
    } else {
      tok := pword(line, &i)
      iskw := false
      for _, kw := range KEYWORDS {
        if tok == kw {
          msg(col, row, COL5|BOLD, COL1, tok)
          iskw = true
          break
        }
      }
      if iskw == false {
        msg(col, row, COL1, COL1, tok)
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
        termbox.SetCell(col, row-1, cell.Ch, COL3, COL4)
      } else {
        termbox.SetCell(col, row-1, cell.Ch, COL1, COL4)
      }
    } else {
      lnoff := lnwidth - 2
      msg(col, row, COL1, COL4, strings.Repeat(" ", lnoff) + "1" + strings.Repeat(" ", cols-1))
      break
    }
  }
}
