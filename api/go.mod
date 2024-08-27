module github.com/cr4shed/did-i-see-it/api

go 1.23.0

replace github.com/cr4shed/did-i-see-it/data => ../data

require (
	github.com/cr4shed/did-i-see-it/data v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
)
