package tremote_plugin

import (
	"time"
)

const MaxButton             uint = 8
const LongPressDelay        time.Duration = 500

// RemoteControlSpec provides access to the invoked button configuration
type RemoteControlSpec struct {
	MultiButton bool        // single-button action if set false

	BIT   uint              // button bitmask
	Label string            // button label

	Fkt      string         // name of short-press action
	Args     string         // short-press arguments
	StrArray []string       // short-press arguments tokenized

	Fktlong      string     // name of long-press action
	Argslong     string     // long-press arguments
	StrArraylong []string   // long-press arguments tokenized
}

type StopFunc func() error
type PrintFunc func(string)

// PluginHelper provides access to runtime information
type PluginHelper struct {
	PrintInfo                   PrintFunc          // fkt to provide runtime info
	PrintStatus                 PrintFunc          // fkt to provide status info
	StopCurrentAudioPlayback    StopFunc           // fkt to stop other currently active audio players
	StopAudioPlayerChan         *chan bool         // chan to receive termination requests
	PauseAudioPlayerChan        *chan bool         // chan to receive pause requests
	PluginIsActive              *bool              // plugin must set true while it is active
	PIdLastPressed              *int               // pid of most recently pressed button
	PLastPressActionDone        *[MaxButton]bool   // button event taken-care-off array
	PLastPressedMS              *[MaxButton]int64  // button is pressed since MS array
}

