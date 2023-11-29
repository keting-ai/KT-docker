# ktdocker
ktdocker is a simple container implementation, imitating the docker.  
This project refers to code in the book "Write docker yourself".
This project isn't completed yet. I haven't implemented many features -- just for fun!  
The basic is: 
1. Using Namespace to isolate resources
2. Using Cgroup to limit the resources
  
Available features are:  
- Run an initial command in the specified container  
- Start run a container (an image, rigorously)  
- Execute a container (this means a container exists, just start run it)  
- Stop a container (stop running)  
- Remove a container (delete the container)
  
Not yet implemented:  
- Deamon process
- Registry
- Commit a container to an image (just need to tar... but too lazy to do that...)  
- Container network  
- Remote Repository
- And so on...

Uses busybox to create an image for this project.  
This project is only tested on Ubuntu Linux. Doesn't work for MacOS and Windows!
