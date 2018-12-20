


MQTT
Message Queue Telemetry Transport
- Low Power
- Cheap
- talk over IP

MQTT Broker
- Manages communication between devices, mobile apps, web apps, and servers.


TCP (Transmission Control Protocol)
- Deals with stream of data

UDP (User Datagram Protocol)
- Alternative to TCP



## Mosquitto


Create Authentication
On the Linux server,
```
sudo mosquitto_passwd -c /etc/mosquitto/passwd <user_name>
```
creates a password file. ```-c``` = create
