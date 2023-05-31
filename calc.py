m = 0.8325
c = 390.0
n = 0.018
T = 282
d1 = 2.955
r1 = 39.5875
d = 20.015
r = 80.05455

import math as math

k = (m * c * n * d1 * (r + 2*d)) / (2 * math.pi * r1 * r1 * T * (r+d))

km = (c * n * d1 * (r + 2*d)) / (2 * math.pi * r1 * r1 * T * (r+d))
nm = 0.0005

kc = (m * n * d1 * (r + 2*d)) / (2 * math.pi * r1 * r1 * T * (r+d))
nc = 5

kn = (m * c * d1 * (r + 2*d)) / (2 * math.pi * r1 * r1 * T * (r+d))
nn = 0.003277050458

kd1 = (m * c * n * (r + 2*d)) / (2 * math.pi * r1 * r1 * T * (r+d)) 
nd1 = 0.05

dr = 1.5451484416 * math.pow(10, -16)

dd = 2.47190935489 * math.pow(10,  -15)

dT = 7.0039093201 * math.pow(10, -16)

dr1 = 1.42161423849 * math.pow(10, -13) 


sum = (km * nm) ** 2 + (kc * nc) ** 2 + (kn * nn) ** 2 + (kd1 * nd1) ** 2 + dr * 0.05 ** 2 + dd * 0.05 ** 2 + dT * 1 ** 2 + dr1 * 0.05 ** 2

print(math.sqrt(sum))