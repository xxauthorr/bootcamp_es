package helpers

import (
	"github.com/xxauthorr/bootcamp_es/database"
	"github.com/xxauthorr/bootcamp_es/models"
)

type TeamHelper struct {
	team        database.Team
	teamupdate  database.TeamProfileUpdate
	dbOperation database.DBoperation
}

func (t TeamHelper) TeamScanAndInsert(team models.TeamReg, user string) error {
	// check weather the leader already have a team
	t.dbOperation.StartTransaction()
	for i := range team.Players {
		if err := t.teamupdate.InsertTeamNotification(team.Players[i], team.TeamName, "Member"); err != nil {
			t.dbOperation.RollBackTransaction()
			return err
		}
	}
	if team.CoLeader != "" {
		if err := t.teamupdate.InsertTeamNotification(team.CoLeader, team.TeamName, "Co-Leader"); err != nil {
			t.dbOperation.RollBackTransaction()
			return err
		}
	}
	if err := t.team.InsertTeam(team, user); err != nil {
		t.dbOperation.RollBackTransaction()
		return err
	}
	t.dbOperation.CommitTransaction()
	return nil
}
