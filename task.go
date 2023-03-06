package BitrixGo

type Task struct {
	Title          string `bitrix:"fields[TITLE]"`
	Description    string `bitrix:"fields[DESCRIPTION]"`
	Responsible_id int    `bitrix:"fields[RESPONSIBLE_ID]"`
	Accomplices    int    `bitrix:"fields[ACCOMPLICES]"`
	Auditors       int    `bitrix:"fields[AUDITORS]"`
	Parent_id      int    `bitrix:"fields[PARENT_ID]"`
	Group_id       int    `bitrix:"fields[GROUP_ID]"`
	Created_by     int    `bitrix:"fields[CREATED_BY]"`
}
