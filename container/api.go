package container

var (
	RUNNING             string = "running"
	STOP                string = "stopped"
	Exit                string = "exited"
	DefaultInfoLocation string = "/var/run/ktdocker/%s/"
	ConfigName          string = "config.json"
	ContainerLogFile    string = "container.log"
	RootUrl             string = "/root"
	MntUrl              string = "/root/mnt/%s"
	WriteLayerUrl       string = "/root/writeLayer/%s"
)

// ContainerInfo holds information about a container.
type ContainerInfo struct {
	PID         string   `json:"pid"`         // PID of the container's init process on the host
	ID          string   `json:"id"`          // Container ID
	Name        string   `json:"name"`        // Container name
	Command     string   `json:"command"`     // Command running in the container
	CreatedTime string   `json:"createTime"`  // Time when the container was created
	Status      string   `json:"status"`      // Current status of the container
	Volume      string   `json:"volume"`      // Data volume of the container
	PortMapping []string `json:"portmapping"` // Port mappings
}
