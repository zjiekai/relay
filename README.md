### TCP Relay

Run this inside the GFW

`$ ./relay -l 0.0.0.0:12345 -r <remote>:12345`


Run this outside the GFW

`$ ./relay -l 0.0.0.0:12345 -r <http_proxy_server>:<port>`

### Changelog

#### 0.0.0
* A working encrypted TCP relay in less than 100 lines of code

#### 0.0.1
* WIP

### See Also

* https://github.com/jpillora/chisel
* https://github.com/eahydra/socks
