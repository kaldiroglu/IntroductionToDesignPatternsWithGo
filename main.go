// Command intro runs the pm1/pm2/pm3 proxy demos and the db (telescoping-constructor)
// demo, mirroring the Java Mains. The complexObject example has no Main in the Java repo,
// so (like the original) it is library-only and is compiled by `go build ./...`.
package main

import (
	"fmt"

	"dev.kaldiroglu/dp/intro/problems/db"
	"dev.kaldiroglu/dp/intro/proxy/pm/pm1"
	"dev.kaldiroglu/dp/intro/proxy/pm/pm2"
	"dev.kaldiroglu/dp/intro/proxy/pm/pm3"
)

func main() {
	fmt.Println("===== Solution 1 (pm1): Citizen -> PM =====")
	pm1.Run()

	fmt.Println()
	fmt.Println("===== Solution 2 (pm2): Citizen -> Proxy -> PM =====")
	pm2.Run()

	fmt.Println()
	fmt.Println("===== Solution 3 (pm3): Citizen -> ProxyPM -> RealPM (served by PMSecretary) =====")
	pm3.Run()

	fmt.Println()
	fmt.Println("===== Problem: telescoping constructors / post-construction setters (db) =====")
	db.Run()
}
