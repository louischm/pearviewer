# PearViewer RoadMap

## Protocol
Using gRPC and protobuf. Bidirectional streaming RPC type.


### Methods

#### Client request
- Upload File
- Upload Dir
- Rename File
- Rename Dir
- Delete File
- Delete Dir
- Move File
- Move Dir
- Create Dir
- List Dir
- Login

#### Server request
- Download File
- Download Dir

### Message

- File
- Dir

## Server

Single server with conf file loaded at start. Put in Raspberry PI, Can be called by multiple client at once.
Transfer packet of data if needed (File exceeding a certain size).

## Client

Connect to a server chosen. Upload/Download one or multiple file at once. 

## GUI

- Search bar at top to search files/dir
- Filter by name/size/date
- ListDirWith scroll bar