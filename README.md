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


