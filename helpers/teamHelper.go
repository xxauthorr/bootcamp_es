package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/models"
)

type TeamHelper struct {
	team        database.Insert
	check       database.Check
	dbOperation database.DBoperation
}

func (t TeamHelper) TeamScanAndInsert(team models.TeamReg, user string) (string, error) {
	// check weather the leader already have a team
	status := t.check.CheckUserHasClan(user)
	if status {
		return "User already in a team", nil
	}
	t.dbOperation.StartTransaction()
	for i := range team.Players {
		player := team.Players
		if err := t.team.InsertTeamNotification(player[i], team.TeamName, "Member"); err != nil {
			t.dbOperation.RollBackTransaction()
			return "", err
		}
	}
	if team.CoLeader != "" {
		if err := t.team.InsertTeamNotification(team.CoLeader, team.TeamName, "Co-Leader"); err != nil {
			t.dbOperation.RollBackTransaction()
			return "", err
		}
	}
	if err := t.team.InsertTeam(team, user); err != nil {
		t.dbOperation.RollBackTransaction()
		return "", err
	}
	t.dbOperation.CommitTransaction()
	return "", nil
}
