package progress

// Daemon interface has a standard set of methods/commands
type Daemon interface {

	// Install the service into the system
	Install() error

	// Remove the service and all corresponding files from the system
	Remove() error

	// Start the service
	Start() error

	// Restart the service
	Restart() error

	// Stop the service
	Stop() error

	// Status - check the service status
	Status() (interface{}, error)

	SetTask([]string) error
}

// New - Create a new daemon
//
// name: name of the service
//
// description: any explanation, what is the service, its purpose
// func NewDaemon() Daemon {
// 	return newDaemon()
// }
