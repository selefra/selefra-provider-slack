# Table: slack_conversation

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_member | bool | X | √ | If true, the user running this query is a member of this conversation. | 
| is_org_shared | bool | X | √ | If true, the conversation is shared between multiple workspaces within the same Enterprise Grid. | 
| purpose | string | X | √ | Purpose of the conversation / channel. | 
| purpose_last_set | timestamp | X | √ | Time when the purpose was last set. | 
| creator | string | X | √ | ID of the user who created the conversation. | 
| is_channel | bool | X | √ | If true, the conversation is a public channel inside the workspace. | 
| is_general | bool | X | √ | If true, this is the #general public channel (even if it's been renamed). | 
| is_im | bool | X | √ | If true, the conversation is a direct message between two individuals or a user and a bot. | 
| num_members | int | X | √ | Number of members in the conversation. Not set if the conversation is individual messages between fixed number of users. | 
| purpose_creator | string | X | √ | User who created the purpose for the conversation. | 
| topic | string | X | √ | Topic of the conversation / channel. | 
| created | timestamp | X | √ | Time when the conversation was created. | 
| is_archived | bool | X | √ | If true, the conversation has been archived. | 
| is_group | bool | X | √ | If true, the conversation is a private channel. | 
| is_private | bool | X | √ | If true, the conversation is privileged between two or more members. | 
| is_ext_shared | bool | X | √ | If true, the conversation is shared with an external workspace. | 
| is_mpim | bool | X | √ | If true, this is an unnamed private conversation between multiple users. | 
| topic_creator | string | X | √ | User who created the topic for the conversation. | 
| topic_last_set | timestamp | X | √ | Time when the topic was last set. | 
| name_normalized | string | X | √ | Name of the conversation normalized into simple ASCII characters. | 
| workspace_domain | string | X | √ | The domain name for the workspace. | 
| id | string | X | √ | ID of the conversation. | 
| name | string | X | √ | Name of the conversation. | 
| is_pending_ext_shared | bool | X | √ | If true, the conversation hopes is awaiting approval to become is_ext_shared. | 
| is_shared | bool | X | √ | If true, the conversation is shared across multiple workspaces. See also is_ext_shared. | 


