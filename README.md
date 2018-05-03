# VehicleFinder

API for retrieving vehicles

## Installation

Add this to your source code

```go
import "github.com/drum445/vehicleFinder"
```

Run in a terminal in your project root

```bash
$ cd src && go get
```

## Usage
Providing you have a local MongoDB instance

```bash
$ go build -o app && ./app
```
### APIs
#### Importing the CSV file into Mongo
```
POST: http://localhost:5000/api/vehicle/import
```

#### Retrieving all vehicles
```
GET: http://localhost:5000/api/vehicle
```
You can add the following URL params:<br>
- make
- shortModel
- longMondel
- trim
- derivative

#### Retrieve one vehicle with image
```
GET: http://localhost:5000/api/vehicle/:id
```

## Contributors

- [drum445](https://github.com/drum445) ed - creator, maintainer
