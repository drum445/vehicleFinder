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
#### Importing the CSV file into MySQL
```
POST: http://localhost:5000/api/vehicle
```

#### Retrieve all available vehicles
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

Example: http://localhost:5000/api/vehicle?make=mercedes&shortModel=B Class&trim=SE

#### Retrieve one vehicle and it's image
```
GET: http://localhost:5000/api/vehicle/:id
```

## Contributors

- [drum445](https://github.com/drum445) ed - creator, maintainer
