# VehicleFinder

API for retrieving vehicles

## Usage
Providing you have a local MySQL instance

Backend:  
```bash
$ cd backend
$ go get
$ go build && ./backend
```
Frontend:  
```bash
$ cd frontend
$ npm build
$ npm run dev
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
- free (will search all columns)
- page (defaults to 1)

Example: http://localhost:5000/api/vehicle?make=mercedes&shortModel=B&trim=SE

#### Retrieve one vehicle and it's image
```
GET: http://localhost:5000/api/vehicle/:id
```

## Contributors

- [drum445](https://github.com/drum445) ed - creator, maintainer
