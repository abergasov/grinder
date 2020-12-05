dev start 
```shell script
bash dev.sh
```

## Front 4 dev
```shell
cd front && npm install && npm run dev
```

## DB & Migrations 
Docker db start 
```shell
bash dev.sh
```

Install dependencies 
```shell
cd helpers && composer install && touch phinx.yml
```
Set db config 
```yaml
paths:
  migrations: '%%PHINX_CONFIG_DIR%%/db/migrations'
  seeds: '%%PHINX_CONFIG_DIR%%/database/seeds'

environments:
  default_migration_table: phinxlog
  default_database: main
  main:
    adapter: mysql
    host: 127.0.0.1
    name: grinder
    user: 2ATCrMn2E2xhp43YL2ge
    pass: pbBEAVndCEVwetYl2wlkWP0qwFFDXFv2Jc
    port: 3019
    charset: utf8

version_order: creation
```
run init migrations
```shell
php app.php migrate
```


create new migration
```shell script
php app.php create UsersTable
```

apply migration
```shell script
cd helpers && php app.php migrate && cd ..
```