package vinamax

import (
	"fmt"
	"math"
	//	"math/rand"
)

//A particle essentially constains a position, magnetisation
type particle struct {
	x, y, z             float64
	m                   vector
	demagnetising_field vector
	u_anis              vector  // Uniaxial anisotropy axis
	rindex              int     //radius index
	msat                float64 // Saturation magnetisation in A/m
	flip                float64 //time of next flip event
	tempnumber          float64

	tempfield vector
	tempm     vector
	previousm vector
	fehlk1    vector
	fehlk2    vector
	fehlk3    vector
	fehlk4    vector
	fehlk5    vector
	fehlk6    vector
	fehlk7    vector
	fehlk8    vector
	fehlk9    vector
	fehlk10   vector
	fehlk11   vector
	fehlk12   vector
	fehlk13   vector
}

//print position and magnitisation of a particle
func (p particle) string() string {
	return fmt.Sprintf("particle@(%v, %v, %v), %v %v %v", p.x, p.y, p.z, p.m[0], p.m[1], p.m[2])
}

//Gives all particles the same specified anisotropy-axis
func Anisotropy_axis(x, y, z float64) {
	uaniscalled = true
	a := norm(vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].u_anis = a
	}
}

//Gives all particles a random anisotropy-axis
func Anisotropy_random() {
	uaniscalled = true
	for i := range universe.lijst {
		phi := rng.Float64() * (2 * math.Pi)
		theta := 2 * math.Asin(math.Sqrt(rng.Float64()))
		universe.lijst[i].u_anis = vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}

//Gives all particles with random magnetisation orientation
func M_random() {
	magnetisationcalled = true
	for i := range universe.lijst {
		phi := rng.Float64() * (2 * math.Pi)
		theta := 2 * math.Asin(math.Sqrt(rng.Float64()))
		universe.lijst[i].m = vector{math.Sin(theta) * math.Cos(phi), math.Sin(theta) * math.Sin(phi), math.Cos(theta)}
	}
}

//Gives all particles a specified magnetisation direction
func M_uniform(x, y, z float64) {
	magnetisationcalled = true
	a := norm(vector{x, y, z})
	for i := range universe.lijst {
		universe.lijst[i].m = a
	}
}

//Sets the saturation magnetisation of all particles in A/m
func Msat(x float64) {
	msatcalled = true
	for i := range universe.lijst {
		universe.lijst[i].msat = x
	}
}
