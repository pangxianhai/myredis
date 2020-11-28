module redis-server

go 1.15

require (
	redis-common v0.0.0
)

replace (
	redis-common v0.0.0 => ../redis-common
)
