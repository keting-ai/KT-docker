package subsystems

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

// Set sets the resource limits for the CPU set subsystem.
func (s *CpusetSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true)
	if err != nil {
		return fmt.Errorf("get cgroup path error: %w", err)
	}

	if res.CpuSet != "" {
		err = os.WriteFile(path.Join(subsysCgroupPath, "cpuset.cpus"), []byte(res.CpuSet), 0644)
		if err != nil {
			return fmt.Errorf("set cgroup cpuset fail: %w", err)
		}
	}
	return nil
}

// Remove removes the specified cgroup.
func (s *CpusetSubSystem) Remove(cgroupPath string) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return fmt.Errorf("get cgroup path error: %w", err)
	}

	return os.RemoveAll(subsysCgroupPath)
}

// Apply adds a process to the cgroup.
func (s *CpusetSubSystem) Apply(cgroupPath string, pid int) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return fmt.Errorf("get cgroup %s error: %w", cgroupPath, err)
	}

	err = os.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return fmt.Errorf("set cgroup proc fail: %w", err)
	}
	return nil
}

// Name returns the name of the subsystem.
func (s *CpusetSubSystem) Name() string {
	return "cpuset"
}
