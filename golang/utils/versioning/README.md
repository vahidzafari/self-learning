# Versioning

One of the most difficult tasks is to automatically and uniquely version commandline utilities, especially when using a CI/CD system. This section presents a technique that uses a GitHub value to version a command-line utility on your local machine.

What we need to do is tell the Go linker that we are going to define the value of the VERSION variable. This happens with the help of the -ldflags flag, which stands for linker flags—this passes values to the cmd/link package, which allows us to change values in imported packages at build time. The -X value that is used requires a key/value pair, where the key is a variable name and the value is the value that we want to set for that key. In our case, the key has the main.Variable form because we change the value of a variable in the main package. As the name of the variable in gitVersion.go is VERSION, the key is main.VERSION.

But first, we need to decide on the GitHub value that we are going to use as the version string. The git rev-list HEAD command returns a full list of commits for the current repository from the latest to the oldest. We only need the last one—the most recent—which we can get using git rev-list -1 HEAD or git rev-list HEAD | head -1. So, we need to assign that value to an environment variable and pass that environment variable to the Go compiler. As this value changes each time you make a commit and you always want to have the latest value, you should reevaluate it each time you execute go build—this will be shown in a while.

```bash
export VERSION=$(git rev-list -1 HEAD)
go build -ldflags "-X main.VERSION=$VERSION" gitVersion.go
```
