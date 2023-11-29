package subsystems

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

// Set sets the resource limits for the CPU subsystem.
func (s *CpuSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true)
	if err != nil {
		return fmt.Errorf("get cgroup path error: %w", err)
	}

	if res.CpuShare != "" {
		err = os.WriteFile(path.Join(subsysCgroupPath, "cpu.shares"), []byte(res.CpuShare), 0644)
		if err != nil {
			return fmt.Errorf("set cgroup cpu share fail: %w", err)
		}
	}
	return nil
}

// Remove removes the specified cgroup.
func (s *CpuSubSystem) Remove(cgroupPath string) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return fmt.Errorf("get cgroup path error: %w", err)
	}

	return os.RemoveAll(subsysCgroupPath)
}

// Apply adds a process to the cgroup.
func (s *CpuSubSystem) Apply(cgroupPath string, pid int) error {
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
func (s *CpuSubSystem) Name() string {
	return "cpu"
}
