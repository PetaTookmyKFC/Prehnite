
All routes begin with `/api/`

* [List Depots](### List Depots)

|Route|Description|
|-|-|
| /depot/list | Returns an array of currently saved depots|


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

