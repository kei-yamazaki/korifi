package payloads

import (
	"net/url"
	"strconv"
	"strings"

	"code.cloudfoundry.org/korifi/api/repositories"
)

type TaskCreate struct {
	Command string `json:"command" validate:"required"`
}

func (p TaskCreate) ToMessage(appRecord repositories.AppRecord) repositories.CreateTaskMessage {
	return repositories.CreateTaskMessage{
		Command:   p.Command,
		SpaceGUID: appRecord.SpaceGUID,
		AppGUID:   appRecord.GUID,
	}
}

type TaskList struct {
	SequenceIDs []int64
}

func (t *TaskList) ToMessage() repositories.ListTaskMessage {
	return repositories.ListTaskMessage{
		SequenceIDs: t.SequenceIDs,
	}
}

func (t *TaskList) SupportedKeys() []string {
	return []string{"sequence_ids"}
}

func (a *TaskList) DecodeFromURLValues(values url.Values) error {
	idsStr := values.Get("sequence_ids")

	var ids []int64
	for _, idStr := range strings.Split(idsStr, ",") {
		if idStr == "" {
			continue
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}

	a.SequenceIDs = ids
	return nil
}
