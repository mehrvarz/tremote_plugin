# TRemote Plugin - base package

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these kind of machines, specifically over Bluetooth. There is no limit to what you can remote control. You can access a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can invoke your own or 3rd party native code plugins.

To build TRemote plugins you must use Go 1.11 and import these repositories:

```
import "github.com/mehrvarz/tremote_plugin"
import "github.com/mehrvarz/log"
```
Find more info in paragraph Compatibility below.


# Helper classes

Repository "tremote_plugin" acts as a interface between the TRemote service and your plugin. 
It provides access to the following helper classes:

## RemoteControlSpec

This struct is handed over to plugins with information specific to the pressed button and it's configuration. Most likely you will make use of the []StrArray element in order to retrieve one or more button specific arguments. How these arguments are used is specific to the individual plugin.

## PluginHelper

This struct is handed over to plugins. It contains several important objects. 

**PrintInfo(info string)**

This function lets a plugin provide runtime info updates.

**PrintStatus(info string)**

This function lets a plugin provide status or error information.

**StopCurrentAudioPlayback()**

This function lets a plugin stop any currently active audio playback. Possibly emitted from another plugin.

**StopAudioPlayerChan *chan bool**

A pointer to a Go channel that lets a plugin receive stop requests.

**PauseAudioPlayerChan *chan bool**

A pointer to a Go channel that lets a plugin receive pause and unpause requests.

**PluginIsActive *bool**

A pointer to a boolian a plugin must set to true while it is active.

**PIdLastPressed *int**

A pointer to an int containing the number of the most recently pressed button, before the button that activated this plugin.

**PLastPressActionDone [tremote_plugin.MaxButton]bool**

An array of booleans containing "taken-care-of" flags for every button.

**PLastPressedMS[tremote_plugin.MaxButton]int64**

An array of int64 elements containg time stamps (MS) of the start time of the most recent button press.


# Compatibility

Support for TRemote plugins started in TRemote v2.0.

In order to create plugins that work with TRemote v2.x, the following conditions must be matched:

- Build your plugins with Go v1.11 (exactly)

  Take advantage of Go Modules and deploy a go.mod file in your project.

- You must import the follwing packages:

  - github.com/mehrvarz/tremote_plugin v1.0.8
  - github.com/mehrvarz/log v1.0.1

- Optional: IF you intend to use any of the following packages, make sure you use the same versions. Check your go.mod file:

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


