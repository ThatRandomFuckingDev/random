package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// this is how I learn a programming language
// as I am learning Go, I also interpret what I need
// than I use it to jot down and understand the language
// also programming AI for this purpose is way better
// it teaches me the importance of data, structure, and more within a language

// user struct
type User struct {
	usrName       string
	usrStatus     string
	usrDisability string
	usrAge        string
}

// user input handling
func inputName() User {
	var usr User

	// limits
	var usrCharLim = 50
	var usrAgeLim = 100

	// pricing
	var pricingAgent int

	fmt.Println("Hello, input your username:")
	fmt.Scan(&usr.usrName)

	// limit handling for name
	if len(usr.usrName) > usrCharLim {
		fmt.Println("Username: Too many characters")
		fmt.Println("You only have:", usrCharLim, "use it wisely")
		var usrCharLim_1 = len(usr.usrName)
		fmt.Println("Your username is:", usrCharLim_1,
			"characters long, it's too long")
		return usr
	} else if len(usr.usrName) < usrCharLim {
		fmt.Println("Your name is:", usr.usrName)
	}

	fmt.Println("Please input your age:")
	fmt.Scan(&usr.usrAge)

	// limit handling for age
	usrAgeInt, err := strconv.Atoi(usr.usrAge)
	if err != nil {
		fmt.Println("Please input a valid age")
		return usr
	} else {
		if usrAgeInt == 0 {
			fmt.Println("Please input an age")
			return usr
		} else {
			if usrAgeInt > usrAgeLim {
				fmt.Println("You are waaaay too old")
				fmt.Println("Age limit:", usrAgeLim)
				return usr
			} else {
				fmt.Println("Your age is:", usr.usrAge)
			}
		}
	}

	fmt.Printf("You are %s, you are %s years of age\n", usr.usrName, usr.usrAge)

	fmt.Println("Are you disabled or a veteran?")
	fmt.Scanln(&usr.usrStatus)
	usr.usrStatus = strings.ToLower(usr.usrStatus)

	// price handling
	if usrAgeInt < 10 {
		pricingAgent = usrAgeInt + 5
		fmt.Println("Price for children:", pricingAgent)
	} else if usrAgeInt < 55 {
		pricingAgent = usrAgeInt + 10
		fmt.Println("Price for adults:", pricingAgent)
	} else {
		pricingAgent = 20
		fmt.Println("Senior discount:", pricingAgent)
	}

	// discount handling
	if usr.usrStatus == "no" || usr.usrStatus == "none" {
		fmt.Println("No discount applied")
		fmt.Printf("\nYour price: $%d\n", pricingAgent)
	} else {
		// veteran handling
		if usr.usrStatus == "veteran" || usr.usrStatus == "vet" {
			pricingAgent = 22
			if usrAgeInt > 55 {
				var pricingAgent_v = pricingAgent - 10
				fmt.Println("Veteran + senior discount applied:", pricingAgent_v)
				fmt.Printf("\nYour price: $%d\n", pricingAgent_v)
				return usr
			} else {
				fmt.Println("Veteran discount applied:", pricingAgent)
				fmt.Printf("\nYour price: $%d\n", pricingAgent)
			}
		}

		// disability handling
		if usr.usrStatus == "disabled" {
			pricingAgent = 25

			fmt.Println("What is your disability?")
			fmt.Scanln(&usr.usrDisability)
			usr.usrDisability = strings.ToLower(usr.usrDisability)

			if usr.usrDisability == "visual" {
				fmt.Println("You are exempt from tax")
				var pricingAgent_vi = pricingAgent - 5
				fmt.Println("Disabled discount applied:", pricingAgent_vi)
				fmt.Printf("\nYour price: $%d\n", pricingAgent_vi)
				return usr
			} else if usrAgeInt > 55 {
				var pricingAgent_vi = pricingAgent - 10
				fmt.Println("Disabled + senior discount applied:", pricingAgent_vi)
				fmt.Printf("\nYour price: $%d\n", pricingAgent_vi)
			}
			if usrAgeInt > 55 {
				var pricingAgent_d = pricingAgent - 10
				fmt.Println("Disabled + senior discount applied:", pricingAgent_d)
        fmt.Printf("\nYour price: $%d\n", pricingAgent_d)
        return usr
      } else {
        fmt.Println("Disabled discount applied:", pricingAgent)
        fmt.Printf("\nYour price: $%d\n", pricingAgent)
      }
    }

    // both handling
		if usr.usrStatus == "both" {
			pricingAgent = 15
			if usrAgeInt > 55 {
				var pricingAgent_d_v = pricingAgent - 5
				fmt.Println("Both + senior discount applied:", pricingAgent_d_v)
				fmt.Printf("\nYour price: $%d\n", pricingAgent_d_v)
			} else {
				fmt.Println("Both + senior discount applied:", pricingAgent)
				fmt.Printf("\nYour price: $%d\n", pricingAgent)
			}
		}
	}

	// print statements for debugging
	fmt.Printf("Your name is: %s\n", usr.usrName)
	fmt.Printf("You are: %s years of age\n", usr.usrAge)

	return usr
}
func factorial(n, r int64) int64 {
	var Sub int64

	if n == 1 {
		return 1
	} else if n == 2 {
		Sub = (n - r) * n
	} else {
		fmt.Println("line: 170 Factorial: Returned a zero")
		return 0
	}

	return Sub
}

func nCr(n, r int64) int64 {
	var calcnCr int64

	nF := factorial(n, r)
	rF := factorial(r, r)
	fA := factorial(n-r, r)

	if nF == 0 || rF == 0 || fA == 0 {
		fmt.Println("line: 185 nCr: Returned a zero")
		return 0
	} else {
		calcnCr = nF / (rF * fA)
	}

	return calcnCr
}

func calcNCR() int64 {
	var val int64

	n := int64(rand.Intn(2) + 1)
	r := int64(rand.Intn(int(n) + 1))
	val = nCr(n, r)

	return val
}

func simplernCr() int64 {
	var values int64
	var calc int64
	var caller int64
	var divide int64

	val := calcNCR()
	n := int64(rand.Intn(1) + 1)
	r := int64(rand.Intn(int(n) + 1))
	calcnCr := nCr(n, r)

	values = val
	calc = calcnCr

	if values > 5000 {
		fmt.Println("Values are waaaay too high")
		return -1
	} else if values < 500 {
		fmt.Println("Values still too high")
		return -1
	} else if values < 50 {
		fmt.Println("Values are still too high")
		return -1
	} else {
		fmt.Println("line: 226 Values are safe")
	}

	if calc > 5000 {
		fmt.Println("Calc: Values are waaaay too high")
		return -1
	} else if calc < 500 {
		fmt.Println("Calc: Values still too high")
		return -1
	} else if calc < 50 {
		fmt.Println("Calc: Values are still too high")
		return -1
	} else {
		fmt.Println("line: 241 Values are safe")
	}

	divide = values / calc
	caller = divide + 10

	if err := caller; err != 0 {
		fmt.Println("Line 236: Caller returned a non-zero value:", err)
		return err
	} else {
		fmt.Println("Caller is safe")
		if values > 50 {
			fmt.Println("RETURN VALUES ARE TOO HIGH")
			return -1
		}

		if calc > 50 {
			fmt.Println("RETURN CALC VALUES ARE TOO HIGH")
			return -1
		}
	}

	if caller == -1 {
		fmt.Println("Caller returned a -1 value")
		panic(caller)
	}

	return caller
}

type Computer struct {
	cpu              string
	gpu              string
	memory           uint8
	storage          int64
	operating_system string
}

func (c Computer) getGPU() string {
	return c.gpu
}

func (c Computer) getCPU() string {
	return c.cpu
}

func (c Computer) getMemory() uint8 {
	return c.memory
}

func (c Computer) getStorage() int64 {
	return c.storage
}

func (c Computer) getOS() string {
	return c.operating_system
}

type computer interface {
	getGPU() string
	getCPU() string
	getMemory() uint8
	getStorage() int64
	getOS() string
}

func computer_specifications(c Computer,
	gpu string, cpu string, os string,
	memory uint8, storage int64) {
	gpu = c.gpu
	cpu = c.cpu
	os = c.operating_system
	memory = c.memory
	storage = c.storage

	fmt.Scanln(&gpu)
	fmt.Scanln(&cpu)
	fmt.Scanln(&os)
	fmt.Scanln(&memory)
	fmt.Scanln(&storage)
}

type Engine struct {
	mpg      uint8
	mpkwh    uint8
	kwh      uint8
	gallons  uint8
	gas      uint8
	electric uint8
}

// Gas engine
func (g Engine) milesLeft() uint8 {
	return g.gallons * g.mpg
}

// Electric engine
func (e Engine) kwhLeft() uint8 {
	return e.kwh * e.mpkwh
}

type vehicle interface {
	milesLeft() uint8
	kwhLeft() uint8
}

// Distance before stop
func distanceLeft(e Engine,
	miles uint8, kwh uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("The number of miles before stop:", miles)
	} else {
		fmt.Println("Out of fuel!")
	}

	if kwh <= e.kwhLeft() {
		fmt.Println("The number of miles before stop:", kwh)
	} else {
		fmt.Println("Recharge!")
	}
}

// EngineFunc
func EngineFunc() {
	var gasEngine = Engine{mpg: 35,
		gallons: 20}
	var electricEngine = Engine{mpkwh: 4,
		kwh: 250}
	distanceLeft(gasEngine, 100, 100)
	distanceLeft(electricEngine, 100, 250)
}

// cost functions bellow
type Cost struct {
	AgeCst        float64
	NameCst       float64
	StatusCst     float64
	CompCstString float64
	CompCst       float64
	VehCst        float64
}

func ageCst(arrAge [14]string) float64 {
	r := rand.Float64()
	var result float64
	for _, age := range arrAge {
		arryLenAge := len(age)
		y := float64(arryLenAge) * r
		z := y - float64(arryLenAge)
		result += z * z
	}
	return result
}

func nameCst(arrName [14]string) float64 {
	rN := rand.Float64()
	var resultN float64
	for _, name := range arrName {
		arryLenName := len(name)
		y := float64(arryLenName) * rN
		z := y - float64(arryLenName)
		resultN += z * z
	}
	return resultN
}

func statusCst(arrStatus [14]string) float64 {
	rS := rand.Float64()
	var resultS float64
	for _, status := range arrStatus {
		arryLenStat := len(status)
		y := float64(arryLenStat) * rS
		z := y - float64(arryLenStat)
		resultS += z * z
	}
	return resultS
}

func compCstString(arrComp [14]Computer) float64 {
	sH := rand.Float64()
	var resultsHs float64
	for _, comp := range arrComp {
		stringHardware := comp.cpu + comp.gpu + comp.operating_system
		y := float64(len(stringHardware)) * sH
		z := y - float64(len(stringHardware))
		resultsHs += z * z
	}
	return resultsHs
}

func compCst(arrComp [14]Computer) float64 {
	sH := rand.Float64()
	var resultsH float64
	for _, comp := range arrComp {
		Hardware := int64(comp.memory) + comp.storage
		y := float64(Hardware) * sH
		z := y - float64(Hardware)
		resultsH += z * z
	}
	return resultsH
}

func vehCst(arrVehicle [14]Engine) float64 {
	vR := rand.Float64()
	var resultsR float64
	for _, veh := range arrVehicle {
		travel := uint8(veh.mpg) + veh.mpkwh + veh.kwh + veh.gallons + veh.gas + veh.electric
		y := float64(travel) * vR
		z := y - float64(travel)
		resultsR += z * z
	}
	return resultsR
}

// epsilon and rate functions
func lrEpsilon() float64 {
	const ep = float64(1000000.0)
	return ep
}

func lrRate() float64 {
	const rt = float64(1000.0)
	return rt
}

// float struct
type Floats struct {
	epsi float64
	rate float64
}

// the main function to categorize the names
func main() {
	// random seed
	rand.Seed(time.Now().UnixNano())

	usr := inputName() // calls user input
	caller := simplernCr()

	// structs
	var fl Floats
	var ct Cost

	// comp array
	arrComp := [14]Computer{
		{"Intel x86_64", "Nvidia", 16, 1500, "Windows"},
		{"AMD x86_64", "AMD", 32, 2800, "Linux"},
		{"Apple_M1", "Nvidia", 16, 500, "macOS"},
		{"Intel_ARM", "Intel", 2, 12, "Android"},
		{"AMD x86_64", "AMD", 32, 2500, "Windows"},
		{"Intel x86_64", "Nvidia", 16, 500, "Windows"},
		{"AMD x86_6", "AMD", 16, 256, "Linux"},
		{"Apple_M1", "Apple", 16, 1000, "macOS"},
		{"Intel_ARM", "Intel", 2, 18, "Android"},
		{"Intel x86_64", "Intel", 32, 3000, "Windows"},
		{"Intel x86_64", "AMD", 16, 500, "Windows"},
		{"AMD x86_64", "AMD", 8, 256, "Linux"},
		{"Apple_M1", "Apple", 16, 1800, "macOS"},
		{"Intel_ARM", "Intel", 4, 20, "Android"},
	}

	// vehicle array
	arrVehicle := [14]Engine{
		{mpg: 35, mpkwh: 0, kwh: 0, gallons: 20,
			gas: 1, electric: 0},
		{mpg: 0, mpkwh: 4, kwh: 250, gallons: 0,
			gas: 0, electric: 1},
		{mpg: 25, mpkwh: 0, kwh: 0, gallons: 15,
			gas: 1, electric: 0},
		{mpg: 0, mpkwh: 3, kwh: 200, gallons: 0,
			gas: 0, electric: 1},
		{mpg: 30, mpkwh: 0, kwh: 0, gallons: 18,
			gas: 1, electric: 0},
		{mpg: 35, mpkwh: 0, kwh: 0, gallons: 20,
			gas: 1, electric: 0},
		{mpg: 0, mpkwh: 4, kwh: 250, gallons: 0,
			gas: 0, electric: 1},
		{mpg: 25, mpkwh: 0, kwh: 0, gallons: 15,
			gas: 1, electric: 0},
		{mpg: 0, mpkwh: 3, kwh: 200, gallons: 0,
			gas: 0, electric: 1},
		{mpg: 30, mpkwh: 0, kwh: 0, gallons: 18,
			gas: 1, electric: 0},
		{mpg: 35, mpkwh: 0, kwh: 0, gallons: 20,
			gas: 1, electric: 0},
		{mpg: 0, mpkwh: 4, kwh: 250, gallons: 0,
			gas: 0, electric: 1},
		{mpg: 25, mpkwh: 0, kwh: 0, gallons: 15,
			gas: 1, electric: 0},
		{mpg: 0, mpkwh: 3, kwh: 200, gallons: 0,
			gas: 0, electric: 1},
	}

	// arrays for the neuron
	arrStatus := [14]string{"none", "none",
		"veteran", "disabled",
		"both", "none", "none", "none",
		"disabled", "none", "none",
		"disabled", "none",
		usr.usrStatus}

	arrAge := [14]string{"10", "15", "78", "55", "38", "35", "43",
		"47", "22", "18", "34", "28", "23", usr.usrAge}

	arrName := [14]string{"Velhma", "Mike", "Ivy",
		"Cassedy", "Shawn", "Reiden", "Chirpie",
		"Jessica", "Emily", "Flako", "Flirben",
		"CPU", "Memory", usr.usrName}

	// calling the cost functions and arrays
	ct.AgeCst = ageCst(arrAge)
	ct.NameCst = nameCst(arrName)
	ct.StatusCst = statusCst(arrStatus)
	ct.CompCstString = compCstString(arrComp)
	ct.CompCst = compCst(arrComp)
	ct.VehCst = vehCst(arrVehicle)

	fmt.Printf("Random AgeCst: %f\n", ct.AgeCst)
	fmt.Printf("Random NameCst: %f\n", ct.NameCst)
	fmt.Printf("Random StatusCst: %f\n", ct.StatusCst)
	fmt.Printf("Random CompCst: %f", ct.CompCst)
	fmt.Printf("Random CompCstString: %f", ct.CompCstString)
	fmt.Printf("Random VehCst: %f", ct.VehCst)

	// 2 values
	value := 5000

	// epsilon and rate values
	fl.epsi = lrEpsilon()
	fl.rate = lrRate()

	// for loop for the dirivatives
	for i := 0; i < value; i++ {

		dsctc := (compCst(arrComp) - ct.CompCst) / fl.epsi
		dsctc *= fl.rate

		dsctcs := (compCstString(arrComp) - ct.CompCst) / fl.epsi
		dsctcs *= fl.rate

		dsctv := (vehCst(arrVehicle) - ct.VehCst) / fl.epsi
		dsctv *= fl.rate

		dsct2 := (nameCst(arrName) - ct.NameCst) / fl.epsi
		dsct2 *= fl.rate

		dsct3 := (statusCst(arrStatus) - ct.StatusCst) / fl.epsi
		dsct3 *= fl.rate

		ct.NameCst += dsct2
		ct.StatusCst += dsct3
	}

	// age value for loop, was testing multiple
	ageVal := uint64(500)

	for i := uint64(0); i < ageVal; i++ {
		dsct := (ageCst(arrAge) - ct.AgeCst) / fl.epsi
		dsct *= fl.rate

		ct.AgeCst += dsct
	}

	// error casting
	errAgeCst := math.IsNaN(ct.AgeCst)
	errNameCst := math.IsNaN(ct.NameCst)
	errStatusCst := math.IsNaN(ct.StatusCst)
	errVehCst := math.IsNaN(ct.VehCst)
	errCompCst := math.IsNaN(ct.CompCst)
	errCompCstString := math.IsNaN(ct.CompCstString)

	// error handling
	if errAgeCst || errNameCst || errStatusCst || errVehCst || errCompCst || errCompCstString {
		if errAgeCst {
			fmt.Println("Error: AgeCst is NaN")
		}
		if errNameCst {
			fmt.Println("Error: NameCst is NaN")
		}
		if errStatusCst {
			fmt.Println("Error: StatusCst is NaN")
		}
		if errCompCstString {
			fmt.Println("Error: CompString is NaN")
		}
		if errCompCst {
			fmt.Println("Error: CompCst is NaN")
		}
		if errVehCst {
			fmt.Println("Error: VehCst is NaN")
		}
		return
		// not 100% sure why many people forget the else statement
	} else {

		// printing arrays in a list
		fmt.Println("-------------------------")
		fmt.Println("List of Users:")
		for i := 0; i < len(arrName); i++ {
			fmt.Printf("%d. %s\n", i+1, arrName[i])
			fmt.Printf("   Age: %s\n", arrAge[i])
			fmt.Printf("   Disability Status: %s\n", arrStatus[i])
			fmt.Printf("   Comp: %v\n", arrComp[i])
			fmt.Printf("   Veh: %d\n", arrVehicle[i])
			fmt.Println("-------------------------")
		}

		if caller == -1 {
			fmt.Println("Caller returned -1")
			return
		} else {
			fmt.Println("line: 398 Caller is safe")

			var actual float64
			var divisible float64
			var percent float64
			var ret float64

			var lenStat float64
			lenStat = float64(len(arrStatus))

			if lenStat == 0 {
				fmt.Println("Can't divide by 0")
				return
			} else {
				for i := 0; i < len(arrStatus); i++ {
					divisible = lenStat
					percent = (actual / divisible) * 100.0
					actual = (percent / 100.0) * divisible
					ret = divisible / percent * 100.0
				}

				if math.IsInf(ret, 0) {
					fmt.Println("Value is infinite")
					return
				} else {
					fmt.Printf("Percent of disabled/veterans: \n%f\n", ret)
					fmt.Printf("\n")
				}
			}
			// printing current user
			fmt.Println("Current user:")
			fmt.Printf("%s\n", usr.usrName)
			fmt.Printf("%s\n", usr.usrAge)
			fmt.Printf("%s\n", usr.usrStatus)
		}
	}
}
