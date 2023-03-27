# Table: slack_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ | Indicates the display name that the user has chosen to identify themselves by in their workspace profile. | 
| bot_id | string | X | √ | If a bot user, this is the unique identifier of the bot. | 
| deleted | bool | X | √ | True if the user has been deleted. | 
| image_24 | string | X | √ | URL of the user profile image, size 24x24 pixels. | 
| image_32 | string | X | √ | URL of the user profile image, size 32x32 pixels. | 
| is_restricted | bool | X | √ | Indicates whether or not the user is a guest user. Use in combination with the is_ultra_restricted field to check if the user is a single-channel guest user. | 
| last_name | string | X | √ | Last name of the user. | 
| color | string | X | √ | Used in some clients to display a special username color. | 
| has_2fa | bool | X | √ | True if two-factor authentication is enabled for the user. | 
| image_192 | string | X | √ | URL of the user profile image, size 192x192 pixels. | 
| is_admin | bool | X | √ | True if the user is an administrator of the current workspace. | 
| is_owner | bool | X | √ | True if the user is an owner of the current workspace. | 
| status_expiration | timestamp | X | √ | Expiration for the user status. | 
| tz_label | string | X | √ | Describes the commonly used name of the timezone. | 
| updated | timestamp | X | √ | Time when the user was last updated. | 
| image_72 | string | X | √ | URL of the user profile image, size 72x72 pixels. | 
| is_app_user | bool | X | √ | True if the user is an owner of the current workspace. | 
| is_bot | bool | X | √ | True if the user is a bot. | 
| is_primary_owner | bool | X | √ | True if the user is the primary owner of the current workspace. | 
| tz | string | X | √ | A human-readable string for the geographic timezone-related region this user has specified in their account. | 
| is_ultra_restricted | bool | X | √ | Indicates whether or not the user is a single-channel guest. | 
| profile_fields | json | X | √ | Custom fields for the profile. | 
| status_text | string | X | √ | Status text the user has set. | 
| display_name_normalized | string | X | √ | The display name, but with any non-Latin characters filtered out. | 
| skype | string | X | √ | Skype handle of the user. | 
| status_emoji | string | X | √ | Status emoji the user has set. | 
| team_id | string | X | √ | The team workspace that the user is a member of. | 
| workspace_domain | string | X | √ | The domain name for the workspace. | 
| id | string | X | √ | Unique identifier for the user. | 
| api_app_id | string | X | √ | If an app user, then this is the unique identifier of the installed Slack application. | 
| email | string | X | √ | Email address of the user. | 
| job_title | string | X | √ | Job title of the user. | 
| phone | string | X | √ | Phone number of the user. | 
| real_name_normalized | string | X | √ | The real_name field, but with any non-Latin characters filtered out. | 
| tz_offset | int | X | √ | Indicates the number of seconds to offset UTC time by for this user's timezone. | 
| image_original | string | X | √ | URL of the user profile image, original size. | 
| is_invited_user | bool | X | √ | True if the user joined the workspace via an invite. | 
| is_stranger | bool | X | √ | If true, this user belongs to a different workspace than the one associated with your app's token, and isn't in any shared channels visible to your app. If false (or this field is not present), the user is either from the same workspace as associated with your app's token, or they are from a different workspace, but are in a shared channel that your app has access to. Read our shared channels docs for more detail. | 
| locale | string | X | √ | IETF language code for the user's chosen display language. | 
| real_name | string | X | √ | The real name that the user specified in their workspace profile. | 
| first_name | string | X | √ | First name of the user. | 
| image_48 | string | X | √ | URL of the user profile image, size 48x48 pixels. | 
| image_512 | string | X | √ | URL of the user profile image, size 512x512 pixels. | 


