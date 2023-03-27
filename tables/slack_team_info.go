package tables

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-slack/slack_client"
	"github.com/selefra/selefra-provider-slack/table_schema_generator"
)

type TableSlackTeamInfoGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSlackTeamInfoGenerator{}

func (x *TableSlackTeamInfoGenerator) GetTableName() string {
	return "slack_team"
}

func (x *TableSlackTeamInfoGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSlackTeamInfoGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSlackTeamInfoGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSlackTeamInfoGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			api, err := slack_client.Connect(ctx, taskClient.(*slack_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			team, err := api.GetTeamInfo()
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- team
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableSlackTeamInfoGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSlackTeamInfoGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_domain").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EmailDomain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("icon").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableSlackTeamInfoGenerator) GetSubTables() []*schema.Table {
	return nil
}
