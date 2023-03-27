package tables

import (
	"context"

	"github.com/selefra/selefra-provider-slack/slack_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-slack/table_schema_generator"
	"github.com/slack-go/slack"
)

type TableSlackGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSlackGroupGenerator{}

func (x *TableSlackGroupGenerator) GetTableName() string {
	return "slack_group"
}

func (x *TableSlackGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSlackGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSlackGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSlackGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			api, err := slack_client.Connect(ctx, taskClient.(*slack_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			groups, err := api.GetUserGroupsContext(ctx, slack.GetUserGroupsOptionIncludeCount(true), slack.GetUserGroupsOptionIncludeDisabled(true), slack.GetUserGroupsOptionIncludeUsers(true))
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, group := range groups {
				resultChannel <- group

			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableSlackGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSlackGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("team_id").ColumnType(schema.ColumnTypeString).Description("Team ID the group is defined in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_external").ColumnType(schema.ColumnTypeBool).Description("True if the group is external facing.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_create").ColumnType(schema.ColumnTypeTimestamp).Description("Date when the group was created.").
			Extractor(column_value_extractor.StructSelector("DateCreate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_update").ColumnType(schema.ColumnTypeTimestamp).Description("Date when the group was last updated.").
			Extractor(column_value_extractor.StructSelector("DateUpdate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_type").ColumnType(schema.ColumnTypeString).Description("The auto_type parameter can be admin for a Workspace Admins group, owner for a Workspace Owners group or null for a custom group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_count").ColumnType(schema.ColumnTypeInt).Description("Number of users in the group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeJSON).Description("List of users (IDs) in the group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_user_group").ColumnType(schema.ColumnTypeBool).Description("True if this is a user group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Description("User who created the group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("prefs").ColumnType(schema.ColumnTypeJSON).Description("The prefs parameter contains default channels and groups (private channels) that members of this group will be invited to upon joining.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workspace_domain").ColumnType(schema.ColumnTypeString).Description("The domain name for the workspace.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 001
				r, err := slack_client.GetCommonColumns(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Domain")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_delete").ColumnType(schema.ColumnTypeTimestamp).Description("Date when the group was deleted.").
			Extractor(column_value_extractor.StructSelector("DateDelete")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_by").ColumnType(schema.ColumnTypeString).Description("User who last updated the group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of the group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("handle").ColumnType(schema.ColumnTypeString).Description("The handle parameter indicates the value used to notify group members via a mention without a leading @ sign.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_by").ColumnType(schema.ColumnTypeString).Description("User who deleted the group.").Build(),
	}
}

func (x *TableSlackGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
