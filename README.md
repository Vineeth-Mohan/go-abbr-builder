# Abbrebation Builder
abbr-builder is a utility to find abbrevation from a corpus of text file and to build a dictinary file based on it

# How to works
If it find text of following formats 
```
The Internet of things (IoT) is the extension of Internet connectivity into physical devices and everyday objects. Embedded with electronics, Internet connectivity, and other forms of hardware (such as sensors), these devices can communicate and interact with others over the Internet, and they can be remotely monitored and controlled.
```
It identified `IoT` as the abbrevation for `Internet of things` , based on the pattern its presented.

# How to run
go build main.go  AbbreviationProcessor.go
./main -text <text file>
