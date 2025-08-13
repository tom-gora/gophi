# gophi

A command-line tool built in Go to display a menu of tasks or execute a specific task based on a JSON configuration file. Intended to be used with rofi, dmenu, or other similar menu applications.

Installation

1.  Ensure you have Go installed.
2.  Clone this repo: `git clone <repository_url> && cd gophi`
3.  Run `go build ./cmd/gophi` to build the executable.
4.  Move the `gophi` executable to a directory in your system's PATH so that your menu application can call it.

Usage

Run `gophi` with the following arguments:

*   `-m, --menu-config STRING`: Path to the JSON menu configuration file.
*   `-e, --exec STRING`: Command to execute as returned from picker.

It is mandatory to use either of those.

Example:

`gophi -m /path/to/config.json` (prints strings to feed into rofi)

`gophi -e "Task Name"` (executes the task with the name "Task Name")

TODOs

*   See into rofi Icon passing API to check if paths to images can be passed as icon string. If so handle paths.
*   Implement error handling and logging improvements so they consistently pass message to script wrapper calling notify-send like my other system setups.

License

MIT License
