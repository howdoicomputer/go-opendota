package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeroStatService_HeroStats(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/heroStats", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"id":1,"name":"npc_dota_hero_antimage","localized_name":"Anti-Mage","primary_attr":"agi","attack_type":"Melee","roles":["Carry","Escape","Nuker"],"img":"/apps/dota2/images/heroes/antimage_full.png?","icon":"/apps/dota2/images/heroes/antimage_icon.png","base_health":200,"base_health_regen":1.5,"base_mana":75,"base_mana_regen":0.9,"base_armor":-1,"base_mr":25,"base_attack_min":27,"base_attack_max":31,"base_str":22,"base_agi":22,"base_int":12,"str_gain":1.3,"agi_gain":2.8,"int_gain":1.8,"attack_range":150,"projectile_speed":0,"attack_rate":1.45,"move_speed":310,"turn_rate":0.5,"cm_enabled":true,"legs":2,"pro_win":44,"pro_pick":84,"hero_id":1,"pro_ban":219,"5000_pick":1876,"5000_win":946,"4000_pick":12119,"4000_win":6514,"2000_pick":27883,"2000_win":14879,"3000_pick":35551,"3000_win":19196,"1000_pick":11670,"1000_win":6034}]`)
	})

	expected := []HeroStat{
		{
			ID:                1,
			Name:              "npc_dota_hero_antimage",
			LocalizedName:     "Anti-Mage",
			PrimaryAttr:       "agi",
			AttackType:        "Melee",
			Roles:             []string{"Carry", "Escape", "Nuker"},
			Img:               "/apps/dota2/images/heroes/antimage_full.png?",
			Icon:              "/apps/dota2/images/heroes/antimage_icon.png",
			BaseHealth:        200,
			BaseHealthRegen:   1.5,
			BaseMana:          75,
			BaseManaRegen:     0.9,
			BaseArmor:         -1,
			BaseMr:            25,
			BaseAttackMin:     27,
			BaseAttackMax:     31,
			BaseStr:           22,
			BaseAgi:           22,
			BaseInt:           12,
			StrGain:           1.3,
			AgiGain:           2.8,
			IntGain:           1.8,
			AttackRange:       150,
			ProjectileSpeed:   0,
			AttackRate:        1.45,
			MoveSpeed:         310,
			TurnRate:          0.5,
			CmEnabled:         true,
			Legs:              2,
			ProWin:            44,
			ProPick:           84,
			HeroID:            1,
			ProBan:            219,
			FiveThousandPick:  1876,
			FiveThousandWin:   946,
			FourThousandPick:  12119,
			FourThousandWin:   6514,
			TwoThousandPick:   27883,
			TwoThousandWin:    14879,
			ThreeThousandPick: 35551,
			ThreeThousandWin:  19196,
			OneThousandPick:   11670,
			OneThousandWin:    6034,
		},
	}

	client := NewClient(httpClient)
	herostats, _, err := client.HeroStatService.HeroStats()
	assert.Nil(t, err)
	assert.Equal(t, expected, herostats)
}
