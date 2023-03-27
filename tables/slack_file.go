package tables

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-slack/slack_client"
	"github.com/slack-go/slack"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-slack/table_schema_generator"
)

type TableSlackFileGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSlackFileGenerator{}

func (x *TableSlackFileGenerator) GetTableName() string {
	return "slack_file"
}

func (x *TableSlackFileGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSlackFileGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSlackFileGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSlackFileGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			api, err := slack_client.Connect(ctx, taskClient.(*slack_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			files, _, err := api.ListFilesContext(ctx, slack.ListFilesParameters{
				Limit: 1000,
				Types: "all",
			})
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, file := range files {
				resultChannel <- file
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableSlackFileGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSlackFileGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("file id").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeInt).Description("file crated at").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("file name").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("file title").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mimetype").ColumnType(schema.ColumnTypeString).Description("file mime type").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_exif_rotation").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ImageExifRotation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("file_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Filetype")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pretty_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrettyType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mode").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("editable").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_external").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsExternal")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExternalType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url_private").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URLPrivate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url_private_download").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URLPrivateDownload")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("edit_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EditLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preview").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_public").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsPublic")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("channels").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ims").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("IMs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comments_count").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("CommentsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_stars").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("NumStars")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_starred").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsStarred")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shares").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("initial_comment").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InitialComment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permalink_public").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_url_shared").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("PublicURLShared")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permalink").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableSlackFileGenerator) GetSubTables() []*schema.Table {
	return nil
}
