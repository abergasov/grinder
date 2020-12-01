dev start 
```shell script
bash dev.sh
```

new migration
```shell script
cd helpers && php app.php create UsersTable && cd ..
```

create migration
```shell script
cd helpers && php app.php migrate && cd ..
```