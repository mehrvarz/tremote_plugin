# runtrp - run TRemote plugin tool

TRemote is a service for ARM based Linux computers. It lets you remote control *things* on these kind of machines, specifically over Bluetooth. There is no limit to what you can remote control on a target device because you can extend the remote control service with with native code plugins.

When developing these plugins it can often be useful to be able to ~run~ them from the command line. This is what the runtrp tool does and this is how to use it:


```
runrtp abc-plugin.so arg1
```

This will behave just like it would, if abc was a reguar command line executable. In which case you would run it like so:

```
abc arg1
```

One difference: in order to hand over multiple arguments to a plugin, you need separate your arguments with pipe characters ('|'): 

```
runrtp abc-plugin.so arg1|arg2|arg3
```


# Download runtrp for linux.ARM6 (Raspberry Pi)

To download runtrp for linux.ARM6, click on the "runrtp" link on top of this page and then click Download.

You can download runtrp for linux.AMD64 [here](https://github.com/mehrvarz/tremote_plugin/tree/master/bin.linux.AMD64).


