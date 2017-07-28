# elapsed
Primitive to create CLI countdown or pomodoro timer.

This is a simple tool and a `bash` script, which in combination with `watch`,
`notify-send` and `paplay` (or `mplayer`) gives a spartan CLI pomodoro timer
and time tracker.

Go compiler is required to build the tool.
Script can be modified (or completely replaced) to use more suitable alternatives.

How to use it:

    # Make
    go build

    # Copy to convenient location
    cp ./elapsed ./pomodoro-timer* ~/bin

    # Run with defaults (change them right in the script if required)
    pomodoro-timer

    # Run with explicite parameters
    pomodoro-timer 40m "Take a break!" "~/sounds/bell.ogg"
    pomodoro-timer 10m "Start work!" "~/sounds/bell.ogg"

    # Run via helper scripts (change parameters within scripts as you want)
    pomodoro-timer-work
    pomodoro-timer-break

It will display something like this:
    
    Pomodoro timer
    2017-03-31T00:32:08+07:00
    29.561768025s

First string is a simple title (can be adjusted in script),
second string is a calculated time of the event,
third string is a remaining time to the event.

When timer ends, script will send desktop notification and play sound,
also timer's title and remaining time will change color and `watch`
will stop updates. Of course, script can be stopped at any time with `Ctrl+C`.

I found it useful to track actual time spent on the task, so I added simple
logging to the script. Script creates `elapsed.log` file in `CSV` format in the
working directory (I usually start it at $HOME, so it conveniently created at
$HOME directory in my case). Whenever script is stopped (by timer or with
`Ctrl+C`) it adds a record to the file with timer title, time when it was
started and stopped and elapsed time.  This file can be easily imported to
Libreoffice Calc, for instance, to calculate time spent on task during week or
month, etc.  Also, it's possible to start timers from different working
directories to track time spent on different tasks.
