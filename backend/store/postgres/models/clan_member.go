package models

type ClanMember struct {
	PlayerTag        string   `gorm:"primaryKey;not null" json:"playerTag"`
	ClanTag          string   `gorm:"primaryKey;not null" json:"clanTag"`
	AddedByDiscordID string   `gorm:"not null" json:"-"`
	ClanRole         ClanRole `gorm:"not null" json:"clanRole"`

	AddedByUser *User        `gorm:"foreignKey:DiscordID;references:AddedByDiscordID" json:"addedByUser,omitempty"`
	Player      *Player      `gorm:"foreignKey:CocTag;references:PlayerTag" json:"player,omitempty"`
	Clan        *Clan        `gorm:"foreignKey:Tag;references:ClanTag" json:"clan,omitempty"`
	Kickpoints  []*Kickpoint `gorm:"foreignKey:PlayerTag,ClanTag;references:PlayerTag,ClanTag" json:"kickpoints,omitempty"`
}

type ClanRole string

const (
	RoleLeader   ClanRole = "leader"
	RoleCoLeader ClanRole = "coLeader"
	RoleElder    ClanRole = "admin"
	RoleMember   ClanRole = "member"
)

func (r ClanRole) String() string {
	return string(r)
}

func (r ClanRole) Format() string {
	switch r {
	case RoleLeader:
		return "Anführer"
	case RoleCoLeader:
		return "Vize-Anführer"
	case RoleElder:
		return "Ältester"
	case RoleMember:
		return "Mitglied"
	default:
		return "Unbekannte Rolle"
	}
}

type ClanMembers []*ClanMember

func (members ClanMembers) Tags() []string {
	if members == nil {
		return nil
	}

	tags := make([]string, len(members))
	for i, member := range members {
		tags[i] = member.PlayerTag
	}

	return tags
}

func (members ClanMembers) TagsDistinct() []string {
	if len(members) == 0 {
		return nil
	}

	tags := make([]string, 1)
	seen := make(map[string]bool)
	for _, member := range members {
		if seen[member.PlayerTag] {
			continue
		}

		tags = append(tags, member.PlayerTag)
		seen[member.PlayerTag] = true
	}

	return tags
}
