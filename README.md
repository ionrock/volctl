# Volctl #

Volctl is a pretty silly web service to run on a Raspberry Pi that lets you adjust the volume via a web interface. This is not a very pretty web interface, but it is an interface nonetheless!

## What does it do?

Volctl just uses `amixer` get the current PCM volume. From there,  you can use a HTML slider to adjust it.

That's it!

## How do I use this essential tool?

There is a `Makefile` that has some targets for building a binary and copying it to your Raspberry Pi. There is also a systemd service file for starting it up when the Raspberry Pi boots.
