package main

/* min -- get min number */
func min(num1, num2 int) int {
  if num1 < num2 { return num1 }
  return num2
}

/* ctoi -- convert string at sri] to integer, increment i */
func ctoi(s string, i *int) int {
  var n, sign int
  skipbl(s, i)
  if s[*i] == '-' {
    sign = -1
  } else {
    sign = 1
  }
  if s[*i] == '+' || s[*i] == '-' { *i++ }
  for *i < len(s) && s[*i] >= '0' && s[*i] <= '9' {
    n = n * 10 + int(s[*i]) - '0'
    *i++
  }
  return sign * n
}

/* getst -- get status code description */
func getst(status stcode) string {
  descr := ""
  switch status {
    case ENDDATA: descr = "ENDDATA"
    case ERR: descr = "ERR"
    case OK: descr = "OK"
  }
  return descr
}
