# Account Merging

**Account Merging** is tool used for merging the accounts to create a listing of people.

# Requirements
* Go 1.18 installed (*mandatory*)
* make (*optional*)

# Run unit tests
```
go test -v ./...
```
or
```
make test
```

# Execute
## from source code
NOTE: will use *accounts.json* from root of this project as file.
```
go run .
```
or
```
make run
```
## build to run
NOTE: **accountmerging** binary file is created in dist folder. Put *accounts.json* in the root folder of the executable file, or specify file path as first argument to any other file in the same format.
```
go build -o dist/accountmerging .
```
or
```
make binary
```
Execute next command from dist folder to use command line argument to specify path to json file:
* Linux\MacOS
```
./accountmerging ../accounts.json
```
* Windows
```
.\accountmerging.exe ..\accounts.json
```