#Course DO407 - Automation with Ansible

### Dynamic Ansible Facts
The examples in this folder show how to create dynamic custom facts. This approach
is very useful to integrate some extra (dynamically collected) informations from
the host.

- The file **selinux_status.fact** checks if SELinux is enabled on the machine.
- The directory **users*list** shows a couple of examples to collect local 
  users informations and output them in JSON format. The examples are written
  in **Golang** and **Python**.

The custom facts must be copied in the `/etc/ansible/facts.d` folder and the
executable bit must be set.
It is **mandatory** to keep the *.fact* extension.
