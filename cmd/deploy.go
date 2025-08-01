package main

import (
	"os/exec"
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

	// TODO: Construct terminal commands

	cmd := exec.Command("rsync", "-P", fullBinPath, deployUserAndHost+":~")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	cli.logger.Debug("Copyied binary to host.", "path", fullBinPath)

	cmd = exec.Command("rsync", "-P", fullUnitPath, deployUserAndHost+":~")
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	cli.logger.Debug("Copyied systemd service to host.", "path", fullUnitPath)

	return nil
}
