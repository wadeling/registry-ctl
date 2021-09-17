package registry

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

func StartRegistry() error {
	cmd := exec.Command("sh", "-c", "/bin/registry  serve /etc/docker/registry/config.yml ")
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	log.Debugf("Start regsitry...")
	if err := cmd.Run(); err != nil {
		log.Errorf("Fail to execute GC: %v, command err: %s", err, errBuf.String())
		return err
	}

	log.Debugf("Successful to start registry ...")
	return nil
}
