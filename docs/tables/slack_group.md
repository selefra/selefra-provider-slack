# Table: slack_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| team_id | string | X | √ | Team ID the group is defined in. | 
| is_external | bool | X | √ | True if the group is external facing. | 
| date_create | timestamp | X | √ | Date when the group was created. | 
| date_update | timestamp | X | √ | Date when the group was last updated. | 
| auto_type | string | X | √ | The auto_type parameter can be admin for a Workspace Admins group, owner for a Workspace Owners group or null for a custom group. | 
| user_count | int | X | √ | Number of users in the group. | 
| users | json | X | √ | List of users (IDs) in the group. | 
| is_user_group | bool | X | √ | True if this is a user group. | 
| created_by | string | X | √ | User who created the group. | 
| prefs | json | X | √ | The prefs parameter contains default channels and groups (private channels) that members of this group will be invited to upon joining. | 
| workspace_domain | string | X | √ | The domain name for the workspace. | 
| date_delete | timestamp | X | √ | Date when the group was deleted. | 
| updated_by | string | X | √ | User who last updated the group. | 
| id | string | X | √ | ID of the group. | 
| name | string | X | √ | Name of the group. | 
| description | string | X | √ | Description of the group. | 
| handle | string | X | √ | The handle parameter indicates the value used to notify group members via a mention without a leading @ sign. | 
| deleted_by | string | X | √ | User who deleted the group. | 


