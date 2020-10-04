# XKCD on Remarkable 

This program fetches a random [XKCD](https://xkcd.com) comic and puts it 
on your [Remarkable](https://remarkable.com/)'s suspend screen.

## Building
```sh
env GOOS=linux GOARCH=arm GOARM=7 go build -o cover.arm
```
## Install
```sh
scp cover.arm root@remarkable-local:/home/root/
scp cover.service root@remarkable-local:/etc/systemd/system/
scp cover.timer root@remarkable-local:/etc/systemd/system/
ssh root@remarkable-local "systemctl daemon-reload; systemctl enable cover; systemctl restart cover;"
ssh root@remarkable-local "systemctl enable cover.timer; systemctl restart cover.timer;"
```
