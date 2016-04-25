### TCP Relay

Run this inside the GFW

`$ ./relay -l 0.0.0.0:12345 -r <remote>:12345 -w -logtostderr`

Run this outside the GFW

`$ ./relay -l 0.0.0.0:12345 -r <http_proxy_server>:<port>`

### Changelog

#### v0.0.1
* A working encrypted TCP relay in less than 100 lines of code

#### v0.1.0
* A monitor package showing how many TCP connections are open

### See Also

* https://github.com/jpillora/chisel
* https://github.com/eahydra/socks
