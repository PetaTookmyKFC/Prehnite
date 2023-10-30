
<p align="center" style="border-bottom: solid 1px; font-size: 1.2rem;">ALL INFORMATION IS SUBJECT TO CHANGE
</p>

All routes begin with `/api/`

* [List Depots](### List Depots)

|Route|Description|
|-|-|
| /depot/list | Returns an array of currently saved depots|
| /depot/create|Creates a new folder to contain the database, also updates the config|
| /depot/recalculate| Removes all containers saved within the depot and add all folders contained within the databaseDepot (Used if a database is added using external tools) |


## Error Responses
Errors returned from the server will be a html page containing the information about the error. JSON responses are not available at this time.

### Basic Responses
These method only returns a single word to confirm that the request was successful. As a result of this this param is used to switch between the formatted error page and a string response. The response will also be placed in a message property in the JSON
```JSON
{
	"Message" : "Wow it worked!"
}
```

### List Depots

<p align="center" style="border-bottom: solid 1px; font-size: 1.2rem;">THIS INFORMATION IS SUBJECT TO CHANGE</p>

Route : `/api/depot/list` 
Method : `Post`
Description: Returns a list of the current containers saved on the server
Params: 
1. B
	This is used to ask the server for a non formatted response.  This should be excluded from the URL entirely to ensure it isn't processed. 
	Response with B=true
	```JSON
	["ExampleDatabase", "Name2"] 
	```
	Response without B=true
```html
<option value="{{ index }}">
	ExampleDatabase
</option>
<option value="{{ index }}">
	Name2
</option>
```

### Depot Info

<p align="center" style="border-bottom: solid 1px; font-size: 1.2rem;">THIS INFORMATION IS SUBJECT TO CHANGE<br>THIS IS NOT ENABLED</p>


Route : `/api/depot/info`
Method : `POST`
Params : 
1. B
	This is used to ask the server for a non formatted response.  This should be excluded from the URL entirely to ensure it isn't processed. 
	Response with B=true
```json
{
	Name : "Example Name",
	Location: "C:/Example",
}
```

2. Depot
	This is used to inform the server what depot you would like information about.


### Create Container

<p align="center" style="border-bottom: solid 1px; font-size: 1.2rem;">THIS IS A BASIC RESPONCE API METHOD</p>

Route : `/api/depot/create`
Method : `POST`

Params : 
1. B
		This is used to ask the server for a non formatted response.  This should be excluded from the URL entirely to ensure it isn't processed.  This method only returns a single word to confirm that the request was successful. As a result of this this param is used to switch between the formatted error page and a string response. The response will also be placed in a message property in the JSON
	Response with B=true
```json
{
	Message : "Example Error - Error Example"
}
```

2. Name
	This is used to inform the server what the container should be called and saved as


### Recalculate Config 
<p align="center" style="border-bottom: solid 1px; font-size: 1.2rem;">THIS IS A BASIC RESPONCE API METHOD</p>
Route : `/api/depot/recalculate`
Method : `POST`
Params : 
1. B
	This is used to ask the server for a non formatted response.  This should be excluded from the URL entirely to ensure it isn't processed.  
	Response with B=true
```json
{ Message: "Recalculated" }
```
