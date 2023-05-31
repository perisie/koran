go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
go get github.com/volatiletech/sqlboiler/v4

# make sure the binary exists here
ls "$GOPATH/bin/sqlboiler"

# also make sure that $PATH points to $GOPATH/bin
sqlboiler --version
