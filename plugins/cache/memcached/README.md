# Memcached cache plugin
The easiest way to demonstrate this plugin is to run a memcached container:

```
$ sudo podman run -d -p 11211:11211 docker.io/library/memcached
```

The container maps to the host port 0.0.0.0:11211. The cache plugin connects to
127.0.0.1:11211 by default.

## Python dependecies
To run this example the **memcached** Python module must be installed. 
It can be installed using **pip** or with yum/dnf:

```
$ sudo yum install -y python3-memcached
```
