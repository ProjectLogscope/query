package authorization

import "fmt"

type authorization struct {
	defaultUserAccess   string
	mapUserAccessFields map[string][]string
}

func New(accessFieldsDefault []string, accessFieldsLimited []string) Authorization {
	return &authorization{
		mapUserAccessFields: map[string][]string{
			UserAccessDefault: accessFieldsDefault,
			UserAccessLimited: accessFieldsLimited,
		},
	}
}

func (*authorization) Validate(authorization string) error {
	if authorization != UserAccessDefault && authorization != UserAccessLimited {
		return fmt.Errorf("authorization value not recognized: %q", authorization)
	}
	return nil
}

func (a *authorization) Fields(userAccess string) []string {
	return a.mapUserAccessFields[userAccess]
}
