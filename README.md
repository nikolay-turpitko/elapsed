# elapsed
Primitive to create CLI countdown or pomodoro timer.

This is a simple tool and a bash script, which in combination with `bash`,
`watch`, `notify-send` and `paplay` (or `mplayer`) gives a spartan CLI pomodoro
timer.

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

