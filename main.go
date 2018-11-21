package main

import (
	"fmt"
	"log"
	"net/http"
	"./models"
	"github.com/rs/cors"
	"encoding/json"
	"io"
	"strconv"
	"net/url"
)

type Env struct {
	db models.Datastore
}

type Winner struct {
	PlayerId string `json:"playerId"`
	Prize float64 `json:"prize"`
}
type ResultJson struct {
	TournamentId string `json:"tournamentId"`
	Winners []Winner `json:"winners"`
}

var env *Env
func init (){
	db, err := models.NewDB("Tournament.db")
	if err != nil {
		log.Panic(err)
	}

	env = &Env{db}
	env.db.CreateTables()
}
func main() {


	mux := http.NewServeMux()
	mux.HandleFunc("/", env.indexHandler)
	mux.HandleFunc("/take", env.take)
	mux.HandleFunc("/fund", env.fund)
	mux.HandleFunc("/announceTournament", env.announceTournament)
	mux.HandleFunc("/joinTournament", env.joinTournament)
	mux.HandleFunc("/resultTournament", env.resultTournament)
	mux.HandleFunc("/balance", env.balance)
	mux.HandleFunc("/reset", env.reset)

	handler := cors.Default().Handler(mux)
	fmt.Println(http.ListenAndServe(":7070", handler))
}


func (env *Env) indexHandler(w http.ResponseWriter, r *http.Request) {

	env.db.CreateTables()

}

func (env *Env) take(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	} else {
		playerId := r.URL.Query().Get("playerId")
		if playerId == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		points := r.URL.Query().Get("points")
		if points == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		fmt.Println("Take from Player id: ", playerId)
		fmt.Println("Take points: ", points)
		pointFloat, err := strconv.ParseFloat(points,  64)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		env.TakeFrom (playerId , pointFloat)

	}

}

func (env *Env) fund(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	} else {

		playerId := r.URL.Query().Get("playerId")
		if playerId == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		points := r.URL.Query().Get("points")
		if points == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		fmt.Println("Fund to Player id: ",playerId)
		fmt.Println("Fund points: ",points)
		pointFloat, err := strconv.ParseFloat(points,  64)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		env.FundIn (playerId , pointFloat)

	}
}

func (env *Env) announceTournament(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	} else {
		tournamentId := r.URL.Query().Get("tournamentId")
		if tournamentId == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		deposit := r.URL.Query().Get("deposit")
		if deposit == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		fmt.Println("Announce Tournament Id: ", tournamentId)
		fmt.Println("Announce deposit: ", deposit)
		depositFloat, err := strconv.ParseFloat(deposit,  64)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		env.db.AddTournament(tournamentId, depositFloat)


	}
}

func (env *Env) joinTournament(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	} else {
		tournamentId := r.URL.Query().Get("tournamentId")
		if tournamentId == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		playerId := r.URL.Query().Get("playerId")
		if playerId == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		fmt.Println("Join Tournament: ", tournamentId)
		fmt.Println("Player id: ", playerId)

		var backerIds []string
		var backerId string

		mapQuery, _ := url.ParseQuery(r.URL.RawQuery)

		if mapQuery["backerId"] != nil {
			for i := range mapQuery["backerId"] {
				fmt.Println("Backers: ", mapQuery["backerId"][i])
				backerId = mapQuery["backerId"][i]
				backerIds = append(backerIds, backerId)
			}
		} else {
			backerIds = append(backerIds, "")
		}

		env.JoinDb(tournamentId, playerId, backerIds)

	}
}

func (env *Env) resultTournament(w http.ResponseWriter, r *http.Request){

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	} else {
		var result ResultJson
		dec := json.NewDecoder(r.Body)
		for {
			if err := dec.Decode(&result); err == io.EOF {
				break
			}
		}

		fmt.Println("Result TournamentId: ", result.TournamentId)
		fmt.Println( "Result playerId: ", result.Winners[0].PlayerId)
		fmt.Println( "Result prize: ", result.Winners[0].Prize)

		env.ResultDb(result.TournamentId, result.Winners[0].PlayerId, result.Winners[0].Prize)

	}
}

func (env *Env) balance(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	} else {

		playerId := r.URL.Query().Get("playerId")
		fmt.Println("Balance for Player id: ", playerId)
		balanceRes := env.db.SelectPlayer(playerId)
		if balanceRes.PlayerId !="" {
			json.NewEncoder(w).Encode(balanceRes)
		} else {
			http.NotFound(w, r)
		}
	}
}

func (env *Env) reset(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	} else {
		env.db.DropTables()

		fmt.Println("DROP TABLE")
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (env *Env) TakeFrom (playerId string, points float64)  {

	player := env.db.SelectPlayer(playerId)
	fmt.Println("Check points from Db before: ", player.Points)

	if (player.Points - points) >= 0 {
		env.db.TakeFromPlayer(playerId, points)
		fmt.Println( "done")
	} else {
		fmt.Println("you can't do that you have low balance")
	}

}

func (env *Env) FundIn(playerId string, points float64)  {

	player := env.db.SelectPlayer(playerId)
	fmt.Println("Check Player id:", player.PlayerId)
	if player.PlayerId == ""{
		env.db.CreateNewPlayer(playerId, points)
		fmt.Println( "create new user done")
	} else {
		env.db.FundInPlayer(playerId, points)
		fmt.Println( "points added")
	}

}

func (env *Env) JoinDb(tournamentId string, playerId string, backerIds []string) {

	tournament := env.db.SelectTournament(tournamentId)


	if !env.db.PlayerInGame (playerId, tournamentId) {

		if backerIds[0] != "" {
			partDeposit := tournament.Deposit / float64(len(backerIds)+1)
			env.TakeFrom(playerId, partDeposit)
			env.db.CreateNewBacker(playerId, playerId, tournamentId, partDeposit)
			for _, backerId := range backerIds {
				env.TakeFrom(backerId, partDeposit) //take from balance
				env.db.CreateNewBacker(playerId, backerId, tournamentId, partDeposit) //put into tournament
			}
			fmt.Println("you join Tournament with help")
		} else {
			env.TakeFrom(playerId, tournament.Deposit)
			env.db.CreateNewBacker(playerId, playerId, tournamentId, tournament.Deposit)
			fmt.Println( "you join yourself Tournament")
		}

	}
	fmt.Println( "you already in Tournament")
}

func (env *Env) ResultDb(tournamentId string, winnerId string, prize float64){

	var backerIds[] string = env.db.SelectBackers(winnerId)


	fmt.Println("Seperate prize for parts: ", len(backerIds))
	prizeAmount := prize/float64(len(backerIds))
	for  _, backerId := range backerIds{
		env.db.FundInPlayer(backerId, prizeAmount)
	}

	fmt.Println("Tournament is finished everybody got their points")
	env.db.ZeroBackers(tournamentId)
}
