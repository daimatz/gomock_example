gomock_example
==============

```sh
go install code.google.com/p/gomock/gomock
go install code.google.com/p/gomock/mockgen
mkdir mock_gomock_example
mockgen github.com/daimatz/gomock_example TwitterService > mock_gomock_example/a.go
go test
```
