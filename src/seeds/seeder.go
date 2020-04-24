package seeds

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/globalsign/mgo/bson"

	"github.com/UserDatabaseApi/src/interface/repository"

	"github.com/UserDatabaseApi/src/models"
)

type outerGame struct {
	ID           bson.ObjectId `json:"-" bson:"_id"`
	PointsGained string        `json:"points_gained" bson:"points_gained"`
	WinStatus    string        `json:"win_status" bson:"win_status"`
	GameType     string        `json:"game_type" bson:"game_type"`
	Created      string        `json:"created" bson:"created"`
}

type OuterUser struct {
	ID        bson.ObjectId `json:"-" bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	LastName  string        `json:"last_name" bson:"last_name"`
	Country   string        `json:"country" bson:"country"`
	City      string        `json:"city" bson:"city"`
	Gender    string        `json:"gender" bson:"gender"`
	BirthDate string        `json:"birth_date" bson:"birth_date"`
}

func RunUserSeeds(ur repository.UserRepository, filePath string) ([]models.User, error) {
	type userSeedObject struct {
		Users []OuterUser `json:"objects"`
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	var outerUsers userSeedObject
	if err = json.Unmarshal(byteValue, &outerUsers); err != nil {
		return nil, err
	}

	users := make([]models.User, 0, len(outerUsers.Users))
	for i := 0; i < len(outerUsers.Users); i++ {
		var user models.User
		if user, err = parseUser(outerUsers.Users[i]); err != nil {
			return nil, err
		}
		user.ID = bson.NewObjectId()
		users = append(users, user)
		if err = ur.InsertUser(user); err != nil {
			return nil, err
		}
	}

	return users, nil
}

func RunGameSeeds(gr repository.GameRepository, filePath string) ([]models.Game, error) {
	type gameSeedObject struct {
		Games []outerGame `json:"objects"`
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	var outerGames gameSeedObject
	if err = json.Unmarshal(byteValue, &outerGames); err != nil {
		return nil, err
	}

	games := make([]models.Game, 0, len(outerGames.Games))
	for _, outerGame := range outerGames.Games {
		game, err := parseGame(outerGame)
		if err != nil {
			return nil, err
		}
		game.ID = bson.NewObjectId()
		if err = gr.InsertGame(game); err != nil {
			log.Error(err)
		}

		games = append(games, game)
	}

	return games, nil
}

func RunUserGameSeeds(ugr repository.UserGameRepository, users []models.User, games []models.Game, gamesCount int) error {
	if users == nil {
		return errors.New("there is no users to make userGame seed")
	}
	if games == nil {
		return errors.New("there is no games to make userGame seed")
	}
	for _, user := range users {
		var userGame models.UserGame
		userGame.ID = bson.NewObjectId()
		userGame.UserID = user.ID
		userGame.GameIDs = make([]bson.ObjectId, 0, gamesCount)
		for j := 0; j < gamesCount; j++ {
			userGame.GameIDs = append(userGame.GameIDs, games[rand.Intn(len(games))].ID)
		}
		if err := ugr.InsertUserGame(userGame); err != nil {
			return err
		}
	}

	return nil
}

func parseGame(oug outerGame) (g models.Game, err error) {
	g.PointsGained, err = strconv.Atoi(oug.PointsGained)
	if err != nil {
		return
	}

	g.WinStatus, err = strconv.Atoi(oug.WinStatus)
	if err != nil {
		return
	}

	g.GameType, err = strconv.Atoi(oug.GameType)
	if err != nil {
		return
	}

	g.Created, err = time.Parse("1/2/2006 3:04 PM", oug.Created)
	if err != nil {
		return
	}

	return
}

func parseUser(oug OuterUser) (user models.User, err error) {
	user = models.User{Email: oug.Email,
		LastName: oug.LastName,
		Country:  oug.Country,
		City:     oug.City,
		Gender:   oug.Gender}

	user.BirthDate, err = time.Parse("Monday, January 2, 2006 3:04 PM", oug.BirthDate)
	if err != nil {
		return
	}

	return
}
