package container

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// RunContainerInitProcess runs the initial process within a container.
func RunContainerInitProcess() error {
	cmdArray := readUserCommand()
	if len(cmdArray) == 0 {
		return fmt.Errorf("run container get user command error, cmdArray is empty")
	}

	setUpMount()

	path, err := exec.LookPath(cmdArray[0])
	if err != nil {
		log.Errorf("exec loop path error: %v", err)
		return err
	}
	log.Infof("find path %s", path)

	if err := syscall.Exec(path, cmdArray, os.Environ()); err != nil {
		log.Errorf("syscall exec error: %v", err)
		return err
	}
	return nil
}

// readUserCommand reads the command to be executed inside the container.
func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := io.ReadAll(pipe)
	if err != nil {
		log.Errorf("init read pipe error: %v", err)
		return nil
	}
	return strings.Fields(string(msg))
}

// setUpMount sets up the necessary filesystem mounts for the container.
func setUpMount() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Errorf("get current location error: %v", err)
		return
	}
	log.Infof("current location is %s", pwd)

	if err := pivotRoot(pwd); err != nil {
		log.Errorf("pivot root error: %v", err)
		return
	}

	// Mount proc and tmpfs.
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
}

// pivotRoot performs the pivot_root operation to change the root filesystem.
func pivotRoot(root string) error {
	// Bind mount root to itself - this is a prerequisite for pivot_root.
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("mount rootfs to itself error: %v", err)
	}

	pivotDir := filepath.Join(root, ".pivot_root")
	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return fmt.Errorf("pivot_root error: %v", err)
	}

	// Change the current working directory to the new root.
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chdir to / error: %v", err)
	}

	// Unmount and remove the temporary pivot directory.
	pivotDir = filepath.Join("/", ".pivot_root")
	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount pivot_root dir error: %v", err)
	}
	return os.Remove(pivotDir)
}
