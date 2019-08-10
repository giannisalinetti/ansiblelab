# Memcached cache plugin
The easiest way to demonstrate this plugin is to run a memcached container:

```
$ sudo podman run -d -p 6379:6379 docker.io/library/redis
```

The container maps to the host port 0.0.0.0:6379. The cache plugin connects to
127.0.0.1:6379 by default.

## Python dependecies
To run this example the **redis** Python module must be installed. 
It can be installed using **pip** or with yum/dnf:

```
$ sudo yum install -y python3-redis
```

