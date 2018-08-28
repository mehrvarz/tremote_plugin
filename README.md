# TRemote Plugin - base package

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these kind of machines, specifically over Bluetooth. There is no limit to what you can remote control. You can access a list of predefined actions, you can execute executables and shell scripts, you can issue http request, and you can invoke your own or 3rd party native code plugins.

To build TRemote plugins you must use Go 1.11 and import this repository:

```
import "github.com/mehrvarz/tremote_plugin"
```

This repository acts as an interface between the TRemote service for Linux and your plugin. To see how this package can be used, take a look at the [TRemote Plugin rpi_gpio](https://github.com/mehrvarz/tremote_plugin_rpi_gpio) project.

This repository provides access to several TRemote Host helper classes:


# RemoteControlSpec

This struct is handed over to plugins with information specific to the pressed button and it's configuration. Most likely you will make use of the []StrArray element in order to retrieve one or more button specific arguments. How these arguments are used is specific to the individual plugin.

# PluginHelper

This struct is handed over to plugins. It contains several important objects. 

**PrintInfo**

This function lets a plugin provide runtime updates.

**PrintStatus**

This function lets a plugin provide status info.

**StopCurrentAudioPlayback**

This function lets a plugin stop any currently active audio player.

**StopAudioPlayerChan**

A Go channel that lets a plugin receive termination requests.

**PauseAudioPlayerChan**

A Go channel that lets a plugin receive pause requests.

**PluginIsActive**

A boolian for the plugin to set true while it is active.

**PIdLastPressed**

Pointer to an int containing the pid of the most recently pressed button.

**PLastPressActionDone**

An array of booleans containing button-event taken-care-of flags.

**PLastPressedMS**

An array of int64 elements prepresenting the time in MS when each button was last pressed.


# Compatibility

The TRemote version that supports plugins is TRemote v2.0.

In order to create plugins to work with TRemote v2.x, you need to make sure the following items are true:

- Build your plugins with Go 1.11 (exactly)

- Use the follwing packages:

  - github.com/mehrvarz/tremote_plugin v1.0.8

  - github.com/mehrvarz/log v1.0.1

- Optional: ~If~ you want to use any of these packages, make sure you use the exact same versions:

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


