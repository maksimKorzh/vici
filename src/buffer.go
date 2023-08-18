package main

import "regexp"
import "strings"
import "strconv"
import "unicode/utf8"
import "github.com/nsf/termbox-go"

/* setbuf (in memory) -- initialize line storage buffer */
func setbuf() {
  buf = []buftype{buftype{txt: ""},}
  curln = 0;
  lastln = 0
}

/* rw -- get current rune width */
func rw(i int) int {
  runes := []rune(buf[curln].txt + " ")
  return utf8.RuneLen(runes[i])
}

/* lnlen -- get length of current buffer row */
func lnlen() int {
  runes := []rune(buf[curln].txt)
  return len(runes)
}

/* cloff -- convert curcl to offset index */
func cloff(i int, s string) int {
  index := 0; offset := 0
  for offset = range s + " " {
    if index == i { return offset }
    index++
  }
  return offset
}

/* inrune -- insert char into line */
func inrune(c rune) {
  dirty = true
  offset := cloff(curcl, buf[curln].txt)
  if c != '\n' {
    lline := buf[curln].txt[:offset]
    rline := buf[curln].txt[offset:]
    buf[curln].txt = lline + string(c) + rline
    curcl = nextcl(curcl)
  } else {
    lline := buf[curln].txt[offset:]
    buf[curln].txt = buf[curln].txt[:offset]
    puttxt(lline)
    curcl = 0
  }
}

/* dlrune -- delete char in line */
func dlrune() {
  dirty = true
  offset := cloff(curcl, buf[curln].txt)
  if curcl > 0 {
    chlen := rw(curcl-1)
    lline := buf[curln].txt[:offset-chlen]
    rline := buf[curln].txt[offset:]
    buf[curln].txt = lline + rline
    curcl = prevcl(curcl)
  } else if curln > 1 {
    lline := buf[curln-1].txt
    rline := buf[curln].txt[offset:]
    stat := OK
    lndelete(curln, curln, &stat)
    curcl = lnlen()
    buf[curln].txt = lline + rline
  }
}

/* lnjoin -- join current line to previous */
func lnjoin() {
  for curcl > 0 { dlrune() }
  if curln > 1 { dlrune() }
}

/* lndelete -- delete lines n1 through n2 */
func lndelete(n1, n2 int, status *stcode) stcode {
  if n1 <= 0 {
    *status = ERR
  } else {
    newbuf := make([]buftype, len(buf) - (n2 - n1)-1)
    copy(newbuf[:n1], buf[:n1])
    copy(newbuf[n1:], buf[n2+1:])
    buf = newbuf
    lastln = lastln - (n2 - n1 + 1)
    if lastln == 0 {
      setbuf()
      puttxt("")
    } else {
      curln = prevln(n1)
    }
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
  newbuf[curln] = buftype{ txt:lin }
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
    cpb = append(cpb, buftype{ txt:buf[i].txt })
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

/* backup -- backup current buffer state to undo buffer */
func backup() {
  ubf = make([]buftype, len(buf))
  copy(ubf, buf)
  ox = curcl
  oy = curln
  ol = lastln
}

/* swapbf -- swap buffer and undo buffer */
func swapbf() {
  tbf := make([]buftype, len(buf))
  copy(tbf, buf)
  tx := curcl
  ty := curln
  tl := lastln
  buf = ubf
  curcl = ox
  curln = oy
  lastln = ol
  ubf = make([]buftype, len(tbf))
  copy(ubf, tbf)
  ox = tx
  oy = ty
  ol = tl
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

/* nextcl -- get next col */
func nextcl(n int) int {
  if curcl < lnlen() {
    return n + 1
  } else {
    if curln < lastln {
      curln = nextln(curln)
      return 0
    } else {
      return lnlen()
    }
  }
}

/* prevcl -- get prev col */
func prevcl(n int) int {
  if curcl > 0 {
    return n - 1
  } else {
    if curln > 1 {
      curln = prevln(curln)
      return lnlen()
    } else {
      return 0
    }
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

/* sptab -- convert tabs to spaces */
func sptab() []buftype {
  dbuf := make([]buftype, len(buf))
  copy(dbuf, buf)
/*
  int idx = 0;
  for (j = 0; j < row->size; j++) {
    if (row->chars[j] == '\t') {
      row->render[idx++] = ' ';
      while (idx % KILO_TAB_STOP != 0) row->render[idx++] = ' ';
    } else {
      row->render[idx++] = row->chars[j];
    }
  }
*/


  for row := 1; row < rows; row++ {
    brow := row + offrw
    if brow >= 1 && brow < len(buf) {
      i := 0; newstr := ""
      for _, ch := range buf[brow].txt {
        if ch == '\t' {
          newstr += " "
          i++
          for i % TABS != 0 {
            newstr += " "
            i++
          }
        } else {
          newstr += string(ch)
          i++
        }
      }
      dbuf[brow].txt = newstr
    }
  }
  return dbuf
}

/* cltab -- convert curcl to tabcl */
func cltab() int {
  rx := 0;
  for col := 0; col < curcl; col++ {
    if col < lnlen() {
      if buf[curln].txt[cloff(col, buf[curln].txt)] == '\t' { rx = rx + (TABS-1) - (rx % TABS) }
      rx++;
    }
  }
  return rx
}

/* doscroll -- scroll buffer based on offrw and offcl (visual mode) */
func doscroll() {
  if curln < 1 { return }
  tabcl = 0
  if curln <= lastln { tabcl = cltab() }
  if curln < offrw+1 { offrw = curln-1 }
  if tabcl < offcl { offcl = tabcl }
  if curln >= offrw + rows+1 { offrw = curln-rows }
  if tabcl >= offcl + cols-lnwidth { offcl = tabcl-cols+lnwidth+1 }
}

/* dorender -- display buffer content (visual mode) */
func dorender() {
  dbuf := sptab()
  for row := 1; row <= rows; row++ {
    brow := row + offrw
    if brow >= 1 && brow < len(buf) {
      lnnum := strconv.Itoa(brow)
      lnoff := lnwidth - len(lnnum)-1
      msg(lnoff, row-1, CCOL, DCOL, lnnum)
      if offcl >= len(dbuf[brow].txt) { continue }
      line := ""
      line = dbuf[brow].txt[cloff(offcl, dbuf[brow].txt):]
      //if offcl > 12 { dbdie(dbuf[brow].txt, "\n", line, offcl, cloff(offcl, dbuf[brow].txt)) }
      if hl == 1 {
        hlsyntax(lnwidth, row-1, line)
      } else {
        msg(lnwidth, row-1, DCOL, DCOL, line)
      } 
    } else if row-1 != 0 {
      msg(0, row-1, BCOL, DCOL, "*")
    }
  }
  hlline(curln-offrw)
}

/* dostat -- display status bar */
func dostat() {
  var modstat string
  fnlen := len(savefile)
  if fnlen > 24 { fnlen = 24 }
  flstat := savefile[:fnlen] + " - " + strconv.Itoa(lastln) + " lines"
  if dirty { flstat += " modified " } else { flstat += " saved" }
  if mode == EDIT { modstat = " EDIT: " } else { modstat = " VIEW: " }
  curstat := " Row " + strconv.Itoa(curln) + ", Col " + strconv.Itoa(tabcl+1) + " "
  uspace := len(modstat) + len(flstat) + len(curstat)
  spaces := ""
  if cols - uspace >= 1 { spaces = strings.Repeat(" ", cols - uspace) }
  message := modstat + flstat + spaces + curstat
  msg(0, rows, NCOL, WCOL, message)
}

/* doshow -- update display */
func doshow(resize bool) {
  if resize {
    cols, rows = termbox.Size()
    if cols < 10 && rows < 3 { return }
    if rows > 2 { rows-- }
  }
  lnwidth = len(strconv.Itoa(len(buf)-1))+1
  termbox.Clear(DCOL, DCOL)
  doscroll()
  dorender()
  dostat()
}
