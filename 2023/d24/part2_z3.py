import sys
import z3

filename = sys.argv[1] if len(sys.argv) > 1 else "input.txt"
lines = open(filename, "r").read().splitlines()

hailstones = [tuple(map(int, l.replace(" @", ",").split(", "))) for l in lines]

"""
If the rock's position is (x, y, z) and velocity is (vx, vy, vz),
then it will hit the hailstone (xh, yh, zh) at
    xh + t*vxh = x + t*vx => t = (x-xh)/(vxh-vx)
    ... => t = (y-yh)/(vyh-vy)
    ... => t = (z-zh)/(vzh-vz)
"""

s = z3.Solver()
x, y, z, vx, vy, vz = z3.Ints("x y, z, vx, vy, vz")

equations  = []
for xh, yh, zh, vxh, vyh, vzh in hailstones:
    s.add((x-xh)*(vyh-vy)==(y-yh)*(vxh-vx))
    s.add((y-yh)*(vzh-vz)==(z-zh)*(vyh-vy))

if s.check() == z3.sat:
    result = s.model().eval(x+y+z)
    print("=>", result)
