# In order to compile Go code and create a binary executable file, you need to use
# the go build command. The generated executable is automatically named after the
# source code filename without the .go file extension.
go build hello-world.go
# go build supports the -o option, which allows you to change the filename and 
# the path of the generated executable file.
go build -o helloWorld hello-world.go

# The go run command builds the named Go package, which in this case is the main
# package implemented in a single file, creates a temporary executable file, executes
# that file, and deletes it once it is done—to our eyes, this looks like using a scripting
# language.
go run hello-world.go
