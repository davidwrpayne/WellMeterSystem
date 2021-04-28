# WellMeterSystem
This Project is to provide a script that uses a Raspberry pi to monitor the height of water in a Water Storage System
It will come with a script you can configure using env vars to point at a recording service.

In my case The recording service is my website.


The script executes a POST to url in the Environment Variable WATER_TANK_MONITOR_SERVICE_URL
The body of the POST is JSON. E.g.  `{'water_level',30.5}`

Where 30.5 is the distance from the sensor in cm.

# TODO
- possible improvement is to calculate the distance based on temprature and humdity if you had those sensors available.
- Use ENV vars for GPIO pin configuration
- Use ENV var to configue the output log location
