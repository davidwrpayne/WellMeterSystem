# WellMeterSystem
This Project is to provide a script that uses a Raspberry pi to monitor the height of water in a Water Storage System
It will come with a script you can configure using env vars to point at a recording service.

In my case The recording service is my website.

The script executes a POST to url in the Environment Variable WATER_TANK_MONITOR_SERVICE_URL
The body of the POST is JSON. E.g.  `{'water_level',30.5}`

Where 30.5 is the distance from the sensor in cm.

# Improved data recording.
Working on a Go system that will record measurements to local disk / storage then attempt to send the metric to the recording website
Then if a successful write to the recording service, it will mark the record for soft delete.



configuration is in config/

software is broken up into different modules

- command line interface code
  - code for 2 commands
    - measure
    - report 
- system of record reporting
  - Publish to internet code
- sensor
  - FakeSensor
  - RaspberryPI Sensor code
- schema 
  - contains the structs that describes the measurements
- repository
  - creation of unpublished record
  - tracks unpublished records
  - publish record
  

# TODO
- possible improvement is to calculate the distance based on temprature and humdity if you had those sensors available.
- Use ENV vars for GPIO pin configuration
- Use ENV var to configure the output log location

- write to the unpublished
- publish by writing to http payne.work
  - move measurement to published folder