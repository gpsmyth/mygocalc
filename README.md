# Using Cobra CLI

## Installing Cobra
`go get -u github.com/spf13/cobra/cobra`
- You know $GOPATH
- Create your project directory outside of $GOPATH and cd into it
- To initialise your Cobra CLI
  - `cobra init --pkg-name mygocalc`

## Editing your Cobra
- In `root.go` file go to `var rootCmd` and have an entry like
```
Run: func(cmd *cobra.Command, args []string) { fmt.Println("Hello CLI")},
```
- The Cobra `Run` command invokes action inside the braces via its own `Command` structure setup using its `Run` method

## Installing and testing
- `go install mygocalc`
- `go run mygocalc`


## Adding numbers
- `cobra add add` add.go file is added
- Add some dubg lines
- `go install mygocalc`
- `go run mygocalc add` and observe output from add.go debug print statements

## Adding number ranges
- Exercise in string slicing, replacing and conversations from []string to strings or []ints
- `go run mygocalc add  --range "8,10"` or
- `go run mygocalc add  --range "8:10"` or
- `go run mygocalc add  --range 8:10` or 

## Typical output
```
go run mygocalc add  --range "8:10"
Range Values [8:10] type value []string
fstatus is  false istatus is false rstatus is true
addRange function 8 10
range from 8 to 10
sum is 27
str: 8:10 string
[8 10]result - 18
```

## Referrences
- https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177
- https://levelup.gitconnected.com/exploring-go-packages-cobra-fce6c4e331d6
- https://github.com/spf13/cobra/blob/master/cobra/README.md

## Debug mode
- in `launch.json` my configuration is
```
"configurations": [
    {
      "name": "Launch test function",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/main.go",
      "env": {
        "GO11MODULE": "on"
      },
      "args": [
        "add",
        "--range",
        "8:10"
      ]
    }
  ]
```