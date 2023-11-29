package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// FindCgroupMountpoint locates the mountpoint for a given cgroup subsystem.
func FindCgroupMountpoint(subsystem string) string {
	file, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		fmt.Printf("Error opening /proc/self/mountinfo: %v\n", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading /proc/self/mountinfo: %v\n", err)
		return ""
	}

	return ""
}

// GetCgroupPath combines the system's cgroup root directory with the operating cgroup directory.
func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountpoint(subsystem)
	if cgroupRoot == "" {
		return "", fmt.Errorf("failed to find cgroup mountpoint for subsystem %v", subsystem)
	}

	fullPath := path.Join(cgroupRoot, cgroupPath)
	_, err := os.Stat(fullPath)

	if err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(fullPath, 0755); err != nil {
				return "", fmt.Errorf("error creating cgroup at %v: %w", fullPath, err)
			}
		}
		return fullPath, nil
	}

	return "", fmt.Errorf("error checking cgroup path %v: %w", fullPath, err)
}
