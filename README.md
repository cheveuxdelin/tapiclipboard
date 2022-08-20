# Tapiclipboard

A superset of the system's clipboard!

<img width="967" alt="Screen Shot 2022-08-02 at 22 15 04" src="https://user-images.githubusercontent.com/68461123/182529329-3bfd8233-9605-4a1e-a706-aec93d072c1e.png">

## What's this?

This is just an CLI App *(soon to able to be run as daemon)* that overpowers the clipboard we all know (CTRL+C, CTRL+V) to a stack system, where you can save multiple inputs to the clipboard, and pop them as FIFO (First in - First out)! 💯⚡️

## How does it work?

Tapiclipboard consists of a server application and a client application, that interact with each other with requests and responses using named pipes. The server is meant to be run as a service/daemon that is still yet to be implemented. It uses the system's commands to directly copy the output to the user's system clipboard for direct use!

## That's great! How can I use it?

As of right now, there's only support for macOS since is the OS I use, I still have to make it cross-platform for Windows and Linux🐧. Since it's not yet implemented as a daemon, the way to run it would be to have Go installed, build both the sever and the client. Have the server running either background or foreground process and just use the client cli as you would... use the command help inside the CLI for more info :)

![nose](https://media.giphy.com/media/pMFmBkBTsDMOY/giphy.gif)
