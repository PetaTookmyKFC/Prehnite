# Layout


LAYOUT

Depot  -< ( Container (multiple Databases / one folder) )
Container -< Database ( sqlite file )

## Depot

<!-- Depot is the controller at allow multiplle containers to be pressent. All databases within the same folder will be grouped into one container. Depot allows the addition of containers from other location. This could be a backup of a container allowing the databases to have the same name as they are in seperate containers. The limit for multiple containers containing the same name is 99, this can be chaged. -->
#### Functions


* [AddDBContainer](#####AddDBContainer) (Location `string`) `error`
* [ListContainers]() ( ) `[]string`
<!-- ##### AddDBContainer 
This created a new container and adds all `.db` files as databases into the container. The new Container is then inserted into the Depot. The error will be `nil` unless the number of containers with the same name exeeds the maximum. Other errors may be thown. Upon Starting the system `./DatabaseDepot` will be created and initilised as the first container ( if it doesn't exit. ) -->

## Container
<!-- This is the controller for the database. Sqlite files will be stored in this object.  -->

* [AddDatabase](######AddDatabase) ( Name `string`, Location `string` ) ( id `string`, error `error`  )
<!-- 
###### AddDatabase -->
