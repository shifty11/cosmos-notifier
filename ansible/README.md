# DaoDao-Notifier docker-compose

Sets up everything to run daodao-notifer.

## Step 1: Set up your own inventory file

Copy inventory file, and make your edits.

```bash
cp samples/inventory.sample inventory
```

## Step 2: Run main playbook to set up a fresh project

```bash
ansible-playbook -i inventory main.yml
```

## Step 3: Update

```bash
ansible-playbook -i inventory main_update.yml
```