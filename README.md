# QGinion

## Usage

1. Create a workspace folder `gocode` in `~`(user home directory, In windows: C:\Users\<username> or In Mac \Users\<username> or In Linux \home\<username>) 
2. Inside `C:\Users\<username>\gocode` create folders `src\github.com\manjeshpv`
3. Open command line and Change present working directory to `C:\Users\<username>\gocode\src\github.com\manjeshpv`
4. Now Clone Repo
```sh
git clone https://github.com/manjeshpv/qginion.git
```
5. Go inside cloned directory `cd qginion`
(At this point open folder in webstorm and right click on `main.go` for more https://plugins.jetbrains.com/plugin/5047?pr=idea)
Update mongodb settings in `main.go`, you can create free mongo database in https://mongolab.com/
6. `go get`
7. `go build`
8. `qginion`
 
 
### Other Info

Workspace: `~\gocode` ie., `C:\Users\ManjeshV\gocode`
`~\gocode\src\githubc.com\manjeshpv\qginion`
In Windows Project Path : ``~\gocode\src\githubc.com\manjeshpv\qginion` ie.,`C:\Users\ManjeshV\gocode\src\githubc.com\manjeshpv\qginion`
In Mac : 

Environment Variables: `nano ~/.bashrc`

```sh
export PATH=$HOME/gocode/bin:$PATH
export GOPATH=$HOME/gocode
```
After `go build` executable created in `~/gocode/bin`

