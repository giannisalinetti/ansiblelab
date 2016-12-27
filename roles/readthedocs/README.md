readthedocs
=========

Readthedocs is a small role to install readthedocs.

Requirements
------------

Python 2.7 installed.

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
