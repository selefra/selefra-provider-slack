# Table: slack_file

## Columns

| Column Name         | Data Type | Uniq | Nullable | Description |
|---------------------|-----------|------|----------|-------------|
| id                  | string    | √  | X       |             |
| created             | timestamp | X   | √      |             |
| name                | string    | X   | √      |             |
| title               | string    | X   | √      |             |
| mimetype            | string    | X   | √      |             |
| image-exif_rotation | int       | X   | √      |             |
| file_type           | string    | X   | √      |             |
| pretty_type         | string    | X   | √      |             |
| user                | string    | X   | √      |             |
| mode                | string    | X   | √      |             |
| editable            | bool      | X   | √      |             |
| is_external         | bool      | X   | √      |             |
| external_type       | string    | X   | √      |             |
| size                | int       | X   | √      |             |
| url_private         | string    | X   | √      |             |
| edit_link           | string    | X   | √      |             |
| preview             | string    | X   | √      |             |
| is_public           | bool      | X   | √      |             |
| channels            | string[]  | X   | √      |             |
| groups              | string[]  | X   | √      |             |
| ims                 | string[]  | X   | √      |             |
| comments_count      | int       | X   | √      |             |
| num_stars           | int       | X   | √      |             |
| is_started          | bool      | X   | √      |             |