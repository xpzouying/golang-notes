from gevent import monkey; monkey.patch_all()
import gevent

def f(i):
    sum = 0
    for i in range(1000000000):
        x = i
        x += 1
        x *= 2
        x /= 3
        x -= 5

        sum += x

    print sum


gevent.joinall([
        gevent.spawn(f, 1),
        gevent.spawn(f, 2),
        gevent.spawn(f, 3),
])