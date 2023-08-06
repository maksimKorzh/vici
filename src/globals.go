package main

import "github.com/nsf/termbox-go"

/* types for vici */
type stcode int
type buftype struct {
  txt string   /* text of line */
  mark bool    /* mark of line */
}

/* const declarations for vici */
const TABS  = 4      /* TAB width */
const TABR  = ' '    /* TAB replace char */
const CPRMT = ":"    /* command prompt */
const APRMT = ">"    /* text prompt */

const (
  VIEW int = iota
  EDIT
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
  PCMD     = 'p'
  NCMD     = 'n'
  QCMD     = 'q'
  ACMD     = 'a'
  DCMD     = 'd'
  CCMD     = 'c'
  ICMD     = 'i'
  EQCMD    = '='
  MCMD     = 'm'
  TCMD     = 't'
  YCMD     = 'y'
  XCMD     = 'x'
  SCMD     = 's'
  ECMD     = 'e'
  FCMD     = 'f'
  RCMD     = 'r'
  WCMD     = 'w'
)

const (
  DCOL = termbox.ColorDefault
  NCOL = termbox.ColorBlack
  WCOL = termbox.ColorWhite
  BCOL = termbox.ColorBlue
  GCOL = termbox.ColorGreen
  CCOL = termbox.ColorCyan
)

/* variables for vici */
var buf[]buftype      /* editor's buffer */
var cpb[]buftype      /* copy buffer */

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
