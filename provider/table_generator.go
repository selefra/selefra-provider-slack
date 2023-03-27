package provider

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-slack/table_schema_generator"
	"github.com/selefra/selefra-provider-slack/tables"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableSlackUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSlackConnectionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSlackConversationGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSlackEmojiGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSlackGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSlackFileGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableSlackTeamInfoGenerator{}),
	}
}
