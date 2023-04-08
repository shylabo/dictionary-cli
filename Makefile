### 1.Build this app first.
build:
	go build .

### 2.Create a database(initialize).
init:
	./dictionary-cli init

### 3.Add a note.
new:
	./dictionary-cli note new

### 4.Show all the lists you registered.
list:
	./dictionary-cli note list
