# keydetect

Simple CLI app that listens to keyboard input and displays the event code and some additional info.

## Background

Recently riced my Arch Linux machine and an update to a newer kernel version switched my keyboard mode.
That led to hours of searching on a separate computer what was happening because I couldn't type simple
characters like `+-` and therefore couldn't type my full password to login. Knowing I had no access to a browser,
I couldn't rely on
a [web-based JS keycode detector](https://www.freecodecamp.org/news/javascript-keycode-list-keypress-event-key-codes/)
so I decided to write a quick utility I can install on any machine to see what keys I'm pressing. :)

## Features

- [x] Show history of keys pressed
- [ ] Show ASCII Char code
- [ ] Show UTF-8 Char code
- [ ] Show KeyCode
- [ ] Show Modifier status in help?
- [x] Show String representation
- [ ] Option to auto-quit after 1 minute
- [ ] Option to limit or hide history
- [ ] Detect keyboard config at startup

## Installation

!TODO

### From source

!TODO

### From AUR or other package managers

## Usage

Simply run `keydetect` and press any key you'd like to know some info about!

All helpers should be displayed by the application but in case they are not, simply run `keydetect --help`.

## License

[MIT](/LICENSE)
