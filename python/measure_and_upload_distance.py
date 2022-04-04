import RPi.GPIO as GPIO
import os
import time
import logging
import requests
import json



# # Define GPIO to use on Pi
GPIO.setmode(GPIO.BCM)
GPIO.setwarnings(False)
GPIO_TRIGGER = 17
GPIO_ECHO = 22

TRIGGER_TIME = 0.00001
MAX_TIME = 0.015  # max time waiting for response in case something is missed
GPIO.setup(GPIO_TRIGGER, GPIO.OUT)  # Trigger
GPIO.setup(GPIO_ECHO, GPIO.IN, pull_up_down=GPIO.PUD_UP)  # Echo

GPIO.output(GPIO_TRIGGER, False)


# This function measures a distance


# puts out_time
def round_trip_distance_cm_to_s(distance):
    # speed of sound =  340M / second
    speed_of_sound_cm_s = 34300.0
    return float(distance) * 2.0 / speed_of_sound_cm_s

def measure():
    logging.debug("Configuring GPIO for measurement")
    # Pulse the trigger/echo line to initiate a measurement
    GPIO.output(GPIO_TRIGGER, True)
    time.sleep(TRIGGER_TIME)
    GPIO.output(GPIO_TRIGGER, False)

    # ensure start time is set in case of very quick return
    start = time.time()
    timeout = start + MAX_TIME

    # set line to input to check for start of echo response
    while GPIO.input(GPIO_ECHO) == 0 and start <= timeout:
        start = time.time()

    if (start > timeout):
        return -1

    stop = time.time()
    timeout = stop + MAX_TIME
    # Wait for end of echo response
    while GPIO.input(GPIO_ECHO) == 1 and stop <= timeout:
        stop = time.time()

    if (stop <= timeout):
        elapsed = stop - start
        distance = float(elapsed * 34300) / 2.0
    else:
        return -1
    return distance


def convert_water_distance_to_water_height_gallons(water_distance):
    pass

def configure_logging():
    logging.basicConfig(
        format='%(asctime)s %(levelname)-8s %(message)s',
        filename='water_tank_measurement.log',
        encoding='utf-8',
        level=logging.DEBUG
    )
    return


def upload_distance(bearer_token, distance):

    url = os.getenv('WATER_TANK_MONITOR_SERVICE_URL')
    myobj = { 'water_level' :  distance } 
    headers = {'Authorization': 'Bearer ' + bearer_token}
    result = requests.post(url, data = json.dumps(myobj), headers=headers)
    if result.status_code != 201:
        logging.error("Failed to get expected status code")
    else:
        logging.info("Measurement uploaded distance: " + str(distance) + " cm status:" + str(result.status_code))


if __name__ == '__main__':
    try:
        configure_logging()
        logging.debug('Starting Measurement Script')
        if True:
            distance = measure()
            bearer_token = os.getenv('MEASUREMENT_UPLOAD_API_BEARER_TOKEN')
            if bearer_token == "":
                logging.debug("Missing environment variable for bearer authentication")
                exit(1)
            upload_distance(bearer_token, distance)
            if distance > -1:
                logging.debug("Measured Distance = %.1f cm" % distance)
            else:
                logging.warn("No Echo!")
        # Reset by pressing CTRL + C
    except KeyboardInterrupt:
        logging.info("Measurement stopped by User")
        GPIO.cleanup()
