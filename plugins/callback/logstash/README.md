# Logstash callback plugin
This plugin must be configured using environment variables only. There is no
supported ini interface.

```
$ LOGSTASH_SERVER=localhost LOGSTASH_PORT=5000 ansible-playbook playbook.yml
```

This plugin requires the **python-logstash** module installed which must be 
installed using `pip3`. 
Since Ansible doesn't resolve the default `/usr/local/lib/python3.7/site-packages`
we are going to install it on `/usr/lib/python3.7/site-packages`. This is 
potentially dangerous and can create conflicts with rpm packages.

```
$ sudo pip3 install -t /usr/lib/python3.7/site-packages/ python3-logstash
```



