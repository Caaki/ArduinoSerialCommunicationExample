import getopt
import struct
import sys
import time
import serial

argumentList = sys.argv[1:]

options = "h:a:v:p:"
long_options = ["Help", "Action=", "Value=", "Port="]

port = "/dev/ttyUSB0"
action = ""
value = ""

try:
    arguments, values = getopt.getopt(argumentList, options, long_options)
    for currentArgument, currentValue in arguments:
        if currentArgument in ("-h", "--Help"):
            print("Displaying Help")
            print("-a --Action")
            print("\t-a send : Send a signal to the Arduino to be sent over the IR")
            print("\t-a receive : Waits for the Arduino to receive and send back data")
            print("-v --Value")
            print("\t-v <RAW_DATA>: This is used in combination with -a send to define what to send to the Arduino")
            print("-p --Port")
            print("\t-p <PORT> : set a custom port for the communication")
            print("\tDefault port is /dev/ttyUSB0")
        elif currentArgument in ("-a", "--Action"):
            action = currentValue
        elif currentArgument in ("-v", "--Value"):
            value = currentValue
        elif currentArgument in ("-p", "--Port"):
            port = currentValue

    if action == "send":
        if not value:
            print("Error, Raw data not provided")
            sys.exit(1)

        ser = serial.Serial(port, 9600, timeout=1)
        ser.setDTR(False)
        time.sleep(1)
        ser.flushInput()
        ser.setDTR(True)
        time.sleep(2)

        values = value.replace(', ', ',')
        values = values.split(",")
        irValues = [int(v) for v in values]
        array_length = len(irValues)

        ser.write("send\n".encode('utf-8'))
        ser.write(struct.pack('<H', array_length))

        for num in irValues:
            ser.write(struct.pack('<H', num)) 

        # print(f"Array of length {array_length} sent. Waiting for response...")

        received_numbers = []
        for _ in range(array_length):
            low_byte = ser.read(1)
            high_byte = ser.read(1)

            if len(low_byte) == 1 and len(high_byte) == 1:
                received_num = struct.unpack('<H', low_byte + high_byte)[0]
                received_numbers.append(received_num)
            else:
                print("Error: Incomplete data received.")
                break

        if irValues == received_numbers:
            print("Success! The arrays match.")
        else:
            print(f"Error: Arrays do not match. Received {len(received_numbers)} numbers.")

        ser.close()

    if action == "receive":
        ser = serial.Serial(port, 9600, timeout=2)
        ser.setDTR(False)
        time.sleep(1)
        ser.flushInput()
        ser.setDTR(True)
        time.sleep(2)
        ser.write("receive\n".encode('utf-8'))

        recieved_num = ser.readline()
        while not recieved_num:
            recieved_num = ser.readline()

        print(f"Received numbers: {recieved_num}")
        # formatedDate = ""
        ser.close()

except getopt.error as err:
    print(str(err))
