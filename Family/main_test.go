package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name        string
		family      Family
		relation    Relationship
		newPerson   Person
		expectError bool
	}{
		{
			name: "Family without child",
			family: Family{
				Members: map[Relationship]Person{
					Father: {
						FirstName: "Ivan",
						LastName:  "Popov",
						Age:       52,
					},
				},
			},
			expectError: false,
			relation:    Child,
			newPerson: Person{
				FirstName: "Artem",
				LastName:  "Popov",
				Age:       0,
			},
		},
		{
			name: "Family with child",
			family: Family{
				Members: map[Relationship]Person{
					Child: {
						FirstName: "Ivan",
						LastName:  "Popov",
						Age:       52,
					},
				},
			},
			expectError: true,
			relation:    Child,
			newPerson: Person{
				FirstName: "Artem",
				LastName:  "Popov",
				Age:       0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			error := test.family.AddNew(test.relation, test.newPerson)
			//native
			if (error != nil) != test.expectError {
				t.Errorf("This family cannot accept new member in role %s with error %s", test.relation, error)
			}

			//testify
			if !test.expectError {
				require.NoError(t, error)
				assert.Contains(t, test.family.Members, test.relation)
				assert.EqualValues(t, test.family.Members[test.relation], test.newPerson, "Family does not contain new person %s", test.newPerson.FirstName)
				return
			}
			assert.Error(t, error)
		})
	}
}
