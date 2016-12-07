# impd - Intermapper Pagerduty notifier

This is a simple program for triggering, acknowledging and resolving incidents
on [pagerduty](https://www.pagerduty.com/).

## Installation

In order to compile from source, golang needs to be installed on the machine
compiling the code ([Go Installation
Instructions](https://golang.org/doc/install)). To ease installation of
dependencies we recommend installing [glide](https://glide.sh/).

1. Download package `go get github.com/blacksails/impd`
1. Install dependency with glide `glide install` or just download the
   dependency `go get github.com/blacksails/pdevent`
1. Compile package `go build`
1. Move binary to the Intermapper Tools folder. The tools folder can be found
   within the Intermapper settings folder. For more info about the location of
   the folder, see the [intermapper
   docs](http://download.intermapper.com/docs/UserGuide/Content/09-Reference/09-03-Files_and_Folders/filesfolders.html).
1. Ensure that the binary has the right permissions so that Intermapper can
   execute it.

## Configuration

1. Add a new notifier with the Command Line type.
1. Insert command with pagerduty service key into the command field.
   ```
   impd -k "<PAGERDUTY_SERVICE_KEY>"
   ```
1. Edit the message and add the following exact message.
   ```
   -e '<Event>' \
   -d '<Smart Name>: <Device Condition>' \
   -i '<Device-IMID>'
   ```
1. Configure the notifier for the probes that you want
