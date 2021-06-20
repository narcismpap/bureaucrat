// Bureaucr.at Coding Challenge
// Author: Narcis M. Pap - https://www.linkedin.com/in/narcismpap/
// London, Jun 2021
// github.com/narcismpap/bureaucrat

package main

import (
	"errors"
	"testing"
)

var gameOfThronesTests = []*TestCases{
	{"Direct", "GoT-005", "GoT-007", "GoT-004", nil},
	{"Indirect", "GoT-003", "GoT-002", "GoT-001", nil},
	{"ErrorLeft", "GoT-XXX", "GoT-002", "", errors.New("staff GoT-XXX not in directory")},
	{"ErrorRight", "GoT-003", "GoT-YYY", "", errors.New("staff GoT-YYY not in directory")},
	{"ErrorSame", "GoT-003", "GoT-003", "", errors.New("ref_one and ref_two are the same person")},
}

var gameOfThronesDirectory = &Staff{
	Ref:  "GoT-001",
	Name: "Daenerys Targaryen",
	ManagerOf: []*Staff{
		{
			Ref:  "GoT-002",
			Name: "Jon Snow",
			ManagerOf: []*Staff{
				{
					Ref:  "GoT-003",
					Name: "Samwell Tarly",
				},
			},
		},

		{
			Ref:  "GoT-004",
			Name: "Cersei Lannister",
			ManagerOf: []*Staff{
				{
					Ref:  "GoT-005",
					Name: "Jaime Lannister",
				},
				{
					Ref:  "GoT-006",
					Name: "Robert Baratheon",
				},
				{
					Ref:  "GoT-007",
					Name: "Gregor Clegane",
				},
			},
		},
	},
}

type TestCases struct {
	name   string
	ref1   StaffReference
	ref2   StaffReference
	expect StaffReference
	err    error
}

func TestClosestManager(t *testing.T) {
	for _, t_case := range gameOfThronesTests {
		d := NewDirectoryQuery(gameOfThronesDirectory, t_case.ref1, t_case.ref2)
		res, err := d.CommonManager()

		if res != t_case.expect {
			t.Errorf("BadResult(%s): exp [%s] got [%s]", t_case.name, t_case.expect, res)
		}

		if err == nil {
			if t_case.err != nil {
				t.Errorf("ExpectedError(%s): [%s]", t_case.name, t_case.err.Error())
			}
		} else {
			if t_case.err == nil {
				t.Errorf("UnexpectedError(%s): [%v]", t_case.name, err)
			} else if t_case.err.Error() != err.Error() {
				t.Errorf("MismatchedError(%s): exp [%v], got [%v]", t_case.name, t_case.err, err)
			}
		}
	}
}
