module github.com/kyoungbinkim/main

go 1.17

replace github.com/kyoungbinkim/seal/seal => ./seal

replace github.com/kyoungbinkim/seal/crypto => ./crypto

replace github.com/kyoungbinkim/seal/filestore => ./filestore

replace github.com/filecoin-project/go-fil-filestore => ./go-fil-filestore

require (
	github.com/kyoungbinkim/seal/crypto v0.0.0-00010101000000-000000000000
	github.com/kyoungbinkim/seal/filestore v0.0.0-20220505074844-04db798935c4
	github.com/kyoungbinkim/seal/seal v0.0.0-00010101000000-000000000000
)

require github.com/filecoin-project/go-fil-filestore v0.0.0-00010101000000-000000000000 // indirect
