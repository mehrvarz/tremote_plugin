# runtrp - run tremote plugin tool

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these kind of machines, specifically over Bluetooth. There is no limit to what you can remote control on the target device because you extend the system with small plugins with native code.

When developing such a plugin it can often be useful to ~run~ it from the command line. This is what the runtrp tool is for. This is how it works:


```
runrtp abc-plugin.so arg1
```

This behaves just as if abc is a regular executable. In which case you would enter the following to run it:

```
abc arg1
```

To download runtrp for linux.ARM6, click the "runrtp" link on top and then click Download.

To download runtrp for linux.AMD64, do the same [here](https://github.com/mehrvarz/tremote_plugin/tree/master/bin.linux.AMD64).


