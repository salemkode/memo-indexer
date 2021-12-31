module github.com/memocash/index

replace github.com/gcash/bchd => ../../../pkg/github.com/gcash/bchd

go 1.16

require (
	github.com/99designs/gqlgen v0.14.0
	github.com/gcash/bchd v0.17.1
	github.com/jchavannes/bchutil v1.1.3
	github.com/jchavannes/btcd v1.1.3
	github.com/jchavannes/btclog v1.1.0
	github.com/jchavannes/btcutil v1.1.3
	github.com/jchavannes/go-mnemonic v0.0.0-20191017214729-76f026914b65
	github.com/jchavannes/jgo v0.0.0-20211112043704-31caacec985a
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/syndtr/goleveldb v1.0.0
	github.com/tyler-smith/go-bip32 v1.0.0
	github.com/vektah/gqlparser/v2 v2.2.0
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)
