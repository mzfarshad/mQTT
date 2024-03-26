# MQTT IN GOLANG

## GOAL

#####   Using the mqtt protocol to store data in the postgres database and receive data from the database

### Save Car

#####   To save data, you can use the   *cars/add-car*   topic to send a JSON, which is an example of a car and has name, company, and color fields.

### Receive Cars

#####   You can use two methods to receive data:

1. #### Receive car by ID
  ##### You can get a car with a specific ID by using the topic *cars/get-car/< car_id >*

2. #### Receive all Cars
  #####   You can use the *cars/all-cars* topic to get all available cars
