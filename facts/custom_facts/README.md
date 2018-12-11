#Course DO407 - Automation with Ansible

### Creazione di custom fact dinamici
Questo esempio mostra come sia possibile generare dinamicamente dei custom facts.
E' uno scenario molto interessante per gli studenti che hanno bisogno di collezionare
informazioni dinamiche sulle macchine oltre ai default fact di Ansible.

- Il file **selinux_status.fact** verifica se SELinux è abilitato sulla macchina.
- La directory **users_list** contiene due esempi per raccogliere gli utenti locali in un output JSON.
Gli esempi mostrano come si possa produrre lo stesso risultato in due linguaggi differenti, **Python** e **Go**.

E' necessario copiare il file in */etc/ansible/facts.d* su tutti i nodi da gestire e
assicurarsi che esso sia eseguibile. E' **importante** che il file mantenga l'estensione **.fact**.

