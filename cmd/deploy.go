package main

import (
	"fmt"
	"path/filepath"
)

func (cli *CLI) Deploy() error {
	// Terminal example:
	// rsync -P ./bin/${SERVICE} ${user}@${PRODUCTION_HOST_IP}:~
	// rsync -P ./remote/production/${SERVICE}.service ${user}@${PRODUCTION_HOST_IP}:~
	// ssh -t ${user}@${PRODUCTION_HOST_IP} '\
	// 	sudo mv ~/${SERVICE}.service /etc/systemd/system/ \
	// 	&& sudo systemctl enable ${SERVICE} \
	// 	&& sudo systemctl restart ${SERVICE} \
	// '

	deployUserAndHost := cli.config.DeployUser + "@" + cli.config.Host
	fullBinPath := filepath.Join(cli.config.BinaryDir, cli.config.ServiceName)
	fullUnitPath := filepath.Join(cli.config.UnitDir, cli.config.ServiceName+".service")

	fmt.Println(fullUnitPath)      // remote/production/blitz.service
	fmt.Println(fullBinPath)       // bin/blitz
	fmt.Println(deployUserAndHost) // deploy_user@host

	// TODO: Construct terminal commands

	return nil
}
