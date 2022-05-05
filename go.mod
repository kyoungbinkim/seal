module github.com/kyoungbinkim/main

go 1.17

replace github.com/kyoungbinkim/seal => ./seal

replace github.com/filecoin-project/go-fil-filestore => ./go-fil-filestore

require github.com/kyoungbinkim/seal v0.0.0-00010101000000-000000000000

require github.com/filecoin-project/go-fil-filestore v0.0.0-20191202230242-40c6a5a2306c // indirect
