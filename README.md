# Cryptoplasm-Core

## Local setup
Follow these steps to configure the go project locally:

* Make sure GO `1.15` is installed!
* Create a new folder somewhere and make sure the `GOPATH` points to that directory (in GOLAND). This folder will contain all future go projects
* Clone this repository in `YOUR_FOLDER\src\github.com`
* The folder should look something like: `C:\go-code\src\github.com\Cryptoplasm-Core` after cloning this, with `GOPATH` pointing to `C:\go-code\`

## Build on Windows
To follow GO workspace structure recommendation, will use a `bin` folder (added to gitignore) where the executables should be built.
Just execute the script in the `scripts` folder from powershell: 

`powershell.exe -ExecutionPolicy Unrestricted .\build.ps1`

## Others
When adding new references in the code (or create new packages with new references) execute these 2 commands to vendor them
* `go mod tidy`
* `go mod vendor`