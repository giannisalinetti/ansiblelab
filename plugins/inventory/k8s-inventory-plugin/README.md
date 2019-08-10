# K8s inventory plugin

To have this plugin work install the `openshift` module under the default
library path:

```
$ sudo pip3 install -t /usr/lib/python3.7/site-packages/ openshift
```

Then call a simple hostlist:
```
$ ansible all -i k8s.yml --list-hosts
```

Kubernetes namespaces will the used to generate Ansible groups as long as 
cluster labels and nodes.

