package main

import (
	"net/url"
	"strings"
)

// ConnectionDetails represents the connection details for a parsed URL.
type ConnectionDetails struct {
	Scheme   string
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

// Parse parses a connection string and returns the connection details.
func Parse(s string) (ConnectionDetails, error) {
	u, err := url.Parse(s)
	if err != nil {
		return ConnectionDetails{}, err
	}
	password, _ := u.User.Password()
	v := ConnectionDetails{
		Scheme:   u.Scheme,
		Username: u.User.Username(),
		Password: password,
		Host:     u.Hostname(),
		Port:     u.Port(),
		Name:     strings.TrimPrefix(u.Path, "/"),
	}
	return v, nil
}
