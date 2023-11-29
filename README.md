# ktdocker
ktdocker is a simple container implementation, imitating the docker
This project refers to code in the book "Write docker yourself".
This project isn't completed yet. I haven't implemented the image and network features like docker has. Also, no remote repository.
The basic is: 
1. Using Namespace to isolate resources
2. Using Cgroup to limit the resources
Available features are: 
	Init a container
	Run a container
	Execute a container
	Stop a container
	Remove a container
Not implemented:
  Commit a container to an image
  Container network
  Remote Repository
