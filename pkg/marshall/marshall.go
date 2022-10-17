package marshall

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	shared "github.com/cazier/resume/pkg/shared"
)

type Resume struct {
	Basics    Basic       `json:"basics"`
	Education []Education `json:"education"`
	Work      []Work      `json:"work"`
	Languages []Language  `json:"languages"`
	Interests []Interest  `json:"interests"`
	Skills    []Skill     `json:"skills"`
	Projects  []Project   `json:"projects"`
}

type Basic struct {
	Name         string              `json:"name"`
	Label        string              `json:"label"`
	Pronouns     string              `json:"pronouns"`
	Image        string              `json:"image,omitempty"`
	Email        string              `json:"email"`
	Phone        string              `json:"phone"`
	Url          string              `json:"url,omitempty"`
	Summary      string              `json:"summary"`
	Location     map[string]string   `json:"location"`
	JsonProfiles []map[string]string `json:"profiles"`
	Profiles     map[string]Profile  `json:"-"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	Url      string `json:"url"`
}

type Education struct {
	Institution   string    `json:"institution"`
	Area          string    `json:"area"`
	StudyType     string    `json:"studyType"`
	StartDate     time.Time `json:"-"`
	JsonStartDate string    `json:"startDate"`
	EndDate       time.Time `json:"-"`
	JsonEndDate   string    `json:"endDate,omitempty"`
	Courses       []string  `json:"courses"`
}

type Work struct {
	Company       string      `json:"company"`
	Position      string      `json:"position"`
	Website       string      `json:"website"`
	StartDate     time.Time   `json:"-"`
	JsonStartDate string      `json:"startDate"`
	EndDate       time.Time   `json:"-"`
	JsonEndDate   string      `json:"endDate,omitempty"`
	Summary       string      `json:"summary"`
	Highlights    []Highlight `json:"highlights"`
}

type Highlight struct {
	Base      string   `json:"base"`
	Lowlights []string `json:"lowlights"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type Interest struct {
	Name string `json:"name"`
}

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type Project struct {
	Name          string            `json:"name"`
	StartDate     time.Time         `json:"-"`
	JsonStartDate string            `json:"startDate"`
	EndDate       time.Time         `json:"-"`
	JsonEndDate   string            `json:"endDate,omitempty"`
	Links         map[string]string `json:"links"`
	Keywords      []string          `json:"keywords"`
	Description   string            `json:"description"`
	Highlights    []Highlight       `json:"highlights"`
}

func LoadJsonFile(path string) Resume {
	data, err := os.ReadFile(path)
	shared.HandleError(err)

	return LoadJsonString(data)
}

func LoadJsonString(data []byte) Resume {
	var output Resume
	json.Unmarshal(data, &output)

	adaptProfiles(&output)
	adaptDates(&output)

	return output
}

func adaptProfiles(resume *Resume) {
	resume.Basics.Profiles = make(map[string]Profile)
	for _, p := range resume.Basics.JsonProfiles {
		var ps *Profile = &Profile{
			p["network"],
			p["username"],
			p["url"],
		}
		resume.Basics.Profiles[strings.ToLower(ps.Network)] = *ps
	}
}

func adaptDates(resume *Resume) {
	for index, education := range resume.Education {
		resume.Education[index].StartDate = parseDate(education.JsonStartDate)
		resume.Education[index].EndDate = parseDate(education.JsonEndDate)
	}

	for index, project := range resume.Projects {
		resume.Projects[index].StartDate = parseDate(project.JsonStartDate)
		resume.Projects[index].EndDate = parseDate(project.JsonEndDate)
	}

	for index, work := range resume.Work {
		resume.Work[index].StartDate = parseDate(work.JsonStartDate)
		resume.Work[index].EndDate = parseDate(work.JsonEndDate)
	}
}

func parseDate(date string) time.Time {
	if date == "" {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		t, err = time.Parse("2006-01", date)
		shared.HandleError(err)
	}
	return t
}
