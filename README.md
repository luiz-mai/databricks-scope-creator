# databricks-scope-creator

🛠 CLI tool used to create multiple scopes with multiple secrets on it and also apply any desired ACLs.

## Installing

First, make sure your `GOBIN` folder (usually `~/go/bin`) in your `PATH`. If not, you can add it by running:
```
export PATH=$PATH:~/go/bin
```

Then, simply run 
```
go install
```

## Running

To run the tool, you'll need to create a JSON file following the example of the `scopes.example.json` file inside this repository. The object will be composed by the scopes names as keys and their corresponding values are objects with `secrets` and `acls` as its properties. On `secrets`, include the `key-secret` pairs and on `acls`, include the `principal-permission` pairs.

With the file on hands, run the tool using:
```
databricks-scope-creator
```

and pass the (relative) file path as input when requested.

## Author

Luiz Mai <lffmai@gmail.com>

## License

MIT
