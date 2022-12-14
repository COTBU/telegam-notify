package model

type JiraPayload struct {
	Id    int `json:"id"`
	Issue struct {
		Id     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Summary     string   `json:"summary"`
			Created     string   `json:"created"`
			Description string   `json:"description"`
			Labels      []string `json:"labels"`
			Priority    string   `json:"priority"`
		} `json:"fields"`
	} `json:"issue"`
	User struct {
		Self         string `json:"self"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		EmailAddress string `json:"emailAddress"`
		DisplayName  string `json:"displayName"`
		Active       string `json:"active"`
	} `json:"user"`
	Changelog struct {
		Items []struct {
			ToString   string  `json:"toString"`
			To         *string `json:"to"`
			FromString string  `json:"fromString"`
			From       *string `json:"from"`
			Fieldtype  string  `json:"fieldtype"`
			Field      string  `json:"field"`
		} `json:"items"`
		Id int `json:"id"`
	} `json:"changelog"`
	Comment struct {
		Self   string `json:"self"`
		Id     string `json:"id"`
		Author struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				X16 string `json:"16x16"`
				X48 string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"author"`
		Body         string `json:"body"`
		UpdateAuthor struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				X16 string `json:"16x16"`
				X48 string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"updateAuthor"`
		Created string `json:"created"`
		Updated string `json:"updated"`
	} `json:"comment"`
	WebhookEvent string `json:"webhookEvent"`
}

/*
{
    "id": 2,
    "timestamp": 1525698237764,
    "issue": {
        "id":"99291",
        "self":"https://jira.atlassian.com/rest/api/2/issue/99291",
        "key":"JRA-20002",
        "fields":{
            "summary":"I feel the need for speed",
            "created":"2009-12-16T23:46:10.612-0600",
            "description":"Make the issue nav load 10x faster",
            "labels":["UI", "dialogue", "move"],
            "priority": "Minor"
        }
    },
    "user": {
        "self":"https://jira.atlassian.com/rest/api/2/user?username=brollins",
        "name":"brollins",
        "key":"brollins",
        "emailAddress":"bryansemail at atlassian dot com",
        "avatarUrls":{
            "16x16":"https://jira.atlassian.com/secure/useravatar?size=small&avatarId=10605",
            "48x48":"https://jira.atlassian.com/secure/useravatar?avatarId=10605"
        },
        "displayName":"Bryan Rollins [Atlassian]",
        "active" : "true"
    },
    "changelog": {
        "items": [
            {
                "toString": "A new summary.",
                "to": null,
                "fromString": "What is going on here?????",
                "from": null,
                "fieldtype": "jira",
                "field": "summary"
            },
            {
                "toString": "New Feature",
                "to": "2",
                "fromString": "Improvement",
                "from": "4",
                "fieldtype": "jira",
                "field": "issuetype"
            }
        ],
        "id": 10124
    },
    "comment" : {
        "self":"https://jira.atlassian.com/rest/api/2/issue/10148/comment/252789",
        "id":"252789",
        "author":{
            "self":"https://jira.atlassian.com/rest/api/2/user?username=brollins",
            "name":"brollins",
            "emailAddress":"bryansemail@atlassian.com",
            "avatarUrls":{
                "16x16":"https://jira.atlassian.com/secure/useravatar?size=small&avatarId=10605",
                "48x48":"https://jira.atlassian.com/secure/useravatar?avatarId=10605"
            },
            "displayName":"Bryan Rollins [Atlassian]",
            "active":true
        },
        "body":"Just in time for AtlasCamp!",
        "updateAuthor":{
            "self":"https://jira.atlassian.com/rest/api/2/user?username=brollins",
            "name":"brollins",
            "emailAddress":"brollins@atlassian.com",
            "avatarUrls":{
                "16x16":"https://jira.atlassian.com/secure/useravatar?size=small&avatarId=10605",
                "48x48":"https://jira.atlassian.com/secure/useravatar?avatarId=10605"
            },
            "displayName":"Bryan Rollins [Atlassian]",
            "active":true
        },
        "created":"2011-06-07T10:31:26.805-0500",
        "updated":"2011-06-07T10:31:26.805-0500"
    },
    "webhookEvent": "jira:issue_updated"
}
*/
