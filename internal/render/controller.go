package render

import (
	"fmt"
	"sort"
	"time"

	"github.com/pkg/errors"
	"github.com/strongdm/comply/internal/config"
	"github.com/strongdm/comply/internal/model"
)

type project struct {
	OrganizationName string
	Name             string
}

type stats struct {
	ControlsTotal     int
	ControlsSatisfied int

	ProcedureTotal      int
	ProcedureOpen       int
	ProcedureOldestDays int

	AuditOpen   int
	AuditClosed int
	AuditTotal  int
}

type control struct {
	Standard    string
	Family      string
	ControlKey  string
	Name        string
	Description string
	Satisfied   bool
	SatisfiedBy []string
}

type controlFamily struct {
	Family   string
	Controls []*control
}

type renderData struct {
	// duplicates Project.OrganizationName
	Name                      string
	Project                   *project
	Stats                     *stats
	Narratives                []*model.Document
	Policies                  []*model.Document
	Procedures                []*model.Procedure
	Standards                 []*model.Standard
	Tickets                   []*model.Ticket
	Controls                  []*control
	ControlsByStandard        map[string][]*control
	ControlFamiliesByStandard map[string][]*controlFamily
	Links                     *model.TicketLinks
}

func load() (*model.Data, *renderData, error) {
	modelData, err := model.ReadData()
	if err != nil {
		return nil, nil, err
	}

	cfg := config.Config()
	project := &project{
		OrganizationName: cfg.Name,
		Name:             fmt.Sprintf("%s Compliance Program", cfg.Name),
	}

	satisfied := model.ControlsSatisfied(modelData)
	controls := make([]*control, 0)
	for _, standard := range modelData.Standards {
		for key, c := range standard.Controls {
			satisfactions, ok := satisfied[key]
			isSatisfied := ok && len(satisfactions) > 0
			controls = append(controls, &control{
				Standard:    standard.Name,
				Family:      c.Family,
				ControlKey:  key,
				Name:        c.Name,
				Description: c.Description,
				Satisfied:   isSatisfied,
				SatisfiedBy: satisfactions,
			})
		}
	}
	sort.Slice(controls, func(i, j int) bool {
		if controls[i].Standard != controls[j].Standard {
			return controls[i].Standard < controls[j].Standard
		}
		return controls[i].ControlKey < controls[j].ControlKey
	})

	controlsByStandard := make(map[string][]*control)
	for _, c := range controls {
		controlsByStandard[c.Standard] = append(controlsByStandard[c.Standard], c)
	}

	// Group controls by standard then family, preserving sorted order within each family
	familiesMap := make(map[string]map[string][]*control)
	for _, c := range controls {
		if familiesMap[c.Standard] == nil {
			familiesMap[c.Standard] = make(map[string][]*control)
		}
		familiesMap[c.Standard][c.Family] = append(familiesMap[c.Standard][c.Family], c)
	}
	controlFamiliesByStandard := make(map[string][]*controlFamily)
	for stdName, familyMap := range familiesMap {
		familyKeys := make([]string, 0, len(familyMap))
		for k := range familyMap {
			familyKeys = append(familyKeys, k)
		}
		sort.Strings(familyKeys)
		families := make([]*controlFamily, 0, len(familyKeys))
		for _, fk := range familyKeys {
			families = append(families, &controlFamily{
				Family:   fk,
				Controls: familyMap[fk],
			})
		}
		controlFamiliesByStandard[stdName] = families
	}

	rd := &renderData{}
	rd.Narratives = modelData.Narratives
	rd.Policies = modelData.Policies
	rd.Procedures = modelData.Procedures
	rd.Standards = modelData.Standards
	rd.Tickets = modelData.Tickets
	rd.Links = &model.TicketLinks{}
	rd.Project = project
	rd.Name = project.OrganizationName
	rd.Controls = controls
	rd.ControlsByStandard = controlsByStandard
	rd.ControlFamiliesByStandard = controlFamiliesByStandard

	ts, err := config.Config().TicketSystem()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error in ticket system configuration")
	}

	tp := model.GetPlugin(model.TicketSystem(ts))
	if tp.Configured() {
		links := tp.Links()
		rd.Links = &links
	}

	return modelData, rd, nil
}

func loadWithStats() (*model.Data, *renderData, error) {
	modelData, renderData, err := load()
	if err != nil {
		return nil, nil, err
	}

	addStats(modelData, renderData)
	return modelData, renderData, nil
}

func addStats(modelData *model.Data, renderData *renderData) {
	stats := &stats{}

	satisfied := model.ControlsSatisfied(modelData)

	for _, std := range renderData.Standards {
		stats.ControlsTotal += len(std.Controls)
		for controlKey := range std.Controls {
			if _, ok := satisfied[controlKey]; ok {
				stats.ControlsSatisfied++
			}
		}
	}

	for _, t := range renderData.Tickets {
		if t.Bool("audit") {
			stats.AuditTotal++
		}

		if t.State == model.Open {
			if t.Bool("comply-procedure") {
				stats.ProcedureOpen++
				if t.CreatedAt != nil {
					age := int(time.Since(*t.CreatedAt).Hours() / float64(24))
					if stats.ProcedureOldestDays < age {
						stats.ProcedureOldestDays = age
					}
				}
			}
			if t.Bool("audit") {
				stats.AuditOpen++
			}
		}
	}

	renderData.Stats = stats
}
