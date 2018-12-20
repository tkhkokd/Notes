

# LINUX

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

## Mosquitto


###Create Authentication

1. First, stop the mosquitto
```sudo stop mosquitto```
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
