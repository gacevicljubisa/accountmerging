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
```
go run .
```
or
```
make run
```
## build to run as binary
NOTE: **accountmerging** binary file is created in dist folder. Put accounts.json in the root folder of the executable file.
```
go build -o dist/accountmerging .
```
or
```
make binary
```
