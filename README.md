# Mqtt In Golang


This is a demo project which uses mqtt broker for communicating between clients and a Golang server. Clients can publish on the topic of **"cars/register"** using the Mosquitto CLI as following:

```bash

  mosquitto_pub -t "cars/register" -m '{"name":"X7","company":"BMW","color":"red"}' -p 1885

```

The **Save Car Consumer** is subscribed and listening to the above topic, and is ready to save the received car message into the database. In case of success, this consumer publishes a new message which can be tracked by subscribing into the topic of **"cars"** as following:

```bash

  mosquitto_sub -t "cars" -p 1885

```
