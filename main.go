package main

import (
	_ "TMail/config"
	_ "TMail/domain"
	_ "TMail/repo"
	_ "TMail/smtp"
)

func main() {
	select {}
}
