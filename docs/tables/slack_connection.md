# Table: slack_connection

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enterprise_id | string | X | √ | ID of the enterprise grid. null if not an enterprise workspace. | 
| bot_id | string | X | √ | ID of the bot making the connection. null if not a bot. | 
| workspace_domain | string | X | √ | The domain name for the workspace. | 
| url | string | X | √ | URL of the workspace. | 
| team | string | X | √ | Name of the workspace team. | 
| user | string | X | √ | Name of the user making the connection. | 
| team_id | string | X | √ | ID of the workspace team. | 
| user_id | string | X | √ | ID of the user making the connection. | 


