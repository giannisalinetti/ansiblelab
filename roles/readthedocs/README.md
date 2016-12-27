Role Name
=========

Readthedocs is a small role to install readthedocs.

Requirements
------------

Any pre-requisites that may not be covered by Ansible itself or the role should be mentioned here. For instance, if the role uses the EC2 module, it may be a good idea to mention in this section that the boto package is required.

Role Variables
--------------

Default variables:
  venv: /opt/rtd

Main vars:

CentOS/RedHat Variables:
pre_pkg:
  - python-devel
  - python2-pip
  - python-setuptools
  - python-virtualenv
  - libxml2-devel
  - libxslt-devel
  - zlib-devel

Dependencies
------------

Example Playbook
----------------

    - hosts: servers
      roles:
         - { role: gbs.readthedocs, venv: /opt/docs }

License
-------

GPL

Author Information
------------------

Gianni Salinetti
