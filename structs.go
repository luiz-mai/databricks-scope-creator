package main

type ScopeSecrets map[string]string
type ScopeACLs map[string]string

type ScopeProperties struct {
	Secrets ScopeSecrets `json:"secrets"`
	ACLs    ScopeACLs    `json:"acls"`
}

type Scopes map[string]ScopeProperties
