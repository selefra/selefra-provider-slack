package tables

import (
	"context"
	"github.com/selefra/selefra-provider-slack/slack_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-slack/table_schema_generator"
)

type TableSlackEmojiGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSlackEmojiGenerator{}

func (x *TableSlackEmojiGenerator) GetTableName() string {
	return "slack_emoji"
}

func (x *TableSlackEmojiGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSlackEmojiGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSlackEmojiGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSlackEmojiGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			api, err := slack_client.Connect(ctx, taskClient.(*slack_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			emojis, err := api.GetEmojiContext(ctx)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for name, url := range emojis {
				resultChannel <- slackEmoji{Name: name, URL: url}

			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type slackEmoji struct {
	Name string
	URL  string
}

func (x *TableSlackEmojiGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSlackEmojiGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the emoji, used in message text.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URL")).Description("URL of the emoji image.").Build(),
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
	}
}

func (x *TableSlackEmojiGenerator) GetSubTables() []*schema.Table {
	return nil
}
