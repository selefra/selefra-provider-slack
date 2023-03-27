# Table: slack_access_log

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| date_first | timestamp | X | √ | Date of the first login in a sequence from this device. | 
| isp | string | X | √ | ISP the login originated from, if available. Often null. | 
| region | string | X | √ | Region the login originated from, if available. Often null. | 
| workspace_domain | string | X | √ | The domain name for the workspace. | 
| user_id | string | X | √ | Unique identifier of the user | 
| user_name | string | X | √ | Name of the user. | 
| ip | string | X | √ | IP address the login came from. | 
| count | int | X | √ | Number of sequential logins from this device. | 
| country | string | X | √ | Country the login originated from, if available. Often null. | 
| date_last | timestamp | X | √ | Date of the last login in a sequence from this device. | 
| user_agent | string | X | √ | User agent of the device used for login. | 


