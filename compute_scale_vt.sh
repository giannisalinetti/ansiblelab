#!/bin/bash

source /home/stack/stackrc

curl -o /home/stack/instackenv-onenode.json http://materials.example.com/instackenv-onenode.json

openstack baremetal import --json /home/stack/instackenv-onenode.json

openstack baremetal node manage compute1

openstack overcloud node introspect --all-manageable --provide

# Define compute capabilities
openstack baremetal node set compute1 --property "capabilities=profile:compute,boot_option:local"

# Fix for VT environment 
openstack baremetal node set compute1 --property memory_mb=6144

# Scale out the number of compute replicas
sed -i  's/^ComputeCount/ComputeCount: 2/' /home/stack/templates/cl210-environment/00-node-info.yaml

# Deploy the overcloud
openstack overcloud deploy --templates /home/stack/templates --environment-directory /home/stack/templates/cl210-environment


