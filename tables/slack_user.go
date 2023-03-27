package tables

import (
	"context"
	"github.com/selefra/selefra-provider-slack/slack_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-slack/table_schema_generator"
)

type TableSlackUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableSlackUserGenerator{}

func (x *TableSlackUserGenerator) GetTableName() string {
	return "slack_user"
}

func (x *TableSlackUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableSlackUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableSlackUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableSlackUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			api, err := slack_client.Connect(ctx, taskClient.(*slack_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			users, err := api.GetUsersContext(ctx)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, user := range users {
				resultChannel <- user

			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableSlackUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableSlackUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("Indicates the display name that the user has chosen to identify themselves by in their workspace profile.").
			Extractor(column_value_extractor.StructSelector("Profile.DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bot_id").ColumnType(schema.ColumnTypeString).Description("If a bot user, this is the unique identifier of the bot.").
			Extractor(column_value_extractor.StructSelector("Profile.BotID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Description("True if the user has been deleted.").
			Extractor(column_value_extractor.StructSelector("Deleted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_24").ColumnType(schema.ColumnTypeString).Description("URL of the user profile image, size 24x24 pixels.").
			Extractor(column_value_extractor.StructSelector("Profile.Image24")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_32").ColumnType(schema.ColumnTypeString).Description("URL of the user profile image, size 32x32 pixels.").
			Extractor(column_value_extractor.StructSelector("Profile.Image32")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_restricted").ColumnType(schema.ColumnTypeBool).Description("Indicates whether or not the user is a guest user. Use in combination with the is_ultra_restricted field to check if the user is a single-channel guest user.").
			Extractor(column_value_extractor.StructSelector("IsRestricted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_name").ColumnType(schema.ColumnTypeString).Description("Last name of the user.").
			Extractor(column_value_extractor.StructSelector("Profile.LastName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("color").ColumnType(schema.ColumnTypeString).Description("Used in some clients to display a special username color.").
			Extractor(column_value_extractor.StructSelector("Color")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_2fa").ColumnType(schema.ColumnTypeBool).Description("True if two-factor authentication is enabled for the user.").
			Extractor(column_value_extractor.StructSelector("Has2FA")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_192").ColumnType(schema.ColumnTypeString).Description("URL of the user profile image, size 192x192 pixels.").
			Extractor(column_value_extractor.StructSelector("Profile.Image192")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_admin").ColumnType(schema.ColumnTypeBool).Description("True if the user is an administrator of the current workspace.").
			Extractor(column_value_extractor.StructSelector("IsAdmin")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_owner").ColumnType(schema.ColumnTypeBool).Description("True if the user is an owner of the current workspace.").
			Extractor(column_value_extractor.StructSelector("IsOwner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_expiration").ColumnType(schema.ColumnTypeTimestamp).Description("Expiration for the user status.").
			Extractor(column_value_extractor.StructSelector("Profile.StatusExpiration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tz_label").ColumnType(schema.ColumnTypeString).Description("Describes the commonly used name of the timezone.").
			Extractor(column_value_extractor.StructSelector("TZLabel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated").ColumnType(schema.ColumnTypeTimestamp).Description("Time when the user was last updated.").
			Extractor(column_value_extractor.StructSelector("Updated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_72").ColumnType(schema.ColumnTypeString).Description("URL of the user profile image, size 72x72 pixels.").
			Extractor(column_value_extractor.StructSelector("Profile.Image72")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_app_user").ColumnType(schema.ColumnTypeBool).Description("True if the user is an owner of the current workspace.").
			Extractor(column_value_extractor.StructSelector("IsAppUser")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_bot").ColumnType(schema.ColumnTypeBool).Description("True if the user is a bot.").
			Extractor(column_value_extractor.StructSelector("IsBot")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_primary_owner").ColumnType(schema.ColumnTypeBool).Description("True if the user is the primary owner of the current workspace.").
			Extractor(column_value_extractor.StructSelector("IsPrimaryOwner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tz").ColumnType(schema.ColumnTypeString).Description("A human-readable string for the geographic timezone-related region this user has specified in their account.").
			Extractor(column_value_extractor.StructSelector("TZ")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_ultra_restricted").ColumnType(schema.ColumnTypeBool).Description("Indicates whether or not the user is a single-channel guest.").
			Extractor(column_value_extractor.StructSelector("IsUltraRestricted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profile_fields").ColumnType(schema.ColumnTypeJSON).Description("Custom fields for the profile.").
			Extractor(column_value_extractor.StructSelector("Profile.Fields")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_text").ColumnType(schema.ColumnTypeString).Description("Status text the user has set.").
			Extractor(column_value_extractor.StructSelector("Profile.StatusText")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name_normalized").ColumnType(schema.ColumnTypeString).Description("The display name, but with any non-Latin characters filtered out.").
			Extractor(column_value_extractor.StructSelector("Profile.DisplayNameNormalized")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("skype").ColumnType(schema.ColumnTypeString).Description("Skype handle of the user.").
			Extractor(column_value_extractor.StructSelector("Profile.Skype")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_emoji").ColumnType(schema.ColumnTypeString).Description("Status emoji the user has set.").
			Extractor(column_value_extractor.StructSelector("Profile.StatusEmoji")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team_id").ColumnType(schema.ColumnTypeString).Description("The team workspace that the user is a member of.").
			Extractor(column_value_extractor.StructSelector("Profile.Team")).Build(),
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
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the user.").
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_app_id").ColumnType(schema.ColumnTypeString).Description("If an app user, then this is the unique identifier of the installed Slack application.").
			Extractor(column_value_extractor.StructSelector("Profile.ApiAppID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Email address of the user.").
			Extractor(column_value_extractor.StructSelector("Profile.Email")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("job_title").ColumnType(schema.ColumnTypeString).Description("Job title of the user.").
			Extractor(column_value_extractor.StructSelector("Profile.Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone").ColumnType(schema.ColumnTypeString).Description("Phone number of the user.").
			Extractor(column_value_extractor.StructSelector("Profile.Phone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("real_name_normalized").ColumnType(schema.ColumnTypeString).Description("The real_name field, but with any non-Latin characters filtered out.").
			Extractor(column_value_extractor.StructSelector("Profile.RealNameNormalized")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tz_offset").ColumnType(schema.ColumnTypeInt).Description("Indicates the number of seconds to offset UTC time by for this user's timezone.").
			Extractor(column_value_extractor.StructSelector("TZOffset")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_original").ColumnType(schema.ColumnTypeString).Description("URL of the user profile image, original size.").
			Extractor(column_value_extractor.StructSelector("Profile.ImageOriginal")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_invited_user").ColumnType(schema.ColumnTypeBool).Description("True if the user joined the workspace via an invite.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_stranger").ColumnType(schema.ColumnTypeBool).Description("If true, this user belongs to a different workspace than the one associated with your app's token, and isn't in any shared channels visible to your app. If false (or this field is not present), the user is either from the same workspace as associated with your app's token, or they are from a different workspace, but are in a shared channel that your app has access to. Read our shared channels docs for more detail.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locale").ColumnType(schema.ColumnTypeString).Description("IETF language code for the user's chosen display language.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("real_name").ColumnType(schema.ColumnTypeString).Description("The real name that the user specified in their workspace profile.").
			Extractor(column_value_extractor.StructSelector("Profile.RealName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_name").ColumnType(schema.ColumnTypeString).Description("First name of the user.").
			Extractor(column_value_extractor.StructSelector("Profile.FirstName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_48").ColumnType(schema.ColumnTypeString).Description("URL of the user profile image, size 48x48 pixels.").
			Extractor(column_value_extractor.StructSelector("Profile.Image48")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_512").ColumnType(schema.ColumnTypeString).Description("URL of the user profile image, size 512x512 pixels.").
			Extractor(column_value_extractor.StructSelector("Profile.Image512")).Build(),
	}
}

func (x *TableSlackUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
