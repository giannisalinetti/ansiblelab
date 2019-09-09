# Rest API lookup
This example uses the url lookup plugin to collect data in json format from an
example REST API service that prints out a list of objects containing data
about employees.

After collecting the results in a variable we want to filter only employees
with a salary higher than 50k and print them in a file in yaml format.
This involves the usage of list concatenation over a loop.

To demonstrate performance metrics, both blockinfile and copy modules are
implemented, and they can be recalled with a different tag.

In this examples blockinfile is about twice faster than copy. Callback plugins
(timer and profile_tasks) have been used to profile the execution of the 
playbook.
