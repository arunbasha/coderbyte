import requests
import numpy as np
import pandas as pd

r = requests.get('https://coderbyte.com/api/challenges/json/age-counting')
res = r.json()['data']
print(res)
len = len(r.json()['data'])
#print (res)

#print (len)

val = res.split(", ")
# print(val)
ages = []
for i in val:
    # print (i)
    key = i.split("=")
    if key[0] == "age":
        ages.append(int(key[1]))
        print (key[1])
count = 0
for age in ages:
    if age >= 50:
        count = count + 1

print(count)

