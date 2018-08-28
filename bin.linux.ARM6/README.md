# runtrp - run tremote plugin tool

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these kind of machines, specifically over Bluetooth. There is no limit to what you can remote control on the target device because you extend the system with small plugins with native code.

When developing such a plugin it can often be useful to ~run~ it from the command line. This is what the runtrp tool is for. This is how it works:


```
runrtp abc-plugin.so arg1
```

This behaves just like it would, if abc was a normal native executable. In which case you would run it like so:

```
abc arg1
```

One difference: to hand over multiple arguments to a plugin, you must use pipe characters ('|'): 

```
runrtp abc-plugin.so arg1|arg2|arg3
```


# Download runtrp for linux.ARM6 (Raspberry Pi)

To download runtrp for linux.ARM6, click the "runrtp" link on top and then click Download.

To download runtrp for linux.AMD64, do the same [here](https://github.com/mehrvarz/tremote_plugin/tree/master/bin.linux.AMD64).


