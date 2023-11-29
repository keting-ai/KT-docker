package actions

import (
	"fmt"
	"io"
	"os"

	"github.com/ktdocker/container"
	log "github.com/sirupsen/logrus"
)

func LogContainer(containerName string) {
	dirURL := fmt.Sprintf(container.DefaultInfoLocation, containerName)
	logFileLocation := dirURL + container.ContainerLogFile
	file, err := os.Open(logFileLocation)

	defer file.Close()
	if err != nil {
		log.Errorf("Log container open file path: %s error: %v", logFileLocation, err)
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Errorf("Log container read file path: %s error: %v", logFileLocation, err)
		return
	}
	fmt.Fprint(os.Stdout, string(content))
}
