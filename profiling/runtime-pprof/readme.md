
go run .\main.go
go tool pprof -http=:8090 .\mutex_profile.prof
