package logger

import (
    "log"
)

func Log( msg string ) {
    log.Printf( "DEPOSITREG: %s", msg )
}