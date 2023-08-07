# vici
Minimalist cross-platform terminal based text editor

# Screenshot
![IMAGE ALT TEXT HERE](https://raw.githubusercontent.com/maksimKorzh/vici/main/assets/vici.png)

# Features
 - visual interface
 - command set (inspired by Linux "ed")
 - syntax highlighting

# Shortcuts
       ESC - enter the 'VIEW' mode
         e - enter the 'EDIT' mode
         q -  quit from the text editor
         w -  write file to disk
         j -  cat current line to previous one
         d -  cut current line to copy buffer
         y -  copy current line to copy buffer
         p -  paste line from copy buffer
         h -  toggle syntax highlighting
         1 -  go to the first line
         $ -  go to the last line
         : -  execute command
    Arrows -  move cursor
    PgDown -  scroll 4 screen rows downwards
      PgUp -  scroll 4 screen rows upwards
      HOME -  move cursor to the begining of the current line
       END -  move cursor to the end of the current line

# Commands
    See GNU ed commands for reference, all but global
    commands are working just like in GNU ed, substitute
    command acts globally by default but can be limited
    if prefixed by the line numbers' range

    While being in 'VIEW' mode, press ":" to enter command mode.
    Input command and press enter to execute it.

       Navigation:

    : 1              go to the first line
    : $              go to the last line
    : 12             go to line number 12
    : .              go to the current line (does nothing, used in expressions)
    : .+10           scroll down 10 lines
    : .-10           scroll up 10 lines
    : $-5            scroll to 5 lines before the last line

       Find/Replace:

    : /pat/          scroll to first "pat" occurrence ("pat" can be regexp )
    : /pat/;//       scroll to the second "pat" occurrence
    : //             scroll to the next "pat" occurrence
    : \\             scroll to the previous "pat" occurence
    : s/pat/sub/     substitute "pat" with "sub" globally ("pat" can be regexp)
    : 2,5s/pat/sub/  substitute "pat" with "sub" within lines 2,5 inclusive

       Edit

    : 3i             insert text at line 3 (line is optional, default is current line)
    : 2a             append text after line 2 (line is optional, default is current line)
    : 1c             change text at line 1 (line is optional, default is current line)  

       I/O:

    : e file.txt     load "file.txt" to the buffer
    : r file.txt     insert content of the "file.txt" to the current line in buffer
    : f file.txt     set current file name to "file.txt"
    : w file.txt     save file as "file.txt"
    : w              save current file
    : q              exit from editor

       Copy/Paste/Move

    : 10,23m41       move lines 10,23 inclusive after line 41
    : 10,23t41       copy lines 10,23 inclusive after line 41
    : 10,23y         copy lines 10,23 inclusive to copy buffer
    : 10,23d         cut lines 10,23 inclusive to copy buffer
    : 41             paste content of the copy buffer after line 41
 

# Usage
    $ vici                # opens editor with 'out.txt' source file name
    $ vici my_file.txt    # opens editor with 'my_file.txt' if it exists,
                          # otherwise sets source filename to 'my_file.txt'

# Latest Release
https://github.com/maksimKorzh/vici/releases/

# Donations
 - paypal "maksymkorzh@gmail.com"
 - patreon "https://www.patreon.com/code_monkey_king"
