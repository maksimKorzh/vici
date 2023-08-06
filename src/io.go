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
  count := 0
  for scanner.Scan() {
    line := scanner.Text()
    puttxt(line)
    count += len(line) + 1
  }
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
  //fmt.Println(count)
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
  if ev.Key == termbox.KeyEsc {
    mode = VIEW
  } else if ev.Ch != 0 {
    if mode == EDIT {
      //insert_rune(ev)
      dirty = true
    } else {
      switch ev.Ch {
        case 'q': termbox.Close(); os.Exit(0)
        case ':': cprompt()
      }
    }
  } else {
    switch ev.Key {
      case termbox.KeyArrowUp: if curln > 1 { curln = prevln(curln) }
      case termbox.KeyArrowDown: if curln < lastln { curln = nextln(curln) }
      case termbox.KeyArrowLeft: curcl = prevcl(curcl)
      case termbox.KeyArrowRight: curcl = nextcl(curcl)
    }
    //if curcl > len(buf[curln].txt) {
    //  curcl = len(buf[curln].txt)
    //  tabcl = curcl
    //}
  }
}

/* getline -- invoke prompt to execute commands */
func getline(prompt string) string {
  doshow(false)
  msg(0, rows+1, DCOL, DCOL, prompt)
  termbox.SetCursor(1, rows+1)
  termbox.Flush()
  command := ""
  for {
    ev := getev()
    msg(0, rows+1, DCOL, DCOL, prompt)
    switch ev.Key {
      case termbox.KeyEsc: return ""
      case termbox.KeyEnter: return command + "\n"
      case termbox.KeySpace: command += " "
      case termbox.KeyBackspace:
      case termbox.KeyBackspace2:
        if len(command) > 0 { command = command[:len(command)-1] }
    }
    if ev.Ch != 0 {
      command += string(ev.Ch)
      msg(1, rows+1, DCOL, DCOL, command)
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
  for {
    lin = getline(CPRMT)
    if lin == "" || lin == "\n" { break }
    i := 0;
    cursave := curln;
    var status stcode
    if getlist(lin, &i, &status) == OK {
      status = docmd(lin, &i, &status)
    }
    if status == ERR {
      msg(0, rows+1, DCOL, DCOL, "?" + strings.Repeat(" ", cols-1))
      curln = min(cursave, lastln)
      termbox.SetCursor(1, rows+1)
      termbox.Flush()
      getev()
    }
  }
  rows++
}
