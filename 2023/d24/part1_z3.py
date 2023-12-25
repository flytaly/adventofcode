from datetime import datetime
import sys
import z3

is_test_data = len(sys.argv) <= 1

filename = sys.argv[1] if len(sys.argv) > 1 else "input.txt"
lines = open(filename, "r").read().splitlines()

hailstones = [tuple(map(int, l.replace(" @", ",").split(", "))) for l in lines]

min_val = 200000000000000
max_val = 400000000000000

if is_test_data:
    min_val, max_val = 7, 27

count = 0

ts = datetime.now()


s = z3.Solver()
x, y = z3.Reals("x y")
s.add(x >= min_val, x <= max_val, y >= min_val, y <= max_val)

for i, hs1 in enumerate(hailstones):
    for hs2 in hailstones[:i]:
        s.push()

        for px, py, _, vx, vy, _ in [hs1, hs2]:
            s.add(vy * (x - px) == vx * (y - py))
            s.add(x > px) if vx > 0 else s.add(x < px)
            s.add(y > py) if vy > 0 else s.add(y < py)

        if s.check() == z3.sat:
            # print(s.model())
            count += 1

        s.pop()


print("=>", count, str((datetime.now() - ts).seconds) + "s")
