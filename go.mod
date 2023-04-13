module github.com/agnivade/modbug/v7

go 1.19

replace github.com/agnivade/modbug/public/v7 => ./public

require (
	github.com/agnivade/modbug/public/v7 v7.0.0
	github.com/agnivade/pgm v0.0.0-20210528073050-e2df0d9cb72d
)
