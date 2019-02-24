package main

import (
	"encoding/json"
	"fmt"

	"github.com/sauerbraten/maitred/pkg/auth"
	"github.com/sauerbraten/waiter/pkg/definitions/role"
)

type User struct {
	Name      string         `json:"name"`
	PublicKey auth.PublicKey `json:"public_key"`
	Role      role.ID        `json:"-"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	proxy := struct {
		User
		Role string `json:"role"`
	}{
		User: *u,
		Role: u.Role.String(),
	}
	return json.Marshal(proxy)
}

func (u *User) UnmarshalJSON(data []byte) error {
	proxy := &struct {
		Name      string
		PublicKey auth.PublicKey `json:"public_key"`
		Role      string         `json:"role"`
	}{}
	err := json.Unmarshal(data, proxy)
	if err != nil {
		return err
	}
	u.Name, u.PublicKey = proxy.Name, proxy.PublicKey
	u.Role = role.Parse(proxy.Role)
	if u.Role == -1 {
		return fmt.Errorf("auth: could not parse '%s' as role", proxy.Role)
	}
	return nil
}
