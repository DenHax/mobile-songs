all:
	run-serv

run-serv:
	CONFIG_PATH=${CONFIG_PATH} . script/run-serv.sh

run-compose:
	CONFIG_PATH=${CONFIG_PATH} . script/deploy-start-compose.sh

auto-start:
	. script/autostart.sh
	@$(MAKE) run-compose
