package main

/* edittype -- types for in-memory version of edit */
type stcode int
type buftype struct {
  txt string   /* text of line */
  mark bool    /* mark of line */
}

/* editcons -- const declarations for edit */
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

/* editvar -- variables for edit */
var buf[]buftype      /* editor's buffer */
var cpb[]buftype      /* copy buffer */

var line1 int         /* first line number */
var line2 int         /* second line number */
var nlines int        /* # of line numbers specified */
var curln int         /* current line -- value of dot */
var lastln int        /* last line -- value of $ */

var pat string        /* pattern */
var lin string        /* input line */
var savefile string;  /* remembered file name */
