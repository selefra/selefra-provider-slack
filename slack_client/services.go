package slack_client

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/slack-go/slack"
)

func Connect(_ context.Context, config *Config) (*slack.Client, error) {
	api := slack.New(config.Token, slack.OptionDebug(false))
	return api, nil
}

type WorkspaceInfo struct {
	Name        string
	ID          string
	Domain      string
	EmailDomain string
}

func GetCommonColumns(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	var workspaceInfo *WorkspaceInfo
	workspace, err := getTeamInfo(ctx, taskClient)
	if err != nil {
		return nil, err
	}

	workspaceInfo = &WorkspaceInfo{
		Name:        workspace.Name,
		ID:          workspace.ID,
		Domain:      workspace.Domain,
		EmailDomain: workspace.EmailDomain,
	}

	return workspaceInfo, nil
}

func getTeamInfo(ctx context.Context, taskClient any) (*slack.TeamInfo, error) {
	api, err := Connect(ctx, taskClient.(*Client).Config)
	if err != nil {
		return nil, err
	}

	data, err := api.GetTeamInfo()
	if err != nil {
		return nil, err
	}

	return data, nil
}
