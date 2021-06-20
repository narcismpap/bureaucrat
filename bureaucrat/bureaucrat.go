// Bureaucr.at Coding Challenge
// Author: Narcis M. Pap - https://www.linkedin.com/in/narcismpap/
// London, Jun 2021
// github.com/narcismpap/bureaucrat

package main

import (
	"errors"
	"fmt"
)

type StaffReference string

type Staff struct {
	Name      string         `json:"name"`
	Ref       StaffReference `json:"ref"`
	ManagerOf []*Staff       `json:"manager_of"`
}

type DirectoryQuery struct {
	directory *Staff
	ref_one   StaffReference
	ref_two   StaffReference
}

func NewDirectoryQuery(dir *Staff, r1 StaffReference, r2 StaffReference) *DirectoryQuery {
	return &DirectoryQuery{
		directory: dir,
		ref_one:   r1,
		ref_two:   r2,
	}
}

func (d *DirectoryQuery) CommonManager() (StaffReference, error) {
	if d.ref_one == d.ref_two {
		return "", errors.New("ref_one and ref_two are the same person")
	}

	if d.directory == nil {
		return "", errors.New("directory is empty")
	}

	managersRefOne, err := d.GetManagers(d.ref_one)
	if err != nil {
		return "", err
	}

	managersRefTwo, err := d.GetManagers(d.ref_two)
	if err != nil {
		return "", err
	}

	for i := len(managersRefOne) - 1; i >= 0; i-- {
		for j := len(managersRefTwo) - 1; j >= 0; j-- {
			if managersRefOne[i] == managersRefTwo[j] {
				return managersRefOne[i], nil
			}
		}
	}

	return "", errors.New("no common manager found")
}

func (d *DirectoryQuery) GetManagers(ref StaffReference) (managers []StaffReference, err error) {
	managers_chan := make(chan StaffReference) // records entire path
	exists := false

	go func() {
		managers_chan <- d.directory.Ref
		exists = discoverManagers(d.directory, ref, managers_chan)
		close(managers_chan)
	}()

	for mg_ref := range managers_chan {
		managers = append(managers, mg_ref)
	}

	if !exists {
		return make([]StaffReference, 0), errors.New(fmt.Sprintf("staff %s not in directory", ref))
	}

	return managers, nil
}

func discoverManagers(d_slice *Staff, ref StaffReference, m_chan chan StaffReference) bool {
	if d_slice == nil {
		return false
	}

	if d_slice.Ref == ref {
		return true
	}

	for _, manager_of := range d_slice.ManagerOf {
		if discoverManagers(manager_of, ref, m_chan) {
			if manager_of.Ref != ref {
				m_chan <- manager_of.Ref
			}
			return true
		}
	}

	return false
}
