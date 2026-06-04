// Package pm1 is Solution 1: the Citizen talks to a concrete PM that listens and
// resolves the problem itself.
package pm1

import "fmt"

type PM struct{}

func (p PM) Listen(problem string) {
	fmt.Println("PM: Listening to you.")
	if p.sortOut(problem) {
		p.resolve(problem)
	}
}

func (p PM) FindJob(name string) {
	fmt.Println("PM: Don't ask me to find a job for you!")
}

func (p PM) sortOut(problem string) bool {
	b := true
	// ...
	return b
}

func (p PM) resolve(problem string) {
	fmt.Println("PM: Please resolve this: " + problem)
}

type Citizen struct {
	name string
	pm   PM
}

func NewCitizen(name string, pm PM) *Citizen {
	return &Citizen{name: name, pm: pm}
}

func (c *Citizen) TellProblem() {
	c.pm.Listen("The problem is ...")
}

func (c *Citizen) AskForJob() {
	c.pm.FindJob(c.name)
}

func Run() {
	pm := PM{}
	citizen := NewCitizen("John", pm)
	citizen.TellProblem()
	citizen.AskForJob()
}
