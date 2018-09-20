# TRemote Plugin - base package

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these kind of machines, specifically over Bluetooth. There is no limit to what you can remote control. You can access a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can invoke your own or 3rd party native code plugins.

To build TRemote plugins you must use Go 1.11 and import these repositories:

```
import "github.com/mehrvarz/tremote_plugin"
import "github.com/mehrvarz/log"
```

# Compatibility

Support for TRemote plugins has started with TRemote v2.0.

In order to create plugins compatible with TRemote v3.0, the following conditions must be met:

- TRemote plugins must be built with Go v1.11

  TRemote plugins are based on "Go Modules".
  For this reason a "go.mod" file must exist in the project folder of any TRemote plugin.
  In it all imported packages must be listed.

- The following packages must be imported:

  - github.com/mehrvarz/tremote_plugin v1.0.12
  - github.com/mehrvarz/log v1.0.1

- Optional: In case any of the following packages are to be included, the same versions must be used:

  - github.com/mehrvarz/go_queue v0.0.0-20180811045238-f34b4ebf5df4
  - github.com/go-ble/ble v0.0.0-20180718090407-11b1dad1df3d
  - github.com/dhowden/tag v0.0.0-20180815181651-82440840077f
  - github.com/pkg/errors v0.8.0
  - github.com/tarm/serial v0.0.0-20180114052751-eaafced92e96
  - github.com/stretchr/testify v1.2.2
  - github.com/pmezard/go-difflib v1.0.0
  - github.com/mattn/go-isatty v0.0.3
  - github.com/mattn/go-colorable v0.0.9 
  - github.com/davecgh/go-spew v1.1.1
  - github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
  - github.com/mgutz/logxi v0.0.0-20161027140823-aebf8a7d67ab



# TRemote plugin main entry

The main entry point every TRemote plugin needs to implement is the Action() method.

```
func Action(
	log log.Logger,
	pid int,
	longpress bool,
	pressedDuration int64,
	homeDirectory string,
	rcs* tremote_plugin.RemoteControlSpec,
	ph tremote_plugin.PluginHelper,
	wg *sync.WaitGroup) error {
	// ...
}
```

Action() needs to return as quickly as possible.
Long running operations must take place in dedicated goroutines.

Action() method arguments:

### log log.Logger

See: [log](https://godoc.org/github.com/alexcesaro/log)

### pid int

Button ID: 0=P1, 1=P2, ...

### longpress bool

false: Action may be a short or longpress.

true: Action was specified as P#L and should be immediately processed as longpress.

### pressedDuration int64

0: Button has just been pressed.

else: Button has just been released. pressedDuration expresses the press duration in milliseconds.

### homeDirectory string

The path to the installation directory.

### rcs* tremote_plugin.RemoteControlSpec

This struct is handed over to plugins with information specific to the pressed button and it's configuration. 
The most important information are arguments in []StrArray handed over from mapping.txt. 
How these arguments may be used is plugin specific.

### ph tremote_plugin.PluginHelper

see PluginHelper below

### wg* sync.WaitGroup

For long term operations the waitgroup methods Add() and Done() must be invoked.


# PluginHelper

This struct is handed over to plugins. It contains several important objects. 

### PrintInfo(info string)

This function lets a plugin provide runtime info updates, such as playing song titles.

### PrintStatus(info string)

This function lets a plugin provide additional status or error info.

### StopCurrentAudioPlayback()

This function lets a plugin stop any currently active audio playback. Possibly emitted from another plugin.

### StopAudioPlayerChan *chan bool

A pointer to a Go channel that lets a plugin receive stop requests.

### PauseAudioPlayerChan *chan bool

A pointer to a Go channel that lets a plugin receive pause and unpause requests.

### PluginIsActive *bool

Do not use.

### PIdLastPressed *int

A pointer to an int containing the number of the most recently pressed button, before the button that activated this plugin.

### PLastPressActionDone [tremote_plugin.MaxButton]bool

An array of booleans containing "taken-care-of" flags for every button.

### PLastPressedMS [tremote_plugin.MaxButton]int64

An array of int64 elements containg time stamps (MS) of the start time of the most recent button press.

### ImageInfo(jpgData []byte, mime string)

This function enables a plugin to hand over raw JPEG data to be used as background image on the client.

### HostCmd(command string, parameter string) string

This function lets a plugin call functionality offered by TRemote host.

"ScreenPower" switches the screen attached to the host on or off. Parameter may be set to "on" or "off".

"AudioMute" switches the audio mute function. Parameter may be set to "on", "off" or "toggle".

Example: ph.HostCmd("ScreenPower","on")


