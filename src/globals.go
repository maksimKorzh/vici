package main

import "github.com/nsf/termbox-go"

/* types for vici */
type stcode int
type buftype struct {
  txt string   /* text of line */
}

type edst struct {
  buf    []buftype
  curcl  int
  curln  int
  lastln int
}

/* const declarations for vici */
const TABS  = 4      /* TAB width */
const TABR  = ' '    /* TAB replace char */
const CPRMT = ":"    /* command prompt */
const APRMT = ">"    /* text prompt */
const SCRDN = ".-4"  /* command to execute on PageDown */
const SCRUP = ".+4"  /* command to execute on PageUp */

const (
  VIEW int = iota
  EDIT
  REPLACE
)

const (
  ENDDATA stcode = iota
  ERR
  OK
)

const (
  BLANK    = ' '
  TAB      = '\t'
  COMMA    = ','
  SEMICOL  = ';'
  CURLINE  = '.'
  PERIOD   = '.'
  LASTLINE = '$'
  SCAN     = '/'
  BACKSCAN = '\\'
  PLUS     = '+'
  MINUS    = '-'
  NEWLINE  = '\n'
  iCMD     = 'i'
  ACMD     = 'A'
  pCMD     = 'p'
  qCMD     = 'q'
  JCMD     = 'J'
  dCMD     = 'd'
  yCMD     = 'y'
  sCMD     = 's'
  HCMD     = 'H'
  eCMD     = 'e'
  fCMD     = 'f'
  rCMD     = 'r'
  RCMD     = 'R'
  uCMD     = 'u'
  UCMD     = 'U'
  wCMD     = 'w'
  xCMD     = 'x'
  hCMD     = 'h'
  jCMD     = 'j'
  kCMD     = 'k'
  lCMD     = 'l'
  ZCMD     = '<'
  XCMD     = '>'
)

/* This is how you can define colorscheme */
//const (
//  COL1 = termbox.ColorDefault
//  COL2 = termbox.ColorBlack
//  COL3 = termbox.ColorWhite
//  COL4 = termbox.ColorBlue
//  COL5 = termbox.ColorGreen
//  COL6 = termbox.ColorCyan
//  COL7 = termbox.ColorYellow
//  COL8 = termbox.ColorYellow
//  COL9 = termbox.ColorMagenta
//  BOLD = termbox.AttrBold
//)

const (
  COL1 = termbox.ColorDefault
  COL2 = termbox.ColorBlack
  COL3 = termbox.ColorWhite
  COL4 = termbox.ColorBlue
  COL5 = termbox.ColorLightMagenta
  COL6 = termbox.ColorDarkGray
  COL7 = termbox.ColorYellow
  COL8 = termbox.ColorYellow
  COL9 = termbox.ColorDarkGray
  BOLD = 0
)

/* variables for vici */
var KEYWORDS = []string {
  "import", "as", "from", "in", "with", "global",
  "and", "or", "fi", "then",
  "try", "except", "pass",
  "if", "else", "elif",
  "for", "do", "done", "while", "break",
  "let", "var", "const", "iota", "type",
  "int", "char", "float", "double", "rune", "byte",
  "func", "function", "def",
  "return", "undefined",
  "async", "await", "finally", "then",
}

var unst []edst       /* undo buffer */
var rest []edst       /* redo buffer */
var buf[]buftype      /* editor's buffer */
var cpb[]buftype      /* copy buffer */
var autoindent int    /* auto indentation toggler */
var hl int            /* syntax highlight toggler (visual mode) */
var mode int          /* visual/command mode toggler */
var dirty bool        /* modified flag */
var lnwidth int       /* line number width (visual mode) */
var rows, cols int    /* number of rows and columns (visual mode) */
var offrw, offcl int  /* offsets for scrolling (visual mode) */
var line1 int         /* first line number */
var line2 int         /* second line number */
var nlines int        /* # of line numbers specified */
var curln int         /* current line -- value of dot */
var curcl int         /* current column (visual mode) */
var tabcl int         /* current column to render assuming TAB chars */
var lastln int        /* last line -- value of $ */
var pat string        /* pattern */
var lin string        /* input line */
var savefile string;  /* remembered file name */
