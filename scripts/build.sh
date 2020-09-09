evaluate_home_dir(){
#Platform specific evaluation of Home DIR
  OS=$(uname)
  BASE_DIR=$(cd "$(pwd)/../.."; pwd)
  PATH_SEP=":"
  if [[ $OS = *"msys"* ]]; then
    CYGPATH=$(cygpath -w /)
    BASE_DIR=$CYGPATH$BASE_DIR
    PATH_SEP=";"
  fi
}

evaluate_home_dir
echo "BASE_DIR" $BASE_DIR
HOME_DIR=$BASE_DIR/restapi-boilerplate
export GOPATH=$GOPATH$PATH_SEP$HOME_DIR

echo "GOPATH" $GOPATH

#Build the code
go env
cd $HOME_DIR/src/app

#Run VET on the Server Code base
echo "Running VET on server code"
go vet $(go list ./... | grep -v generated) 2>&1

#Run FMT on the Server Code base
echo "Running fmt on server code"
go fmt $(go list ./... | grep -v generated)

#Run golint on the Server Code base
echo "Running golint on server code"
golint $(go list ./... | grep -v generated|swagger)

#Run go test on the Server Code base
#echo "Running go test on server code"
#go test $(go list ./... | grep -v generated)

#Run gosec on the Server Code base
#echo "Running gosec on server code"
#gosec -exclude=G104 ./...

#Build the Code
echo "Building server binary"
go install
