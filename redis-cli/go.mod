module redis-cli

go 1.15

require (
	github.com/carmark/pseudo-terminal-go v0.0.0-20151106093136-5a48ae24c6f5
	redis-common v0.0.0
)

replace redis-common v0.0.0 => ../redis-common
