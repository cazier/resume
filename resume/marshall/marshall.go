package marshall

import (
	"encoding/json"
	"os"
	"strings"

	shared "resume/main/shared"
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
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Courses     []string `json:"courses"`
}

type Work struct {
	Company    string      `json:"company"`
	Position   string      `json:"position"`
	Website    string      `json:"website"`
	StartDate  string      `json:"startDate"`
	EndDate    string      `json:"endDate"`
	Summary    string      `json:"summary"`
	Highlights []Highlight `json:"highlights"`
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
	Name        string            `json:"name"`
	StartDate   string            `json:"startDate"`
	EndDate     string            `json:"endDate"`
	Links       map[string]string `json:"links"`
	Keywords    []string          `json:"keywords"`
	Description string            `json:"description"`
	Highlights  []Highlight       `json:"highlights"`
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
