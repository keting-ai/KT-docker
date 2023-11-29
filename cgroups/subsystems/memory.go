package subsystems

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

// Set sets the resource limits for the memory subsystem.
func (s *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true)
	if err != nil {
		return fmt.Errorf("get cgroup path error: %w", err)
	}

	if res.MemoryLimit != "" {
		err = os.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644)
		if err != nil {
			return fmt.Errorf("set cgroup memory fail: %w", err)
		}
	}
	return nil
}

// Remove removes the specified cgroup.
func (s *MemorySubSystem) Remove(cgroupPath string) error {
	subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false)
	if err != nil {
		return fmt.Errorf("get cgroup path error: %w", err)
	}

	return os.RemoveAll(subsysCgroupPath)
}

// Apply adds a process to the cgroup.
func (s *MemorySubSystem) Apply(cgroupPath string, pid int) error {
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
func (s *MemorySubSystem) Name() string {
	return "memory"
}
