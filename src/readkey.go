package main

import "os"
import "github.com/nsf/termbox-go"

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
      }
    }
  } else {
    switch ev.Key {
      /* commands */
    }
    //if current_col > len(text_buffer[current_row]) { current_col = len(text_buffer[current_row]) }
  }
}
