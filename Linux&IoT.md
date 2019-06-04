

# LINUX

### Process

```htop``` Shows all processes

```kill -HUP PID```
```PID```: Process ID
```-HUP``` could be ```-SIGHUP```: Signal Hangup

### Packages
List packages
```
apt list --installed
```
Updates the package list, get information about newest version of the packages and their dependencies.
```
sudo apt-get update
```
Upgrades the existing packages, if APT knows about the new versions by running the above command ```apt-get update```
```
sudo apt-get upgrade
```

###Sessions

Keep the server session running, use ```screen```

Create new session
```screen```, with a session name ```screen -S SESSION_NAME```

List Sessions
```screen -ls```

Detache (Leave the screen, != terminate)
```Ctrl+<a>, <d>```

Resume to the screen
```screen -r SESSION_ID```

Split the screen (Vertical)
```Ctrl+<a>, <|>(vertical line)```

Undo screen split
```Ctrl+<a>, <X>```

Terminate screen
```Ctrl+<a>, <K>```


## Mosquitto

### Mosquitto Configuration



### Create Authentication

```conf```file must say
```
allow_anonymous false
password_file PASSWORD_FILE_PATH
```

Create password file in ```etc\mosquitto ```
```mosquitto_passwd -b passwordfile USER_NAME PASSWORD```


1. First, stop the mosquitto
```sudo service mosquitto stop```
2. creates a password file. ```-c``` = create
```
sudo mosquitto_passwd -c /etc/mosquitto/passwd <user_name>
```
3. Edit the config file.
```sudo gedit /etc/mosquitto/mosquitto.conf```
4. Add
```
password_file /etc/mosquitto/passwd
allow_anonymous false
```
5. Re-start the Mosquitto Broker by,
```
mosquitto -c /etc/mosquitto/mosquitto.conf
```



