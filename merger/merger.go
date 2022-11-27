package merger

import (
	"encoding/json"
)

type Account struct {
	Application json.Number
	Emails      []Email
	Name        Name
}

type Person struct {
	Applications []string
	Emails       []string
	Name         string
}

type (
	Email          string
	Name           string
	Application    string
	EmailSet       map[Email]struct{}
	ApplicationSet map[Application]struct{}
)

type Union struct {
	Emails       EmailSet
	Applications ApplicationSet
}

type AccountMerger struct {
	owners  map[Email]Name           // email to owner name
	apps    map[Email]ApplicationSet // email to the set of applications
	parents map[Email]Email          // email to a parent
	unions  map[Email]Union          // email as root to the set of emails and applications under the same account
}

func (am *AccountMerger) Merge(accounts []Account) (persons []Person) {
	am.owners = make(map[Email]Name)
	am.apps = make(map[Email]ApplicationSet)
	am.parents = make(map[Email]Email)
	am.unions = make(map[Email]Union)

	// initialization
	for _, account := range accounts {
		for _, email := range account.Emails {
			am.owners[email] = account.Name
			if _, ok := am.apps[email]; !ok {
				am.apps[email] = make(ApplicationSet)
			}
			am.apps[email].add(Application(account.Application))
			am.parents[email] = email
		}
	}

	// set parent for each
	for _, account := range accounts {
		rootEmail := account.Emails[0]
		for i := 1; i < len(account.Emails); i++ {
			am.parents[account.Emails[i]] = rootEmail
		}
	}

	// union all
	for _, account := range accounts {
		parentEmail := am.parents[account.Emails[0]]
		if _, ok := am.unions[parentEmail]; !ok {
			am.unions[parentEmail] = Union{
				Emails:       make(EmailSet),
				Applications: make(ApplicationSet),
			}
		}
		for i := 0; i < len(account.Emails); i++ {
			am.unions[parentEmail].Emails.add(account.Emails[i])
			for application := range am.apps[account.Emails[i]] {
				am.unions[parentEmail].Applications.add(application)
			}
		}
	}

	// create slice of persons
	persons = make([]Person, 0, len(am.unions))
	for parent := range am.unions {

		apps := make([]string, 0, len(am.unions[parent].Applications))
		for app := range am.unions[parent].Applications {
			apps = append(apps, string(app))
		}

		emails := make([]string, 0, len(am.unions[parent].Emails))
		for email := range am.unions[parent].Emails {
			emails = append(emails, string(email))
		}

		persons = append(persons, Person{
			Applications: apps,
			Emails:       emails,
			Name:         string(am.owners[parent]),
		})
	}

	return persons
}

func (es EmailSet) add(email Email) {
	es[email] = struct{}{}
}

func (as ApplicationSet) add(application Application) {
	as[application] = struct{}{}
}
