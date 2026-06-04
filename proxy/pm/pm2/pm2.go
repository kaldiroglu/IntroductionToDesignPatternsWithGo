// Package pm2 is Solution 2: a Proxy sorts out and delegates to the PM; the Citizen
// still knows it is talking to a Proxy.
package pm2

import "fmt"

type PM struct{}

func (p PM) Listen(problem string) {
	fmt.Println("PM: Listening to you.")
	p.resolve(problem)
}

func (p PM) FindJob(name string) {
	fmt.Println("PM: Don't ask me to find a job for you!")
}

func (p PM) resolve(problem string) {
	fmt.Println("PM: Please resolve this: " + problem)
}

type Proxy struct {
	pm PM
}

func NewProxy(pm PM) *Proxy {
	return &Proxy{pm: pm}
}

func (pr *Proxy) Listen(problem string) {
	fmt.Println("Proxy: Listening to you.")
	if pr.sortOut(problem) {
		pr.delegate(problem)
	}
}

func (pr *Proxy) FindJob(name string) {
	fmt.Println("Proxy: 'I'll find out what I can do for you!'")
}

func (pr *Proxy) delegate(problem string) {
	pr.pm.Listen(problem)
}

func (pr *Proxy) sortOut(problem string) bool {
	b := true
	// ...
	return b
}

type Citizen struct {
	name  string
	proxy *Proxy
}

func NewCitizen(name string, proxy *Proxy) *Citizen {
	return &Citizen{name: name, proxy: proxy}
}

func (c *Citizen) TellProblem() {
	c.proxy.Listen("The problem is ...")
}

func (c *Citizen) AskForJob() {
	c.proxy.FindJob(c.name)
}

func Run() {
	pm := PM{}
	proxy := NewProxy(pm)
	citizen := NewCitizen("John", proxy)
	citizen.TellProblem()
	citizen.AskForJob()
}
