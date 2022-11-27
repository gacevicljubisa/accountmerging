package merger_test

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
	"testing"

	am "github.com/gacevicljubisa/accountmerging/merger"
)

func TestMerge(t *testing.T) {
	testTable := []struct {
		name     string
		input    []am.Account
		expected []am.Person
	}{
		{
			name:     "empty",
			input:    []am.Account{},
			expected: []am.Person{},
		},
		{
			name: "accounts",
			input: []am.Account{
				{
					Application: json.Number("1"),
					Emails:      []am.Email{am.Email("a@gmail.com"), am.Email("b@gmail.com")},
					Name:        "A",
				},
				{
					Application: json.Number("1"),
					Emails:      []am.Email{am.Email("c@gmail.com"), am.Email("d@gmail.com")},
					Name:        "C",
				},
				{
					Application: json.Number("2"),
					Emails:      []am.Email{am.Email("a@yahoo.com")},
					Name:        "A",
				},
				{
					Application: json.Number("3"),
					Emails:      []am.Email{am.Email("a@yahoo.com"), am.Email("a@gmail.com")},
					Name:        "A",
				},
			},
			expected: []am.Person{
				{
					Applications: []string{"1", "2", "3"},
					Emails:       []string{"a@gmail.com", "b@gmail.com", "a@yahoo.com"},
					Name:         "A",
				},
				{
					Applications: []string{"1"},
					Emails:       []string{"c@gmail.com", "d@gmail.com"},
					Name:         "C",
				},
			},
		},
		{
			name: "accounts_dif_name",
			input: []am.Account{
				{
					Application: json.Number("1"),
					Emails:      []am.Email{am.Email("a@gmail.com"), am.Email("b@gmail.com")},
					Name:        "B",
				},
				{
					Application: json.Number("1"),
					Emails:      []am.Email{am.Email("c@gmail.com"), am.Email("d@gmail.com")},
					Name:        "C",
				},
				{
					Application: json.Number("2"),
					Emails:      []am.Email{am.Email("a@yahoo.com")},
					Name:        "C",
				},
				{
					Application: json.Number("3"),
					Emails:      []am.Email{am.Email("a@yahoo.com"), am.Email("a@gmail.com")},
					Name:        "A",
				},
			},
			expected: []am.Person{
				{
					Applications: []string{"1", "2", "3"},
					Emails:       []string{"a@gmail.com", "b@gmail.com", "a@yahoo.com"},
					Name:         "A",
				},
				{
					Applications: []string{"1"},
					Emails:       []string{"c@gmail.com", "d@gmail.com"},
					Name:         "C",
				},
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			merger := am.AccountMerger{}
			result := merger.Merge(test.input)

			// sort result and expected and then compare
			resultMap := getSortedPersonsMap(result)
			expectedMap := getSortedPersonsMap(test.expected)

			if !reflect.DeepEqual(expectedMap, resultMap) {
				t.Errorf("response expected: %v, got: %v", test.expected, result)
			}
		})
	}
}

func getSortedPersonsMap(persons []am.Person) (pm map[string]am.Person) {
	pm = make(map[string]am.Person, len(persons))
	for _, person := range persons {
		sort.Strings(person.Applications)
		sort.Strings(person.Emails)
		pm[strings.Join(person.Emails, ",")] = person
	}

	return pm
}
