# hotpocket 
This is a hot reload tooling for go

## Usage

 First need to have a json file in the root of your project. Name it as a `hotpocket.json`. It Should look somewhat like this
```
{
	"Command":"code", // here you specify command you wnat to execute on each reload
	"Arguments":[".."], // here you specify the arguments for that command
	"ExceptionFiles":[".md", ".gitignore"] // here you specify file extensions you dont want to listen for
}
```

After that you have to start your project with the hotpocket. ❗ YOU HAVE TO START IT IN THE ROOT OF YOUR PROJECT ❗
```
hotpocket
```

## Install

If you have go installed just run the followint command
```
go install github.com/rasulob-emirlan/hotpocket
```

Or you can compile it on your own. After you run the command bellow you will have the binary in ./bin folder
```bash
make build
```
