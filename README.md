# VehicleFinder

API for retrieving vehicles

## Installation
Run in a terminal in your project root

```bash
$ go get
```

## Usage
Providing you have a local MySQL instance

```bash
$ go build && ./vehicleFinder
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
- page (defaults to 1)

#### Retrieve one vehicle with image
```
GET: http://localhost:5000/api/vehicle/:id
```

## Contributors

- [drum445](https://github.com/drum445) ed - creator, maintainer
