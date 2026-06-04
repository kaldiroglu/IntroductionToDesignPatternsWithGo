// Package pm3 is Solution 3: PM is an interface (the Subject). RealPM and ProxyPM
// implement it; PMSecretary is a factory that hands the Citizen a ProxyPM typed as PM.
package pm3

import "fmt"

type PM interface {
	Listen(problem string)
	FindJob(name string)
}

type RealPM struct{}

func (r RealPM) Listen(problem string) {
	fmt.Println("RealPM: Listening to you.")
	r.resolve(problem)
}

func (r RealPM) FindJob(name string) {
	fmt.Println("RealPM: Don't ask me to find a job for you!")
}

func (r RealPM) resolve(problem string) {
	fmt.Println("RealPM: Please resolve this: " + problem)
}

type ProxyPM struct {
	pm PM
}

func NewProxyPM(pm PM) *ProxyPM {
	return &ProxyPM{pm: pm}
}

func (pr *ProxyPM) Listen(problem string) {
	fmt.Println("Proxy: Listening to you.")
	if pr.sortOut(problem) {
		pr.delegate(problem)
	}
}

func (pr *ProxyPM) FindJob(name string) {
	fmt.Println("Proxy: 'I'll find out what I can do for you!'")
}

func (pr *ProxyPM) delegate(problem string) {
	pr.pm.Listen(problem)
}

func (pr *ProxyPM) sortOut(problem string) bool {
	b := true
	// ...
	return b
}

type PMSecretary struct {
	pm PM
}

func NewPMSecretary() *PMSecretary {
	return &PMSecretary{pm: NewProxyPM(RealPM{})}
}

// getMePM is unexported, mirroring the package-private getMePM() in Java; the Citizen,
// which lives in the same package, calls it.
func (s *PMSecretary) getMePM() PM {
	return s.pm
}

type Citizen struct {
	name string
	pm   PM
}

func NewCitizen(name string, secretary *PMSecretary) *Citizen {
	return &Citizen{name: name, pm: secretary.getMePM()}
}

func (c *Citizen) TellProblem() {
	c.pm.Listen("The problem is ...")
}

func (c *Citizen) AskForJob() {
	c.pm.FindJob(c.name)
}

func Run() {
	fmt.Println("Everything starts with a citizen coming to PM Secretary and asking for PM")
	secretary := NewPMSecretary()
	citizen := NewCitizen("John", secretary)
	citizen.TellProblem()
	citizen.AskForJob()
}
