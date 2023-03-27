package tables

import (
	"context"
	"github.com/selefra/selefra-provider-slack/slack_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-slack/table_schema_generator"
)

type TableSlackConnectionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSlackConnectionGenerator{}

func (x *TableSlackConnectionGenerator) GetTableName() string {
	return "slack_connection"
}

func (x *TableSlackConnectionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSlackConnectionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSlackConnectionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSlackConnectionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			api, err := slack_client.Connect(ctx, taskClient.(*slack_client.Client).Config)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			conn, err := api.AuthTestContext(ctx)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- conn
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableSlackConnectionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSlackConnectionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("enterprise_id").ColumnType(schema.ColumnTypeString).Description("ID of the enterprise grid. null if not an enterprise workspace.").
			Extractor(column_value_extractor.StructSelector("EnterpriseID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bot_id").ColumnType(schema.ColumnTypeString).Description("ID of the bot making the connection. null if not a bot.").
			Extractor(column_value_extractor.StructSelector("BotID")).Build(),
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
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("URL of the workspace.").
			Extractor(column_value_extractor.StructSelector("URL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team").ColumnType(schema.ColumnTypeString).Description("Name of the workspace team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user").ColumnType(schema.ColumnTypeString).Description("Name of the user making the connection.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team_id").ColumnType(schema.ColumnTypeString).Description("ID of the workspace team.").
			Extractor(column_value_extractor.StructSelector("TeamID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("ID of the user making the connection.").
			Extractor(column_value_extractor.StructSelector("UserID")).Build(),
	}
}

func (x *TableSlackConnectionGenerator) GetSubTables() []*schema.Table {
	return nil
}
