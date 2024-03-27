# MQTT IN GOLANG

## GOAL

#####   Using the MQTT protocol, a JSON file which is an example of a car and has name, company and color fields is stored in the Postgres database, and it can find a car with a specific ID or a list of all available cars using the topics below. Get the database

### Save Car

#####   publish :  *cars/add-car*  
#####   subscribe : *response/save-car*

### Receive Cars

#####   You can use two methods to receive data:

1. #### Receive car by ID
  #####  publish : *cars/get-car/< car_id >*
  #####  subscribe : *response/car*

2. #### Receive all Cars
  #####   publish : *cars/all-cars*
  #####   subscribe : *response/all-cars*
