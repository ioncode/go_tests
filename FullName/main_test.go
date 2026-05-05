package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		value    User
		expected string
	}{
		{
			name: "Misha Popov",
			value: User{
				FirstName: "Misha",
				LastName:  "Popov",
			},
			expected: "Misha Popov",
		},
		{
			name: "Complex name",
			value: User{
				FirstName: "Wölfe Schlegel Steinhausen Bergerdorff vor altern waren gewissenhaft Schäfers wessen Schafe waren wohl Gepflege und Sorgfältigkeit beschützen von Angreifen durch ihr raubgierig Feinde welche vor altern zwölftausend Jahres vor an die Erscheinen wan der erste Erdemensch der Raumschiff gebrauch Licht als sein Ursprung von Kraft gestart sein lange Fahrt hin zwischen sternartig Raum auf der Suche nach die Stern welche gehabt bewohnbar Planeten Kreise drehen sich und wohin der neu Rasse von verständig Menschlichkeit konnte fortplanzen und sich erfreuen an lebenslänglich Freude und Ruhe mit nicht ein Furcht vor Angreifen von anderer intelligent Geschöpfs von hin zwischen sternartig Raum.",
				LastName:  "",
			},
			expected: "Wölfe Schlegel Steinhausen Bergerdorff vor altern waren gewissenhaft Schäfers wessen Schafe waren wohl Gepflege und Sorgfältigkeit beschützen von Angreifen durch ihr raubgierig Feinde welche vor altern zwölftausend Jahres vor an die Erscheinen wan der erste Erdemensch der Raumschiff gebrauch Licht als sein Ursprung von Kraft gestart sein lange Fahrt hin zwischen sternartig Raum auf der Suche nach die Stern welche gehabt bewohnbar Planeten Kreise drehen sich und wohin der neu Rasse von verständig Menschlichkeit konnte fortplanzen und sich erfreuen an lebenslänglich Freude und Ruhe mit nicht ein Furcht vor Angreifen von anderer intelligent Geschöpfs von hin zwischen sternartig Raum. ",
		},
		{
			name: "Noname",
			value: User{
				FirstName: "",
				LastName:  "",
			},
			expected: " ",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.value.FullName())
			if actual := test.value.FullName(); actual != test.expected {
				t.Errorf("For user %s expected result %s not equal with actual %s in case %s", test.value.LastName, test.expected, actual, test.name)
			}
		})
	}
}
