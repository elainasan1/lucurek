package commands

import (
	"fmt"
	"strings"
	"time"

	"filec2/Handlers/sessions"
	"filec2/db"

	"github.com/alexeyco/simpletable"
)

func init() {
	Register(&Command{
		Name: "ongoing",

		Admin: false,
		Execute: func(session *sessions.Session, cmd []string) error {
			/*session.Println("  ID       Target       Port    Time     Method  Ending in")
			session.Println("╔════╦════════════════╦═══════╦════════╦════════╦════════════╗")

			allSessions, _ := db.OngoingAttacks()

			for i, s := range allSessions {
				session.Println(fmt.Sprintf("║ %s ║ %s ║ %s ║ %s ║ %s ║ %s ║", fillSpace(fmt.Sprintf("%d", (i+1)), 2, " "), fillSpace(fmt.Sprintf("%s", s.Target), 14, " "), fillSpace(fmt.Sprintf("%d", s.Port), 5, " "), fillSpace(fmt.Sprintf("%d", s.Duration), 6, " "), fillSpace(fmt.Sprintf("%s", s.Method), 6, " "), fillSpace(fmt.Sprintf("%s Seconds", fmt.Sprintf("%d", time.Duration(time.Until(time.Unix(s.End, 0)).Seconds()))), 10, " ")))
			}

			_, err := session.Println("╚════╩════════════════╩═══════╩════════╩════════╩════════════╝")
			return err*/
			table := simpletable.New()
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "#"},
					{Align: simpletable.AlignCenter, Text: "Target"},
					{Align: simpletable.AlignCenter, Text: "Port"},
					{Align: simpletable.AlignCenter, Text: "Duration"},
					{Align: simpletable.AlignCenter, Text: "Method"},
					{Align: simpletable.AlignCenter, Text: "Ending"},
				},
			}

			allattacks, err := db.OngoingAttacks()
			if err != nil {
				return err
			}
			if len(allattacks) == 0 {
				session.Println("    There Is No Ongoing Attacks")
				return nil
			}

			for i, atk := range allattacks {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
					{Text: atk.Target},
					{Text: fmt.Sprint(atk.Port)},
					{Text: fmt.Sprint(atk.Duration)},
					{Text: atk.Method},
					{Text: fmt.Sprintf("%d", time.Duration(time.Until(time.Unix(atk.End, 0)).Seconds()))},
				}

				table.Body.Cells = append(table.Body.Cells, r)
			}

			table.SetStyle(simpletable.StyleCompactLite)
			_, err = session.Println(strings.Replace("    "+table.String(), "\n", "\r\n    ", -1), "\r\n")
			return err
		},
	})
}
