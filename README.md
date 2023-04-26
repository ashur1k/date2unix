# Date/time converter to UNIX timestamp and back

A Go program that converts the date/time from string format to UNIX timestamp format and back.
## Usage

The program uses command-line flags to control its behavior. Available flags:

- `-utc` specifies that the time returned should be in UTC (Coordinated Universal Time) format.
- `-milli` indicates that milliseconds must be included in the returned UNIX time stamp.
- `-format` allows the user to specify his own date/time format in golang "time.Time" format.
- `-help` outputs help on how to use the program.

## Examples of use:

```sh
    # Convert date/time to UNIX time stamp
    $ date2unix "2023-04-25 13:30:00".
    Timestamp: 1682356200

    # Convert UNIX timestamp to date/time
    $ date2unix 1682356200
    2023-04-25 13:30:00 +0300 MSK
```

## Installation

To install the program, you must:

    Install Go version 1.13 or higher
    Clone the repository: `git clone https://github.com/ashur1k/date2unix.git`
    Go to the folder with the program: `cd date2unix`
    Build the program: `go build -o date2unix`.
    Run the program: `./date2unix`.

## License

The program is licensed under the MIT license. More information can be found in the LICENSE file.