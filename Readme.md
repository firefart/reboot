# reboot

Reboot is a simple golang based windows service that initiates a reboot upon receiving a specific message on a TCP server.

I connect to my windows machine mostly via RDP. I have the strange phenomenon that it stops responding to any kind of samba, netbios, rdp or other requests. This behaviour is fixed after a physical login but that's not always possible.

Without these ports responding using tools like psshutdown, WMI or remote registry is not possible.

Thus I needed a hacky way to restart the machine and so this service was created.

You can connect via netcat or telnet to port 1234 and send the `REBOOT` keyword. After that the program will ask for a password the user can supply and if it's correct a clean reboot will be executed. All attempts and reboots will be logged to the Application Eventlog.

**Attention:** This program is not considered safe and only contains basic authentication mechanisms. It should only be used in places where you have complete control over your network. Please do not use this on enterprise networks.

## Install

- build the binary on a windows machine using `go build` or by executing `make.bat`

- copy `reboot.exe` and `password.conf` to a directory (for example `C:\ProgramFiles\reboot\`)

- edit `password.conf` and put in a password of your choice

- open an elevated command or powershell prompt and cd into the installation directory

- execute `reboot.exe install` to install the service

- execute `reboot.exe start` to start the service

- check with `netstat -an` if port 1234 is listening

## Example

```text
nc x.x.x.x 1234
REBOOT
Please enter password: password
Initiating reboot...
Done
```
