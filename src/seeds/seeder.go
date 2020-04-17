package seeds

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/UserDatabaseApi/src/interface/repository"

	"github.com/UserDatabaseApi/src/models"
)

type outerUserGame struct {
	ID           bson.ObjectId `json:"-" bson:"_id"`
	PointsGained string        `json:"points_gained" bson:"points_gained"`
	WinStatus    string        `json:"win_status" bson:"win_status"`
	GameType     string        `json:"game_type" bson:"game_type"`
	Created      string        `json:"created" bson:"created"`
}

func RunUserSeeds(ur repository.UserRepository, filePath string) ([]models.User, error) {
	type userSeedObject struct {
		Users []models.User `json:"objects"`
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	var users userSeedObject

	if err = json.Unmarshal(byteValue, &users); err != nil {
		return nil, err
	}

	for i := 0; i < len(users.Users); i++ {
		users.Users[i].ID = bson.NewObjectId()
		if err = ur.InsertUser(users.Users[i]); err != nil {
			return nil, err
		}
	}

	return users.Users, nil
}

func RunUserGamesSeeds(ugr repository.UserGameRepository, users []models.User, filePath string) error {
	type userGameSeedObject struct {
		UserGames []outerUserGame `json:"objects"`
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	var outerUserGames userGameSeedObject
	if err = json.Unmarshal(byteValue, &outerUserGames); err != nil {
		return err
	}

	userGames := make([]models.UserGame, 0, len(outerUserGames.UserGames))
	for _, oug := range outerUserGames.UserGames {
		ug, err := parseUserGame(oug)
		if err != nil {
			return err
		}
		userGames = append(userGames, ug)
	}

	for i := 0; i < len(users); i++ {
		random := rand.Intn(len(userGames) - 5000)
		if err := setGamesForUser(ugr, users[i].ID, userGames[random:random+5000]); err != nil {
			return err
		}
	}

	return nil
}

func parseUserGame(oug outerUserGame) (ug models.UserGame, err error) {

	ug.PointsGained, err = strconv.Atoi(oug.PointsGained)
	if err != nil {
		return
	}

	ug.WinStatus, err = strconv.Atoi(oug.WinStatus)
	if err != nil {
		return
	}

	ug.GameType, err = strconv.Atoi(oug.GameType)
	if err != nil {
		return
	}

	ug.Created, err = time.Parse("1/2/2006 3:04 PM", oug.Created)
	if err != nil {
		return
	}

	return
}

func setGamesForUser(ugr repository.UserGameRepository, userID bson.ObjectId, userGames []models.UserGame) error {
	for i := 0; i < len(userGames); i++ {
		userGames[i].UserID = userID
	}
	if err := ugr.InsertUserGames(userGames); err != nil {
		return err
	}
	return nil
}
