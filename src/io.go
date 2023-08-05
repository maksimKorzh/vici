package main

import "os"
import "fmt"
import "bufio"
import "strings"

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
  fmt.Println(count)
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
  fmt.Println(count)
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

/* getline -- read user input from STDIN */
func getline() string {
  var line string
  var ch rune
  for ch != '\n' {
    _, err := fmt.Scanf("%c", &ch)
    if err != nil { panic(err) }
    line += string(ch)
  }
  line = strings.Replace(line, "\r", "", -1)
  return line
}
