package repository

import (
	"github.com/UserDatabaseApi/src/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type gameRepository struct {
	games *mgo.Collection
}

func NewGameRepository(db *mgo.Session, databaseName string) GameRepository {
	return &gameRepository{db.DB(databaseName).C("games")}
}

// GameRepository is interface for game entity
type GameRepository interface {
	FindGameByID(id bson.ObjectId) (*models.Game, error)
	FindGame(game models.Game) (*models.Game, error)
	InsertGames(games []models.Game) error
	InsertGame(game models.Game) error
}

// FindGameByID find game by id
func (gr *gameRepository) FindGameByID(id bson.ObjectId) (*models.Game, error) {
	var game models.Game

	err := gr.games.FindId(id).One(&game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

// InsertGames create game objects inside db
func (gr *gameRepository) InsertGames(games []models.Game) error {
	for _, g := range games {
		if err := gr.games.Insert(g); err != nil {
			return err
		}
	}

	return nil
}

// InsertGame create game object inside db
func (gr *gameRepository) InsertGame(game models.Game) error {
	if err := gr.games.Insert(game); err != nil {
		return err
	}

	return nil
}

// FindGame find game by parameters except id
func (gr *gameRepository) FindGame(incomingGame models.Game) (*models.Game, error) {
	var game models.Game
	err := gr.games.Find(bson.M{"points_gained": incomingGame.PointsGained,
		"win_status": incomingGame.WinStatus,
		"game_type":  incomingGame.GameType,
		"created":    incomingGame.Created}).One(&game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}
