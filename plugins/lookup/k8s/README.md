# K8s lookup plugin

To run this demo the `openshift` python3 module must be installed:

```
$ sudo pip3 install -t /usr/lib/python3.7/site-packages/ openshift
```

## Demo playbook
The demo runs against a local minikube installation. To configure minikube
follow these instructions:
https://kubernetes.io/docs/tasks/tools/install-minikube/
  
The plugin uses the default kubeconfig file stored in *~/.kube/config*. To 
pass a custom kubeconfig, credentials or token, follow the plugin docs:

```
$ ansible-doc -t lookup k8s
```
