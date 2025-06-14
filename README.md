# VICI
Minimalist cross-platform terminal based VIM-like text editor

# Screenshot
![IMAGE ALT TEXT HERE](https://raw.githubusercontent.com/maksimKorzh/vici/main/assets/vici.png)

# Project idea
Combine powerful Linux **ed** command set with the visual interface,<br>
hence the name: (V)isual (I)nterface, (C)ommand (I)nput.<br>
Unlikely VI/VIM this project is focused on simplicity of usage<br>
and should be treated as **visual ed**, which essentially it is.

# Features
 - visual text editing
 - ED command set + some VIM motion commands
 - rudimentary syntax highlighting
 - unlimited undo/redo

# Commands

       NAVIGATION:

    : 1                go to the first line
    : $                go to the last line
    : 12               go to line number 12
    : .                go to the current line (does nothing, used in expressions)
    : .+10             scroll down 10 lines
    : .-10             scroll up 10 lines
    : $-5              scroll to 5 lines before the last line
    : h                move cursor left
    : j                move cursor down
    : k                move cursor up
    : l                move cursor right
    : <                go to first char in line
    : >                go to last char in line

       EDIT:

    : a                toggle auto indentation
    : i                switch to INSERT mode
    : R                switch to REPLACE mope
    : A                go to last char in line, switch to INSERT mode
    : J                join next line to current one
    : x                delete char under the cursor
    : u                undo last change
    : U                redo last change

       COPY / PASTE:

    : 10,23y           copy lines 10,23 inclusive to copy buffer
    : 10,23d           cut lines 10,23 inclusive to copy buffer
    : 41p              paste content of the copy buffer after line 41

       FIND / REPLACE:

    : /pat/            scroll to first "pat" occurrence ("pat" can be regexp )
    : /pat/;//         scroll to the second "pat" occurrence
    : //               scroll to the next "pat" occurrence
    : \\               scroll to the previous "pat" occurence
    : s/pat/sub/       substitute "pat" with "sub" globally ("pat" can be regexp)
    : 2,5s/pat/sub/    substitute "pat" with "sub" within lines 2,5 inclusive

       I/O:

    : e file.txt       load "file.txt" to the buffer
    : r file.txt       insert content of the "file.txt" to the current line in buffer
    : f file.txt       set current file name to "file.txt"
    : w file.txt       save file as "file.txt"
    : w                save current file
    : q                exit from editor

      MISC:

    : H                toggle syntax highlighting if available

# Shortcuts
Simple no bulk (no line range) commands are working as shortcuts, e.g. while
being in NORMAL mode press **i** to switch to INSERT mode,
or **y**/**d** to yank/delete a single line, but there are a plenty
of exceptions, listed below:

                 ESC   switch to NORMAL mode
              : or /   execute command
                   n   find next pattern
                   N   find prev pattern
              Ctrl-r   redo last change
              Arrows   move cursor
    Ctrl-d or PgDown   scroll 4 screen rows downwards
      Ctrl-u or PgUp   scroll 4 screen rows upwards
           0 or HOME   move cursor to the begining of the current line
            $ or END   move cursor to the end of the current line
            r <char>   replace single char
                   g   go to first line
                   G   go to last line
                   w   go to the beginning of the next word
                   e   go to the end of the next word
                   b   go to the beginning of the previous word

# Usage
    $ vici             open editor with "out.txt" source file name
    $ vici file.txt    open editor with "file.txt" if it exists,
                       otherwise set source filename to "file.txt"

# Latest Release
https://github.com/maksimKorzh/vici/releases/

# Donations
 - paypal "maksymkorzh@gmail.com"
 - patreon "https://www.patreon.com/code_monkey_king"
