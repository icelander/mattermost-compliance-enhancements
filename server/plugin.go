package main

import (
	"sync"

	"github.com/mattermost/mattermost-server/v6/plugin"
	"github.com/mattermost/mattermost-server/v6/model"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

func (p *Plugin) MessageWillBeUpdated(c *plugin.Context, newPost, oldPost *model.Post) (*model.Post, string) {
	configuration := p.getConfiguration()

	if configuration.disabled {
		return newPost, ""
	}

	// When editing a post, the userID for newPost and oldPost are the same so we look it up from the session
	currentSession, err := p.API.GetSession(c.SessionId)

	if err != nil {
		p.API.LogError("Could not get session, not updating editedBy value", "Session ID:", c.SessionId, "Error", err)
		return newPost, ""
	}

	updatingUserId := ""
	// If the post is being edited and the person editing isn't the original person
	if (newPost.UserId != currentSession.UserId) {
		p.API.LogInfo("Updating EditedBy Field", "updatingUserId", updatingUserId, "newPost.UserId", newPost.UserId)
		updatingUserId = currentSession.UserId	
	} else {
		updatingUserId = newPost.UserId
	}
	
	postProps := newPost.GetProps()
	postProps["editedBy"] = updatingUserId
	newPost.SetProps(postProps)
	
	return newPost, ""
}