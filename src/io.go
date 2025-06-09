package main

import "os"
import "bufio"
import "strings"
import "github.com/nsf/termbox-go"

/* doread -- read "fil" after line n */
func doread(n int, fil string) stcode {
  file, err := os.Open(fil)
  if err != nil { return ERR }
  defer file.Close()
  curln = n
  scanner := bufio.NewScanner(file)
  buf := make([]byte, 0, 64*1024)
  scanner.Buffer(buf, 1024*1024)
  count := 0
  for scanner.Scan() {
    line := scanner.Text()
    puttxt(line)
    count += len(line) + 1
  }
  if count == 0 { puttxt("") }
  dirty = false
  return OK
}

/* dowrite -- write lines n1..n2 into file */
func dowrite(n1, n2 int, fil string) stcode {
  file, err := os.Create(fil)
  if err != nil { return ERR }
  defer file.Close()
  writer := bufio.NewWriter(file)
  count := 0
  for line := n1; line <= n2; line++ {
    newline := "\n"
    if line == len(buf) { newline = "" }
    writeln := buf[line].txt + newline
    _, err = writer.WriteString(writeln)
    if err != nil { return ERR }
    count += len(writeln)
  }
  writer.Flush()
  dirty = false
  return OK
}

/* getfn -- get file name from lin[i]... */
func getfn(lin string, i *int, fil *string) stcode {
  var k int
  var stat stcode
  stat = ERR
  if lin[*i+1] == BLANK {
    *i += 2
    k = getword(lin, i, fil)    /* get new filename */
    if (k > 0) { stat = OK }
  } else if lin[*i+1] == NEWLINE && savefile != "" {
    *fil = savefile
    stat = OK
  }
  if stat == OK && savefile == "" {
    savefile = *fil    /* save if no old one */
  }
  return stat
}

/* getev -- get termbox event (visual mode) */
func getev() termbox.Event {
  var ev termbox.Event
  switch event := termbox.PollEvent(); event.Type {
     case termbox.EventKey: ev = event
     case termbox.EventError: panic(event.Err)
   }
   return ev
}

/* readkey -- process keyboard input (visual mode) */
func readkey() {
  ev := getev()
  if ev.Key == termbox.KeyCtrlU {
    execcom(SCRDN)
  } else if ev.Key == termbox.KeyCtrlD {
    execcom(SCRUP)
  } else if ev.Key == termbox.KeyCtrlR {
    execcom("U")
  } else if ev.Key == termbox.KeyEsc {
    mode = VIEW
  } else if ev.Ch != 0 {
    if mode == EDIT {
      inrune(rune(ev.Ch))
    } else if mode == REPLACE {
      rerune(rune(ev.Ch))
    } else {
      switch ev.Ch {
        case ':': cprompt()
        case '/': cprompt()
        case 'i': execcom("i")
        case 'r': execcom("R"); dostat(); termbox.Flush(); readkey(); curcl -= 1; mode = VIEW
        case 'R': execcom("R")
        case 'H': execcom("H")
        case 'A': execcom("A")
        case 'g': execcom("1")
        case 'G': execcom("$")
        case 'y': execcom("y")
        case 'J': execcom("J")
        case 'x': execcom("x")
        case 'p': execcom("p")
        case 'd': execcom("d")
        case 'u': execcom("u")
        case 'n': execcom("//")
        case 'N': execcom("\\\\")
        case 'h': execcom("h")
        case 'j': execcom("j")
        case 'k': execcom("k")
        case 'l': execcom("l")
        case '0': execcom("<")
        case '$': execcom(">")
        case 'b': execcom("\\[a-zA-Z0-9_]+\\")
        case 'e': execcom("/\\w\\W/")
        case 'w': if execcom("/\\W\\w/") == OK { curcl = nextcl(curcl) }
      }
    }
  } else {
    switch ev.Key {
      case termbox.KeyArrowUp: execcom("k")
      case termbox.KeyArrowDown: execcom("j")
      case termbox.KeyArrowLeft: execcom("h")
      case termbox.KeyArrowRight: execcom("l")
      case termbox.KeyHome: execcom("<")
      case termbox.KeyEnd: execcom(">")
      case termbox.KeyPgup: execcom(SCRDN)
      case termbox.KeyPgdn: execcom(SCRUP)
    }
    if mode == EDIT {
      switch ev.Key {
        case termbox.KeySpace: inrune(' ')
        case termbox.KeyTab: for i := 0; i < TABS; i++ { inrune(' ') }
        case termbox.KeyBackspace: dlrune()
        case termbox.KeyBackspace2: dlrune()
        case termbox.KeyEnter:
          line := buf[curln].txt
          line = line[:len(line)-len(strings.TrimLeft(line, " "))]
          inrune('\n')
          for i := 0; i < len(line); i++ { inrune(' ') }
      }
    } else if mode == REPLACE {
      if ev.Key == termbox.KeySpace { rerune(' ') }
    }
  }
  if curcl > lnlen() { curcl = lnlen() }
}

/* getline -- invoke prompt to execute commands */
func getline(prompt string) string {
  doshow(false)
  msg(0, rows+1, COL1, COL1, prompt)
  termbox.SetCursor(1, rows+1)
  termbox.Flush()
  command := ""
  for {
    ev := getev()
    msg(0, rows+1, COL1, COL1, prompt)
    switch ev.Key {
      case termbox.KeyArrowUp: if curln > 1 { curln = prevln(curln) }; return ""
      case termbox.KeyArrowDown: if curln < lastln { curln = nextln(curln) }; return ""
      case termbox.KeyEsc: return ""
      case termbox.KeyEnter: return command + "\n"
      case termbox.KeySpace: command += " "
      case termbox.KeyTab: command += "\t"
      case termbox.KeyBackspace: if len(command) > 0 { command = command[:len(command)-1] }
      case termbox.KeyBackspace2: if len(command) > 0 { command = command[:len(command)-1] }
    }
    if ev.Ch != 0 {
      command += string(ev.Ch)
      msg(1, rows+1, COL1, COL1, strings.Replace(command, "\t", " ", -1))
    };
    cmdlen := 0
    for _,ch := range command { if ch > 0 { cmdlen++} }
    termbox.SetCursor(cmdlen+1, rows+1)
    for i := len(command)+1; i < cols; i++ {
      termbox.SetChar(i, rows+1, rune(' '))
    }
    termbox.Flush()
  }
}

/* cprompt -- invoke prompt to execute commands */
func cprompt() {
  rows--
  lin = getline(CPRMT)
  if lin == "" || lin == "\n" { return }
  i := 0;
  cursave := curln;
  var status stcode
  if getlist(lin, &i, &status) == OK {
    status = docmd(lin, &i, &status)
  }
  if status == ERR {
    msg(0, rows+1, COL1, COL1, "?" + strings.Repeat(" ", cols-1))
    curln = min(cursave, lastln)
    termbox.SetCursor(1, rows+1)
    termbox.Flush()
  }
  rows++
}
