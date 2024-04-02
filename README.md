# Coffee shop
## Requirements
### Windows
1. wsl version 2, installed with:
    - go compiler
    - Bazel
    - gazelle => gazelle go package
1. bazel installed

### Debian  
1. Bazel
1. go compiler

## Update go mod file commands
1. Open a wsl terminal
1. Run the command:
    ```
        bazel run //:gazelle-update-repos 
    ```
## Run the project
```
    bazel run //...
```