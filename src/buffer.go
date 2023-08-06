package main

import "fmt"
import "regexp"
import "strings"
import "strconv"

/* setbuf (in memory) -- initialize line storage buffer */
func setbuf() {
  buf = []buftype{buftype{txt: "", mark: false},}
  curln = 0;
  lastln = 0
}

/* doprint -- print lines n1 through n2 */
func doprint (n1, n2 int, c rune) stcode {
  if (n1 <= 0) {
    return ERR
  } else {
    for i := n1; i <= n2; i++ {
      if c == NCMD { fmt.Print(i, "\t") }
      fmt.Println(buf[i].txt)
    }
    curln = n2;
    return OK
  }
}

/* lnappend -- append lines after "line" */
func lnappend(line int) stcode {
  var inline string
  var stat stcode
  var done bool
  curln = line
  stat = OK
  done = false
  for done == false && stat == OK {
    inline = getline()
    if inline[0] == PERIOD && inline[1] == NEWLINE {
      done = true
    } else if puttxt(inline[:len(inline)-1]) == ERR {
      stat = ERR
    }
  }
  return stat
}

/* lndelete -- delete lines n1 through n2 */
func lndelete(n1, n2 int, status *stcode) stcode {
  if n1 <= 0 {
    *status = ERR
  } else {
    newbuf := make([]buftype, len(buf) - (n2 - n1))
    copy(newbuf[:n1], buf[:n1])
    copy(newbuf[n1:], buf[n2+1:])
    buf = newbuf
    lastln = lastln - (n2 - n1 + 1)
    curln = prevln(n1)
    *status = OK
  }
  return *status
}

/* puttxt (to buffer) -- put text from lin after curln */
func puttxt (lin string) stcode {
  lastln++
  curln++
  newbuf := make([]buftype, lastln+1)
  copy(newbuf[:curln], buf[:curln])
  newbuf[curln] = buftype{ txt: lin, mark: false }
  copy(newbuf[curln+1:], buf[curln:])
  buf = newbuf
  return OK
}

/* move -- move line1 through line2 after line3 */
func move(line3 *int) stcode {
  if line1 <= 0 || (*line3 >= line1 && *line3 < line2) {
    return ERR
  } else {
    blkmove(line1, line2, *line3);
    if (*line3 > line1) {
      curln = *line3
    } else {
      curln = *line3 + (line2 - line1 + 1)
    }
    return OK
  }
}

/* cp -- copy lines n1,n2 to copy buffer */
func cp(n1, n2 int) {
  cpb = []buftype{}
  for i := n1; i <= n2; i++ {
    cpb = append(cpb, buftype{txt: buf[i].txt, mark: buf[i].mark })
  }
}

/* dup -- duplicate line1 through line2 after line3 */
func dup(line3 *int) stcode {
  if line1 <= 0 || (*line3 >= line1 && *line3 < line2) {
    return ERR
  } else {
    blkcopy(line1, line2, *line3);
    if (*line3 > line1) {
      curln = *line3
    } else {
      curln = *line3 + (line2 - line1 + 1)
    }
    return OK
  }
}

/* blkmove -- move block of lines n1..n2 to after n3 */
func blkmove(n1, n2, n3 int) {
  if n3 < n1-1 {
    reverse(n3+1, n1-1)
    reverse(n1, n2)
    reverse(n3+1, n2)
  } else if n3 > n2 {
    reverse(n1, n2)
    reverse(n2+1, n3)
    reverse(n1, n3)
  }
}

/* blkcopy -- copy block of lines n1..n2 to after n3 */
func blkcopy(n1, n2, n3 int) {
  cp(n1,n2)
  curln = n3
  for i := n1; i <= n2; i++ {
    puttxt(cpb[i-n1].txt)
  }
}

/* reverse -- reverse buf[n1]...buf[n2] */
func reverse(n1, n2 int) {
var temp buftype
  for n1 < n2 {
    temp = buf[n1]
    buf[n1] = buf[n2]
    buf[n2] = temp
    n1 = n1 + 1
    n2 = n2 - 1
  }
}

/* nextln -- get line after n */
func nextln(n int) int {
  if n >= lastln {
    return 0
  } else {
    return n + 1
  }
}

/* prevln -- get line before n */
func prevln(n int) int {
  if n <= 0 {
    return lastln
  } else {
    return n - 1
  }
}

/* subst -- substitute pattern with "sub" */
func subst(sub string) stcode {
  var line string
  var stat stcode
  stat = ERR
  for i := line1; i <= line2; i++ {
    line = buf[i].txt
    r, err := regexp.Compile(pat)
    if err != nil {
      stat = ERR
    } else {
      if r.MatchString(line) {
        curln = i
        buf[i].txt = r.ReplaceAllString(line, sub)
      }
    }
  }
  stat = OK
  return stat
}

/* doscroll -- scroll buffer based on offrw and offcl (visual mode) */
func doscroll() {
  if curln < offrw { offrw = curln }
  if curcl < offcl { offcl = curcl }
  if curln >= offrw + rows { offrw = curln-rows+1 }
  if curcl >= offcl + cols-lnwidth { offcl = curcl-cols+lnwidth+1 }
}

/* dorender -- display buffer content (visual mode) */
func dorender() {
  for row := 0; row < rows; row++ {
    brow := row + offrw
    if brow >= 1 && brow < len(buf) {
      lnnum := strconv.Itoa(brow)
      lnoff := lnwidth - len(lnnum)-1
      msg(lnoff, row-1, CCOL, DCOL, lnnum)
      line := buf[brow].txt[offcl:]
      msg(curcl-offcl+lnwidth, row-1, DCOL, DCOL, line)
    }
    hlline(0, row)
    dostat()
  }
}

/* dostat -- display status bar */
func dostat() {
  var modstat string
  fnlen := len(savefile)
  if fnlen > 24 { fnlen = 24 }
  flstat := savefile[:fnlen] + " - " + strconv.Itoa(len(buf)) + " lines"
  if dirty { flstat += " modified " } else { flstat += " saved" }
  if mode == EDIT { modstat = " EDIT: " } else { modstat = " VIEW: " }
  curstat := " Row " + strconv.Itoa(curln+1) + ", Col " + strconv.Itoa(curcl+1) + " "
  uspace := len(modstat) + len(flstat) + len(curstat)
  spaces := strings.Repeat(" ", cols - uspace)
  message := modstat + flstat + spaces + curstat
  msg(0, rows, NCOL, WCOL, message)
}
