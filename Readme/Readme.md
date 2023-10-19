
The entire program is stored within a single `depot`. This depot is in charge of holding references to all containers stored within the system. This depot is also used to create folders for new containers as well as save information about the containers to a config file. This config file is just the depot saved in json form. An example of this will be : 
```json
{
	"containerName" : {
		"Name" : string,
		"Location": string
	}
}
```

The config file only contains the name of the container including the name that is used to identify the container within the depot, and the location where the container is stored. As multiple containers can have the same name but stored in different locations, the Name within the container can be different from the name that is used to identify the container. 
