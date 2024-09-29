all:
	run-serv

run-serv:
	CONFIG_PATH=${CONFIG_PATH} . script/run-serv.sh

auto-start:
	. script/autostart.sh
