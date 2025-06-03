package main

type editorState struct {
  buf    []buftype
  curcl  int
  curln  int
  lastln int
}

var undoStack []editorState
var redoStack []editorState

func backup() {
  snapshot := editorState{
    buf:    make([]buftype, len(buf)),
    curcl:  curcl,
    curln:  curln,
    lastln: lastln,
  }
  copy(snapshot.buf, buf)
  undoStack = append(undoStack, snapshot)
  // Clear redo stack on new action
  redoStack = nil
}

func undo() {
  if len(undoStack) == 0 {
    return // nothing to undo
  }
  // Save current state to redo
  snapshot := editorState{
    buf:    make([]buftype, len(buf)),
    curcl:  curcl,
    curln:  curln,
    lastln: lastln,
  }
  copy(snapshot.buf, buf)
  redoStack = append(redoStack, snapshot)

  // Restore from undo stack
  last := undoStack[len(undoStack)-1]
  undoStack = undoStack[:len(undoStack)-1]
  buf = make([]buftype, len(last.buf))
  copy(buf, last.buf)
  curcl = last.curcl
  curln = last.curln
  lastln = last.lastln
}

func redo() {
  if len(redoStack) == 0 {
    return // nothing to redo
  }
  // Save current state to undo
  snapshot := editorState{
    buf:    make([]buftype, len(buf)),
    curcl:  curcl,
    curln:  curln,
    lastln: lastln,
  }
  copy(snapshot.buf, buf)
  undoStack = append(undoStack, snapshot)

  // Restore from redo stack
  last := redoStack[len(redoStack)-1]
  redoStack = redoStack[:len(redoStack)-1]
  buf = make([]buftype, len(last.buf))
  copy(buf, last.buf)
  curcl = last.curcl
  curln = last.curln
  lastln = last.lastln
}
