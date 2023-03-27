package tables

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-slack/slack_client"
	"github.com/selefra/selefra-provider-slack/table_schema_generator"
	"github.com/slack-go/slack"
)

type TableSlackConversationGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSlackConversationGenerator{}

func (x *TableSlackConversationGenerator) GetTableName() string {
	return "slack_conversation"
}

func (x *TableSlackConversationGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSlackConversationGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSlackConversationGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSlackConversationGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			api, err := slack_client.Connect(ctx, taskClient.(*slack_client.Client).Config)

			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := &slack.GetConversationsParameters{Limit: 1000, Types: []string{"public_channel", "private_channel", "im", "mpim"}}

			for {
				conversations, cursor, err := api.GetConversations(opts)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, conversation := range conversations {
					resultChannel <- conversation

				}
				if cursor == "" {
					break
				}
				opts.Cursor = cursor
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableSlackConversationGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSlackConversationGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("is_member").ColumnType(schema.ColumnTypeBool).Description("If true, the user running this query is a member of this conversation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_org_shared").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation is shared between multiple workspaces within the same Enterprise Grid.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purpose").ColumnType(schema.ColumnTypeString).Description("Purpose of the conversation / channel.").
			Extractor(column_value_extractor.StructSelector("Purpose.Value")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purpose_last_set").ColumnType(schema.ColumnTypeTimestamp).Description("Time when the purpose was last set.").
			Extractor(column_value_extractor.StructSelector("Purpose.LastSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator").ColumnType(schema.ColumnTypeString).Description("ID of the user who created the conversation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_channel").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation is a public channel inside the workspace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_general").ColumnType(schema.ColumnTypeBool).Description("If true, this is the #general public channel (even if it's been renamed).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_im").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation is a direct message between two individuals or a user and a bot.").
			Extractor(column_value_extractor.StructSelector("IsIM")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_members").ColumnType(schema.ColumnTypeInt).Description("Number of members in the conversation. Not set if the conversation is individual messages between fixed number of users.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purpose_creator").ColumnType(schema.ColumnTypeString).Description("User who created the purpose for the conversation.").
			Extractor(column_value_extractor.StructSelector("Purpose.Creator")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("topic").ColumnType(schema.ColumnTypeString).Description("Topic of the conversation / channel.").
			Extractor(column_value_extractor.StructSelector("Topic.Value")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Time when the conversation was created.").
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_archived").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation has been archived.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_group").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation is a private channel.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_private").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation is privileged between two or more members.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_ext_shared").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation is shared with an external workspace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_mpim").ColumnType(schema.ColumnTypeBool).Description("If true, this is an unnamed private conversation between multiple users.").
			Extractor(column_value_extractor.StructSelector("IsMpIM")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("topic_creator").ColumnType(schema.ColumnTypeString).Description("User who created the topic for the conversation.").
			Extractor(column_value_extractor.StructSelector("Topic.Creator")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("topic_last_set").ColumnType(schema.ColumnTypeTimestamp).Description("Time when the topic was last set.").
			Extractor(column_value_extractor.StructSelector("Topic.LastSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name_normalized").ColumnType(schema.ColumnTypeString).Description("Name of the conversation normalized into simple ASCII characters.").Build(),
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
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the conversation.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the conversation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_pending_ext_shared").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation hopes is awaiting approval to become is_ext_shared.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_shared").ColumnType(schema.ColumnTypeBool).Description("If true, the conversation is shared across multiple workspaces. See also is_ext_shared.").Build(),
	}
}

func (x *TableSlackConversationGenerator) GetSubTables() []*schema.Table {
	return nil
}
