package main

import "os"
import "regexp"
import "strings"
import "github.com/nsf/termbox-go"

/* getlist -- get list of line nums at lin[i], increment i */
func getlist(lin string, i *int, status *stcode) stcode {
  var num int
  var done bool
  line2 = 0;
  nlines = 0;
  done = getone(lin, i, &num, status) != OK
  for *i < len(lin) && done == false {
    line1 = line2
    line2 = num
    nlines++
    if lin[*i] == SEMICOL { curln = num }
    if lin[*i] == COMMA || lin[*i] == SEMICOL {
      *i++
      done = getone(lin, i, &num, status) != OK
    } else { done = true }
  }
  nlines = min(nlines, 2);
  if nlines == 0 { line2 = curln }
  if nlines <= 1 { line1 = line2 }
  if (*status != ERR) { *status = OK }
  return *status
}

/* getone -- get one line number expression */
func getone(lin string, i, num *int, status *stcode) stcode {
  var istart, mul, pnum int
  istart = *i;
  *num = 0;
  if getnum(lin, i, num, status) == OK {    /*  1st term */
    for *i < len(lin) {
      skipbl(lin, i)
      if *i >= len(lin) { *status = ENDDATA; break }
      if lin[*i] != PLUS && lin[*i] != MINUS {
        *status = ENDDATA
      } else {
        if *i < len(lin) && lin[*i] == PLUS {
          mul = +1
        } else {
          mul = -1
        }; *i++
        if getnum(lin, i, &pnum, status) == OK {
          *num = *num + mul * pnum
        }
        if *status == ENDDATA { *status = ERR }
      }
      if *status != OK { break }
    }
  }
  if *num < 0 || *num > lastln { *status = ERR }
  if *status != ERR {
    if *i <= istart {
      *status = ENDDATA
    } else {
      *status = OK
    }
  }
  return *status
}

/* getnum -- get single line number component */
func getnum(lin string, i, num *int, status *stcode) stcode {
  *status = OK
  skipbl(lin, i)
  if *i < len(lin) {
    if lin[*i] >= '0' && lin[*i] <= '9' {
        *num = ctoi(lin, i)
        *i--    /* move back; to be advanced at end */
    } else if lin[*i] == CURLINE {
      *num = curln
    } else if lin[*i] == LASTLINE {
      *num = lastln
    } else if lin[*i] == SCAN || lin[*i] == BACKSCAN {
      if (optpat(lin, i) == ERR) {
        *status = ERR  /* build pattern */
      } else {
        if *i < len(lin) {
          *status = patscan(rune(lin[*i]), num)
        } else {
          *status = ERR
        }
      }
    } else { *status = ENDDATA }
  }
  if (*status == OK) { *i++ }    /* next character to be examined */
  return *status
}

/* optpat -- get optional pattern from lin[i], increment i */
func optpat(lin string, i *int) stcode {
  if *i >= len(lin)-1 { return ERR }
  if lin[*i+1:] == "\n" {
    *i = 0
  } else if lin[*i+1] == lin[*i] {    /* repeated delimiter */
   *i++    /* leave existing pattern alone */
  } else {
    pat = strings.Split(lin[*i:], string(lin[*i]))[1]
    *i += len(pat)+1
    if pat == "" { *i = 0 }
    if *i == 0 { return ERR }
  }
  return OK
}

/* patscan -- find next occurrence of pattern after line n */
func patscan (way rune, n *int) stcode {
  var done bool
  var line string
  var stat stcode
  *n = curln
  stat = ERR
  done = false;
  for {
    if way == SCAN {
      *n = nextln(*n)
    } else {
      *n = prevln(*n)
    }
    line = buf[*n].txt
    r, err := regexp.Compile(pat)
    if err != nil {
      stat = ERR
    } else {
      if r.MatchString(line) {
        curcl = r.FindStringIndex(line)[0]
        stat = OK
        done = true
      }
    }
    if *n == curln || done { break }
  }
  return stat
}

/* setdef -- set defaulted line numbers */
func setdef(def1, def2 int, status *stcode) stcode {
  if (nlines == 0) {
    line1 = def1;
    line2 = def2
  }
  if line1 > line2 || line1 <= 0 {
    *status = ERR
  } else {
    *status = OK;
  }
  return *status
}

/* skipbl -- skip blanks and tabs at s[i]... */
func skipbl(s string, i *int) {
  if *i >= len(s) {return }
  for s[*i] == BLANK || s[*i] == TAB {
    *i ++
    if *i >= len(s) {return }
  }
}

/* getword -- get word from s[i] into out */
func getword(s string, i *int, out *string) int {
  skipbl(s, i)
  name := strings.Split(s[*i:], " ")[0]
  name = name[:len(name)-1]
  *out = name
  *i += len(*out)-1
  if len(s) == 0 {
    return 0
  } else {
    return *i
  }
}

/* ckp -- check for "p" after command */
func ckp (lin string, i *int,  pflag *bool, status *stcode) stcode {
  skipbl(lin, i);
  if lin[*i] == PCMD {
    *i++
    *pflag = true
  } else {
    *pflag = false
  }

  if lin[*i] == NEWLINE {
    *status = OK
  } else {
    *status = ERR
  }
  return *status
}

/* docmd -- handle all commands except globals */
func docmd (lin string, i *int, status *stcode) stcode {
  var fil, sub string
  var line3 int
  var pflag bool
  pflag = false;    /* may be set by d, m, s */
  *status = ERR;
  if lin[*i] == COMMA {
    line1 = 1
    line2 = lastln
    nlines = 1
    *i++
  }
  if lin[*i] == NEWLINE {
    if nlines == 0 { line2 = nextln(curln) }
    if line2 > 0 && line2 < len(buf) {
      curln = line2
      *status = OK
    }
  } else if lin[*i] == QCMD {
    if lin[*i+1] == NEWLINE && nlines == 0 {
      termbox.Close()
      os.Exit(0)
    }
  } else if lin[*i] == ACMD {
    if lin[*i+1] == NEWLINE {
      *status = lnappend(line2)
      if *status == OK { dirty = true }
    }
  } else if lin[*i] == CCMD {
    if lin[*i+1] == NEWLINE {
      if setdef(curln, curln, status) == OK {
        if lndelete(line1, line2, status) == OK {
          *status = lnappend(prevln(line1))
          if *status == OK { dirty = true }
        }
      }
    }
  } else if lin[*i] == DCMD {
    *i++
    if ckp(lin, i, &pflag, status) == OK {
      if setdef(curln, curln, status) == OK {
        cp(line1, line2)
        if lndelete(line1, line2, status) == OK {
          dirty = true
          if nextln(curln) != 0 {
            curln = nextln(curln)
          }
        }
      }
    }
  } else if lin[*i] == ICMD {
    if lin[*i+1] == NEWLINE {
      dirty = true
      if line2 == 0 {
        *status = lnappend(0)
      } else {
        *status = lnappend(prevln(line2))
      }
    }
  } else if lin[*i] == JCMD {
    lnjoin()
    *status = OK
  } else if lin[*i] == NCMD {
    inrune('\n')
    *status = OK
  } else if lin[*i] == MCMD {
    *i++
    if getone(lin, i, &line3, status) == ENDDATA { *status = ERR }
    if *status == OK {
      if ckp(lin, i, &pflag, status) == OK {
        if setdef(curln, curln, status) == OK {
          *status = move(&line3)
          if *status == OK { dirty = true }
        }
      }
    }
  } else if lin[*i] == TCMD {
    *i++
    if getone(lin, i, &line3, status) == ENDDATA { *status = ERR }
    if *status == OK {
      if ckp(lin, i, &pflag, status) == OK {
        if setdef(curln, curln, status) == OK {
          *status = dup(&line3)
          if *status == OK { dirty = true }
        }
      }
    }
  } else if lin[*i] == YCMD {
    *i++
    if ckp(lin, i, &pflag, status) == OK {
      if setdef(curln, curln, status) == OK {
        cp(line1, line2)
        *status = OK
      }
    }
  } else if lin[*i] == PCMD {
    *i++
    if ckp(lin, i, &pflag, status) == OK {
      curln = line2
      if len(cpb) > 0 {
        for i := 0; i < len(cpb); i++ {
          puttxt(cpb[i].txt)
        }
        dirty = true
        *status = OK
      } else {
        *status = ERR
      }
    }
  } else if lin[*i] == SCMD {
    *i++
    if optpat(lin, i) == OK {
      if *i < len(lin) {
        sub = strings.Split(lin[*i:], string(lin[*i]))[1]
        *i += len(sub)+2
        if *i < len(lin) {
          if ckp(lin, i, &pflag, status) == OK {
            if setdef(1, lastln, status) == OK {
              *status = subst(sub)
              if *status == OK { dirty = true }
            }
          }
        }
      }
    }
  } else if lin[*i] == ECMD {
    if nlines == 0 {
      if getfn(lin, i, &fil) == OK {
        savefile = fil
        setbuf()
        *status = doread(0, fil)
      }
    }
  } else if lin[*i] == FCMD {
    if nlines == 0 {
      if getfn(lin, i, &fil) == OK {
        savefile = fil
        *status = OK
      }
    }
  } else if lin[*i] == RCMD {
    if getfn(lin, i, &fil) == OK {
      *status = doread(line2, fil)
      if *status == OK { dirty = true }
    }
  } else if lin[*i] == WCMD {
    if getfn(lin, i, &fil) == OK {
      if setdef(1, lastln, status) == OK {
        *status = dowrite(line1, line2, fil)
      }
    } else {
      *i++
      if savefile != "" && lin[*i] == 'q' {
        dowrite(1, lastln, savefile)
        termbox.Close()
        os.Exit(0)
      }
    }
  }
  return *status
}

/* execcom -- execute command (visual mode) */
func execcom(com string) {
  var i int
  var status stcode
  lin = com + "\n"
  if getlist(lin, &i, &status) == OK {
    docmd(lin, &i, &status)
  }
}
