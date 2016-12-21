#!/usr/bin/python

# Basic Hello World Ansible module
# Tested with ansible/hacking/test-module utility

from ansible.module_utils.basic import *
import sys

def main():
    module = AnsibleModule(argument_spec={})
    response = {"hello": "world", "platform": sys.platform, "python_version": sys.version}
    module.exit_json(changed=False, meta=response)

if __name__ == '__main__':
  main()

