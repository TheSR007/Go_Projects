# Example Server
## Example Server in GO Using Fiber

### Features
- [X] HTTP(S)
- [X] WebSocket
- [X] Database(MongoDB)
- [X] UDP(Enet)

## Usage
Navigate to the this directory and
```cmd
go run .
```
For The **TCP and UDP Server <br>
for Enet Client**
```cmd
cd Enet_Client
go run
```
for the server and the enet client. <br> 
To build
```cmd
go build .
```
goto Enet_Client folder and
```cmd
go build
```
build Directory Also has **HTTP and HTTPS builds** for **Windows** which runs on **localhost 80 and 443** respectively **(need manual setup for /db endpoint else you will get error as string)**
**Enet_Client.exe** is also provided to checkout **UDP client** which **connects to the Server, and starts ping pong**
<br>
### **Note** 
- Make Sure The **MongoDB** is Setup Correctly for **Localhost** and you Have a database Named **DatabaseName** and a collection named **CollectionName**
- You can change them and rebuild as you need, these are just placeholder and only requred if you wanna use the prebuilt binaries
- Run The Prebuilds in [**Terminal**](https://www.microsoft.com/store/productId/9N0DX20HK701) for Better View
- Also the **Config of Server is in main.go**
## DOC
* `/`
	+ **GET**: "Server is running"
	+ **POST**: "Server is running"
* `/ws` (WebSocket)
	+ **Messages**
		- **ping**: Server will return a **pong**
* `/db` (Database)
	+ **GET**: Returns the MongoDB Database Documents (need to setup first)


