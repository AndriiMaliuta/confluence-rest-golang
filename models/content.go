package models

import "time"

type VersionE struct {
	Number int `json:"number"`
}

type Metadata struct {
	Labels LabelArray `json:"labels"`
}

type Icon struct {
	Path      string `json:"path"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsDefault bool   `json:"isDefault"`
}

type UserDetails struct {
	Business string
	Personal string
}

type User struct {
	Type           string      `json:"type"`
	Username       string      `json:"username"`
	UserKey        string      `json:"userKey"`
	AccountId      string      `json:"accountId"`
	AccountType    string      `json:"accountType"`
	EMail          string      `json:"email"`
	PublicName     string      `json:"publicName"`
	DisplayName    string      `json:"displayName"`
	ProfilePicture Icon        `json:"profilePicture"`
	TimeZone       string      `json:"timeZone"`
	Details        UserDetails `json:"details"`
}

type Version struct {
	By        User   `json:"by"`
	When      string `json:"when"`
	Number    int    `json:"number"`
	MinorEdit bool   `json:"minorEdit"`
}

type ContentHistory struct {
	Latest      bool
	CreatedBy   User
	CreatedDate string  `json:"createdDate"`
	LastUpdated Version `json:"lastUpdated"`
}

//type View struct {
//	Value          string `json:"value"`
//	Representation string `json:"representation"`
//}

type Storage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

type Body struct {
	//View    View    `json:"view"`
	Storage Storage `json:"storage"`
}

type Ancestor struct {
	Id string `json:"id"`
}

type Expandable struct {
	//Container Content `json:"container"`
	Operations string `json:"operations"`
}

type Label struct {
	Id     string `json:"id"`
	Prefix string `json:"prefix"`
	Name   string `json:"name"`
}

type LabelArray struct {
	Results []Label `json:"results"`
	start   int
	limit   int
	size    int
	Links   GenericLinks `json:"_links"`
}

type Extensions struct {
	Location   string `json:"location"`
	Resolution string `json:"resolution"`
}

type Content2 struct {
	Id         string         `json:"id"`
	Type       string         `json:"type"`
	Status     string         `json:"status"`
	Title      string         `json:"title"`
	Body       Body           `json:"body"`
	Version    Version        `json:"version"`
	Space      Space          `json:"space"`
	History    ContentHistory `json:"history"`
	Links      GenericLinks   `json:"_links"`
	Ancestors  []Ancestor     `json:"ancestors"`
	Expandable Expandable     `json:"_expandable"`
	Extensions Extensions     `json:"extensions"`
}

type Content struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Title   string `json:"title"`
	Space   Space  `json:"space"`
	History struct {
		Latest    bool `json:"latest"`
		CreatedBy struct {
			Type           string `json:"type"`
			AccountID      string `json:"accountId"`
			AccountType    string `json:"accountType"`
			Email          string `json:"email"`
			PublicName     string `json:"publicName"`
			TimeZone       string `json:"timeZone"`
			ProfilePicture struct {
				Path      string `json:"path"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsDefault bool   `json:"isDefault"`
			} `json:"profilePicture"`
			DisplayName            string `json:"displayName"`
			IsExternalCollaborator bool   `json:"isExternalCollaborator"`
			Expandable             struct {
				Operations    string `json:"operations"`
				PersonalSpace string `json:"personalSpace"`
			} `json:"_expandable"`
			Links struct {
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"createdBy"`
		CreatedDate time.Time `json:"createdDate"`
		Expandable  struct {
			LastUpdated     string `json:"lastUpdated"`
			PreviousVersion string `json:"previousVersion"`
			Contributors    string `json:"contributors"`
			NextVersion     string `json:"nextVersion"`
		} `json:"_expandable"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"history"`
	Version struct {
		By struct {
			Type           string `json:"type"`
			AccountID      string `json:"accountId"`
			AccountType    string `json:"accountType"`
			Email          string `json:"email"`
			PublicName     string `json:"publicName"`
			TimeZone       string `json:"timeZone"`
			ProfilePicture struct {
				Path      string `json:"path"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
				IsDefault bool   `json:"isDefault"`
			} `json:"profilePicture"`
			DisplayName            string `json:"displayName"`
			IsExternalCollaborator bool   `json:"isExternalCollaborator"`
			Expandable             struct {
				Operations    string `json:"operations"`
				PersonalSpace string `json:"personalSpace"`
			} `json:"_expandable"`
			Links struct {
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"by"`
		When                time.Time `json:"when"`
		FriendlyWhen        string    `json:"friendlyWhen"`
		Message             string    `json:"message"`
		Number              int       `json:"number"`
		MinorEdit           bool      `json:"minorEdit"`
		SyncRev             string    `json:"syncRev"`
		SyncRevSource       string    `json:"syncRevSource"`
		ConfRev             string    `json:"confRev"`
		ContentTypeModified bool      `json:"contentTypeModified"`
		Expandable          struct {
			Collaborators string `json:"collaborators"`
			Content       string `json:"content"`
		} `json:"_expandable"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"version"`
	MacroRenderedOutput struct {
	} `json:"macroRenderedOutput"`
	Body struct {
		Storage struct {
			Value           string        `json:"value"`
			Representation  string        `json:"representation"`
			EmbeddedContent []interface{} `json:"embeddedContent"`
			Expandable      struct {
				Content string `json:"content"`
			} `json:"_expandable"`
		} `json:"storage"`
		Expandable struct {
			Editor              string `json:"editor"`
			AtlasDocFormat      string `json:"atlas_doc_format"`
			View                string `json:"view"`
			ExportView          string `json:"export_view"`
			StyledView          string `json:"styled_view"`
			Dynamic             string `json:"dynamic"`
			Editor2             string `json:"editor2"`
			AnonymousExportView string `json:"anonymous_export_view"`
		} `json:"_expandable"`
	} `json:"body"`
	Extensions struct {
		Position int `json:"position"`
	} `json:"extensions"`
	Expandable struct {
		ChildTypes          string `json:"childTypes"`
		Container           string `json:"container"`
		Metadata            string `json:"metadata"`
		Operations          string `json:"operations"`
		SchedulePublishDate string `json:"schedulePublishDate"`
		Children            string `json:"children"`
		Restrictions        string `json:"restrictions"`
		Ancestors           string `json:"ancestors"`
		Descendants         string `json:"descendants"`
	} `json:"_expandable"`
	Links GenericLinks `json:"_links"`
}