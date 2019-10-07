#! /usr/bin/python3

from PIL import Image

data_payload = open("data_raw", "r")

data_converted = []

count = 0
for line in data_payload:
    for data in line.split(":"):
        data_converted.append(int(data, 16))
        count+=1

data_converted.reverse()

bytestr = bytes(data_converted)
image = Image.frombytes('1', (264,176), bytestr, 'raw')
image.save("out.bmp")
