# elapsed
Primitive to create CLI countdown or pomodoro timer.

This is a simple tool and a `bash` script, which in combination with `watch`,
`notify-send` and `paplay` (or `mplayer`) gives a spartan CLI pomodoro timer.

Go compiler is required to build the tool.
Script can be modified (or completely replaced) to use more suitable alternatives.

How to use it:

    # Make
    go build

    # Copy to convenient location
    cp ./elapsed ./pomodoro-timer ~/bin

    # Run with defaults (change them right in the script if required)
    pomodoro-timer

    # Run with explicite parameters
    pomodoro-timer 40m "Take a break!" "~/sounds/bell.ogg"
    pomodoro-timer 10m "Start work!" "~/sounds/bell.ogg"

It will display something like this:
    
    Pomodoro timer
    2017-03-31T00:32:08+07:00
    29.561768025s

First string is a simple title (can be adjusted in script),
second string is a calculated time of the event,
third string is a remaining time to the event.

When timer ends, it'll send desktop notification and play sound,
also timer's title and remaining time will change color and `watch`
will stop updates.
