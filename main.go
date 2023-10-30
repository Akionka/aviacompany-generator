package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type (
	Airport struct {
		IataCode string `db:"iata_code"`
		City     string `db:"city"`
		Timezone string `db:"timezone"`
		Name     string `db:"name"`
	}

	Attendant struct {
		ID        int    `db:"id"`
		GivenName string `db:"given_name"`
		LastName  string `db:"last_name"`
	}

	BookingOffice struct {
		ID          int    `db:"id"`
		Address     string `db:"address"`
		PhoneNumber string `db:"phone_number"`
	}

	Cashier struct {
		ID         int    `db:"id"`
		Login      string `db:"login"`
		LastName   string `db:"last_name"`
		FirstName  string `db:"first_name"`
		MiddleName string `db:"middle_name"`
		Password   []byte `db:"password"`
	}

	Line struct {
		LineCode   string  `db:"line_code"`
		DepTime    string  `db:"dep_time"`
		FlightTime string  `db:"flight_time"`
		BasePrice  float64 `db:"base_price"`
		DepAirport string  `db:"dep_airport"`
		ArrAirport string  `db:"arr_airport"`
	}

	PlaneModel struct {
		IcaoTypeDes  string `db:"icao_type_des"`
		Manufacturer string `db:"manufacturer"`
		Model        string `db:"model"`
	}

	Plane struct {
		IataCode  string `db:"iata_code"`
		ModelCode string `db:"model_code"`
	}

	Pilot struct {
		ID        int    `db:"id"`
		GivenName string `db:"given_name"`
		LastName  string `db:"last_name"`
	}

	Flight struct {
		ID         int       `db:"id"`
		DepDate    time.Time `db:"dep_date"`
		ActDepTime time.Time `db:"act_dep_time"`
		ActArrTime time.Time `db:"act_arr_time"`
		LineCode   string    `db:"line_code"`
		PlaneCode  string    `db:"plane_code"`
		PilotID    int       `db:"pilot_id"`
		CopilotID  int       `db:"copilot_id"`
	}

	Seat struct {
		ID        int    `db:"id"`
		Number    string `db:"number"`
		Class     string `db:"class"`
		ModelCode string `db:"model_code"`
	}

	Tariff struct {
		ID          int     `db:"id"`
		Name        string  `db:"name"`
		IsAvaliable bool    `db:"is_avaliable"`
		Description string  `db:"description"`
		Multiplier  float64 `db:"multiplier"`
	}

	Purchase struct {
		ID              int       `db:"id"`
		Date            time.Time `db:"date"`
		BookingOfficeID int       `db:"booking_office_id"`
		TotalPrice      float64   `db:"total_price"`
		ContactPhone    string    `db:"contact_phone"`
		ContactEmail    string    `db:"contact_email"`
		CashierID       int       `db:"cashier_id"`
		TariffID        int       `db:"tariff_id"`
	}

	Passenger struct {
		ID             int       `db:"id"`
		LastName       string    `db:"last_name"`
		GivenName      string    `db:"given_name"`
		Sex            bool      `db:"sex"`
		PassportNumber string    `db:"passport_number"`
		BirthDate      time.Time `db:"birth_date"`
	}

	Ticket struct {
		ID          int `db:"id"`
		PurchaseID  int `db:"purchase_id"`
		PassengerID int `db:"passenger_id"`
	}

	FlightTicket struct {
		ID       int `db:"id"`
		FlightID int `db:"flight_id"`
		SeatID   int `db:"seat_id"`
		TicketID int `db:"ticket_id"`
	}

	FlightAttendant struct {
		AttendantID int `db:"attendant_id"`
		FlightID    int `db:"flight_id"`
	}

	PilotFliesPlaneModel struct {
		PilotID   int    `db:"pilot_id"`
		ModelCode string `db:"model_code"`
	}
)

var (
	dbUsername string
	dbPassword string
	dbHost     string
)

func init() {
	flag.StringVar(&dbUsername, "Username", "root", "username of the db user")
	flag.StringVar(&dbPassword, "Password", "password", "password of the db user")
	flag.StringVar(&dbHost, "Host", "localhost", "host of the db")
}

func main() {
	flag.Parse()
	rand.Seed(1)
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s)/aviacompany?parseTime=true&time_zone=%%27GMT%%27", dbUsername, dbPassword, dbHost))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	var (
		airports              = []Airport{}
		attendants            = []Attendant{}
		bookingOffices        = []BookingOffice{}
		cashiers              = []Cashier{}
		lines                 = []Line{}
		planeModels           = []PlaneModel{}
		planes                = []Plane{}
		pilots                = []Pilot{}
		flights               = []Flight{}
		seats                 = []Seat{}
		tariffs               = []Tariff{}
		purchases             = []Purchase{}
		passengers            = []Passenger{}
		tickets               = []Ticket{}
		flightTickets         = []FlightTicket{}
		flightAttendants      = []FlightAttendant{}
		pilotFliesPlaneModels = []PilotFliesPlaneModel{}
	)

	if err := db.Select(&airports, "SELECT * FROM airport"); err != nil {
		log.Fatal(err)
	}

	if err := db.Select(&bookingOffices, "SELECT * FROM booking_office"); err != nil {
		log.Fatal(err)
	}
	if err := db.Select(&planes, "SELECT * FROM plane"); err != nil {
		log.Fatal(err)
	}
	if err := db.Select(&planeModels, "SELECT * FROM plane_model"); err != nil {
		log.Fatal(err)
	}
	if err := db.Select(&tariffs, "SELECT * FROM tariff"); err != nil {
		log.Fatal(err)
	}

	const (
		attendantCount = 200
		cashierCount   = 25
		lineCount      = 200
		pilotCount     = 500
		passengerCount = 1_000_000
	)

	for i := 0; i < attendantCount; i++ {
		isMale := rand.Float64() < 0.1
		var (
			names     = femaleNames
			lastNames = femaleLastNames
			midNames  = femaleMidNames
		)
		if isMale {
			names = maleNames
			lastNames = maleLastNames
			midNames = maleMidNames
		}
		attendants = append(attendants, Attendant{
			ID:        i + 1,
			GivenName: names[rand.Intn(len(names))] + " " + midNames[rand.Intn(len(midNames))],
			LastName:  lastNames[rand.Intn(len(lastNames))],
		})
	}

	for i := 0; i < cashierCount; i++ {
		isMale := rand.Float64() < 0.5
		var (
			names     = femaleNames
			lastNames = femaleLastNames
			midNames  = femaleMidNames
		)
		if isMale {
			names = maleNames
			lastNames = maleLastNames
			midNames = maleMidNames
		}
		cashiers = append(cashiers, Cashier{
			ID:         i + 1,
			Login:      fmt.Sprintf("login%d", i+1),
			LastName:   lastNames[rand.Intn(len(lastNames))],
			FirstName:  names[rand.Intn(len(names))],
			MiddleName: midNames[rand.Intn(len(midNames))],
			Password:   []byte(fmt.Sprintf("password%d", i+1)),
		})
	}

	for i := 0; i < pilotCount; i++ {
		isMale := rand.Float64() < 0.99
		var (
			names     = femaleNames
			lastNames = femaleLastNames
			midNames  = femaleMidNames
		)
		if isMale {
			names = maleNames
			lastNames = maleLastNames
			midNames = maleMidNames
		}
		pilots = append(pilots, Pilot{
			ID:        i + 1,
			GivenName: names[rand.Intn(len(names))] + " " + midNames[rand.Intn(len(midNames))],
			LastName:  lastNames[rand.Intn(len(lastNames))],
		})
		for _, pm := range planeModels {
			pilotFliesPlaneModels = append(pilotFliesPlaneModels, PilotFliesPlaneModel{PilotID: i + 1, ModelCode: pm.IcaoTypeDes})
		}
	}

	t := 1
	for i := 1; i <= 5; i++ {
		for j := 'A'; j <= 'D'; j++ {
			seats = append(seats, Seat{
				ID:        t,
				Number:    fmt.Sprintf("%d%c", i, j),
				Class:     "J",
				ModelCode: "B738",
			})
			t++
		}
	}
	for i := 6; i <= 10; i++ {
		for j := 'A'; j <= 'F'; j++ {
			seats = append(seats, Seat{
				ID:        t,
				Number:    fmt.Sprintf("%d%c", i, j),
				Class:     "W",
				ModelCode: "B738",
			})
			t++
		}
	}
	for i := 11; i <= 30; i++ {
		for j := 'A'; j <= 'F'; j++ {
			seats = append(seats, Seat{
				ID:        t,
				Number:    fmt.Sprintf("%d%c", i, j),
				Class:     "Y",
				ModelCode: "B738",
			})
			t++
		}
	}

	for i := 0; i < lineCount; i++ {
		depAirport := airports[rand.Intn(len(airports))].IataCode
		arrAirport := airports[rand.Intn(len(airports))].IataCode
		for depAirport == arrAirport {
			arrAirport = airports[rand.Intn(len(airports))].IataCode
		}
		depTime := time.Duration(rand.Intn(24))*time.Hour + time.Duration(rand.Intn(12)*5)*time.Minute
		flightTime := time.Duration(rand.Intn(5))*time.Hour + time.Duration(rand.Intn(12)*5)*time.Minute
		lines = append(lines, Line{
			LineCode:   fmt.Sprintf("JR%d", i+1),
			DepTime:    fmt.Sprintf("%02d:%02d:00", int(depTime.Hours()), int(depTime.Minutes())%60),
			FlightTime: fmt.Sprintf("%02d:%02d:00", int(flightTime.Hours()), int(flightTime.Minutes())%60),
			BasePrice:  float64(rand.Intn(10000) + 2000),
			DepAirport: depAirport,
			ArrAirport: arrAirport,
		})
	}

	t0 := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Now()

	t = 0
	for t0.Before(t1) {
		for _, line := range lines {
			plane := planes[rand.Intn(len(planes))]
			pilotID := rand.Intn(pilotCount) + 1
			copilotID := rand.Intn(pilotCount) + 1
			for copilotID == pilotID {
				copilotID = rand.Intn(pilotCount) + 1
			}
			depTime, _ := time.Parse(time.TimeOnly, line.DepTime)
			flightTime, _ := time.Parse(time.TimeOnly, line.FlightTime)
			actDepTime := time.Date(t0.Year(), t0.Month(), t0.Day(), depTime.Hour(), depTime.Minute()+rand.Intn(60), 0, 0, time.UTC)
			actArrTime := time.Date(t0.Year(), t0.Month(), t0.Day(), depTime.Hour()+flightTime.Hour(), depTime.Minute()+rand.Intn(60)+flightTime.Minute()+actDepTime.Minute(), 0, 0, time.UTC)
			flights = append(flights, Flight{
				ID:         t + 1,
				DepDate:    t0,
				ActDepTime: actDepTime,
				ActArrTime: actArrTime,
				LineCode:   line.LineCode,
				PlaneCode:  plane.IataCode,
				PilotID:    pilotID,
				CopilotID:  copilotID,
			})
			t++
		}
		t0 = t0.AddDate(0, 0, 1)
	}

	passportNumbers := map[string]struct{}{}

	t1 = time.Date(1960, time.January, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Now().AddDate(-20, 0, 0)

	for i := 0; i < passengerCount; i++ {
		isMale := rand.Float64() < 0.5
		var (
			names     = femaleNames
			lastNames = femaleLastNames
			midNames  = femaleMidNames
		)
		if isMale {
			names = maleNames
			lastNames = maleLastNames
			midNames = maleMidNames
		}
		var (
			givenName      = names[rand.Intn(len(names))] + " " + midNames[rand.Intn(len(midNames))]
			lastName       = lastNames[rand.Intn(len(lastNames))]
			birthDate      = t1.Add(time.Duration(rand.Int63n(t2.Unix()-t1.Unix())) * time.Second)
			region         = rand.Intn(89) + 1
			year           = (birthDate.Year() + 18) % 100
			passportNumber string
		)
		for {
			passportNumber = fmt.Sprintf("%02d%02d%06d", region, year, rand.Intn(999999)+1)
			if _, ok := passportNumbers[passportNumber]; !ok {
				passportNumbers[passportNumber] = struct{}{}
				break
			}
		}
		passengers = append(passengers, Passenger{
			ID:             i + 1,
			LastName:       lastName,
			GivenName:      givenName,
			BirthDate:      birthDate,
			PassportNumber: passportNumber,
			Sex:            isMale,
		})
	}

	for _, flight := range flights {
		attendants := make(map[int]bool)
		for len(attendants) < 5 {
			attendantID := rand.Intn(attendantCount) + 1
			if !attendants[attendantID] {
				attendants[attendantID] = true
				flightAttendants = append(flightAttendants, FlightAttendant{
					AttendantID: attendantID,
					FlightID:    flight.ID,
				})
			}
		}

		ps := map[int]struct{}{}
		seatsTaken := 0
		seatsAvailable := len(seats)
		for seatsTaken < seatsAvailable {
			purchases = append(purchases, Purchase{
				ID:              len(purchases) + 1,
				Date:            flight.DepDate.Add(-(time.Hour*24 + time.Second*time.Duration(rand.Int63n(86400*7)))),
				BookingOfficeID: rand.Intn(len(bookingOffices)) + 1,
				TotalPrice:      10_000,
				TariffID:        rand.Intn(len(tariffs)) + 1,
				ContactPhone:    fmt.Sprintf("7%03d%07d", rand.Intn(999)+1, rand.Intn(9999999)+1),
				ContactEmail:    fmt.Sprintf("email%d@gmail.com", rand.Intn(999999999)+1),
				CashierID:       rand.Intn(11) + 1,
			})

			tCount := min(rand.Intn(5)+1, seatsAvailable-seatsTaken)
			for i := 0; i < tCount; i++ {
				var p int
				for {
					p = rand.Intn(len(passengers))
					if _, ok := ps[p]; !ok {
						ps[p] = struct{}{}
						break
					}
				}

				passenger := passengers[p]
				tickets = append(tickets, Ticket{
					ID:          len(tickets) + 1,
					PassengerID: passenger.ID,
					PurchaseID:  purchases[len(purchases)-1].ID,
				})
				flightTickets = append(flightTickets, FlightTicket{
					FlightID: flight.ID,
					SeatID:   seatsTaken + 1,
					TicketID: len(tickets),
				})
				seatsTaken++
			}
		}
	}

	tx := db.MustBegin()

	wg := &sync.WaitGroup{}

	for i := 0; i < len(attendants); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO attendant (id, given_name, last_name)
				VALUES (:id, :given_name, :last_name)`,
				attendants[i:min(len(attendants), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	attendants = nil

	for i := 0; i < len(passengers); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO passenger (id, last_name, given_name, sex, passport_number, birth_date) VALUES (:id, :last_name, :given_name, :sex, :passport_number, :birth_date);`,
				passengers[i:min(len(passengers), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	passengers = nil

	for i := 0; i < len(cashiers); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO cashier (id, login, last_name, first_name, middle_name, password)
				VALUES (:id, :login, :last_name, :first_name, :middle_name, :password)`,
				cashiers[i:min(len(cashiers), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	cashiers = nil

	for i := 0; i < len(lines); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO line (line_code, dep_time, flight_time, base_price, dep_airport, arr_airport)
				VALUES (:line_code, :dep_time, :flight_time, :base_price, :dep_airport, :arr_airport)`,
				lines[i:min(len(lines), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	lines = nil

	for i := 0; i < len(pilots); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO pilot (id, given_name, last_name)
				VALUES (:id, :given_name, :last_name)`,
				pilots[i:min(len(pilots), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	pilots = nil

	for i := 0; i < len(pilotFliesPlaneModels); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO pilot_flies_plane_model (pilot_id, model_code)
				VALUES (:pilot_id, :model_code)`,
				pilotFliesPlaneModels[i:min(len(pilotFliesPlaneModels), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	pilotFliesPlaneModels = nil

	for i := 0; i < len(seats); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO seat (id, number, class, model_code)
				VALUES (:id, :number, :class, :model_code)`,
				seats[i:min(len(seats), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	seats = nil

	for i := 0; i < len(flights); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO flight (id, dep_date, act_dep_time, act_arr_time, line_code, plane_code, pilot_id, copilot_id)
				VALUES (:id, :dep_date, :act_dep_time, :act_arr_time, :line_code, :plane_code, :pilot_id, :copilot_id)`,
				flights[i:min(len(flights), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	flights = nil

	for i := len(purchases) - 1; i >= 0; i -= 5000 {
		// wg.Add(1)
		// go func(i int) {
		_, err := db.NamedExec(
			`INSERT INTO purchase (id, date, booking_office_id, total_price, contact_phone, contact_email, cashier_id, tariff_id)
				VALUES (:id, :date, :booking_office_id, :total_price, :contact_phone, :contact_email, :cashier_id, :tariff_id)`,
			purchases[max(0, i-4999):i+1])
		// wg.Done()
		if err != nil {
			log.Fatal(err)
		}
		purchases = purchases[:max(0, i-4999)]
		// }(i)
	}
	// wg.Wait()
	purchases = nil

	for i := 0; i < len(flightAttendants); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO flight_attendant (attendant_id, flight_id)
				VALUES (:attendant_id, :flight_id)`,
				flightAttendants[i:min(len(flightAttendants), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	flightAttendants = nil

	for i := 0; i < len(tickets); i += 1000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT INTO ticket (id, purchase_id, passenger_id)
				VALUES (:id, :purchase_id, :passenger_id)`,
				tickets[i:min(len(tickets), i+1000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	tickets = nil

	for i := 0; i < len(flightTickets); i += 10000 {
		wg.Add(1)
		go func(i int) {
			_, err := db.NamedExec(
				`INSERT IGNORE INTO flight_in_ticket (flight_id, seat_id, ticket_id)
				VALUES (:flight_id, :seat_id, :ticket_id)`,
				flightTickets[i:min(len(flightTickets), i+10000)])
			wg.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
	wg.Wait()
	flightTickets = nil

	tx.Commit()

}
