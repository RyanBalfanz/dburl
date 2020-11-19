package main

import (
	"fmt"
	"testing"
)

func ExampleParse() {
	d, err := Parse("postgres://username:password@host.test:5432/dbname")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", d)
	// Output:
	// {Scheme:postgres Username:username Password:password Host:host.test Port:5432 Name:dbname}
}

func TestParse_WithScheme(t *testing.T) {
	cases := []struct {
		connection string
		want       interface{}
	}{
		{
			connection: "postgres://",
			want: ConnectionDetails{
				Scheme:   "postgres",
				Username: "",
				Password: "",
				Host:     "",
				Port:     "",
				Name:     "",
			},
		},
	}
	for _, c := range cases {
		got, err := Parse(c.connection)
		if err != nil {
			t.FailNow()
		}
		if got != c.want {
			t.Errorf("Parse(%s) got %+v, wanted %+v", c.connection, got, c.want)
		}
	}
}

func TestParse_WithHost(t *testing.T) {
	cases := []struct {
		connection string
		want       interface{}
	}{
		{
			connection: "postgres://localhost",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "localhost", Port: "", Name: ""},
		},
		{
			connection: "postgres://localhost:",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "localhost", Port: "", Name: ""},
		},
		{
			connection: "postgres://localhost:/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "localhost", Port: "", Name: ""},
		},
		{
			connection: "postgres://host.test",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://host.test:",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://host.test:/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://127.0.0.1",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "127.0.0.1", Port: "", Name: ""},
		},
		{
			connection: "postgres://0.0.0.0",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "0.0.0.0", Port: "", Name: ""},
		},
	}
	for _, c := range cases {
		got, err := Parse(c.connection)
		if err != nil {
			t.FailNow()
		}
		if got != c.want {
			t.Errorf("Parse(%s) got %+v, wanted %+v", c.connection, got, c.want)
		}
	}
}

func TestParse_WithDatabaseName(t *testing.T) {
	cases := []struct {
		connection string
		want       interface{}
	}{
		{
			connection: "postgres:///dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "", Port: "", Name: "dbname"},
		},
		{
			connection: "postgres://:/dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "", Port: "", Name: "dbname"},
		},
		{
			connection: "postgres://host.test/dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "host.test", Port: "", Name: "dbname"},
		},
		{
			connection: "postgres://host.test:/dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "host.test", Port: "", Name: "dbname"},
		},
		{
			connection: "postgres://host.test:5432/dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "", Password: "", Host: "host.test", Port: "5432", Name: "dbname"},
		},
		{
			connection: "postgres://username@host.test:5432/dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "host.test", Port: "5432", Name: "dbname"},
		},
		{
			connection: "postgres://username:password@host.test:5432/dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "host.test", Port: "5432", Name: "dbname"},
		},
	}
	for _, c := range cases {
		got, err := Parse(c.connection)
		if err != nil {
			t.FailNow()
		}
		if got != c.want {
			t.Errorf("Parse(%s) got %+v, wanted %+v", c.connection, got, c.want)
		}
	}
}

func TestParse_WithUsername(t *testing.T) {
	cases := []struct {
		connection string
		want       interface{}
	}{
		{
			connection: "postgres://username@",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "", Port: "", Name: ""},
		},
		{
			connection: "postgres://username@host.test",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://username@host.test/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://username@host.test:",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://username@host.test:/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://username@:",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "", Port: "", Name: ""},
		},
		{
			connection: "postgres://username@:/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "", Port: "", Name: ""},
		},
		{
			connection: "postgres://username@:5432",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "", Port: "5432", Name: ""},
		},
		{
			connection: "postgres://username@:5432/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "", Port: "5432", Name: ""},
		},
		{
			connection: "postgres://username@:5432/dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "", Host: "", Port: "5432", Name: "dbname"},
		},
	}
	for _, c := range cases {
		got, err := Parse(c.connection)
		if err != nil {
			t.FailNow()
		}
		if got != c.want {
			t.Errorf("Parse(%s) got %+v, wanted %+v", c.connection, got, c.want)
		}
	}
}

func TestParse_WithUsernameAndPassword(t *testing.T) {
	cases := []struct {
		connection string
		want       interface{}
	}{
		{
			connection: "postgres://username:password@",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "", Port: "", Name: ""},
		},
		{
			connection: "postgres://username:password@host.test",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://username:password@host.test/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://username:password@host.test:",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://username:password@host.test:/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "host.test", Port: "", Name: ""},
		},
		{
			connection: "postgres://username:password@:",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "", Port: "", Name: ""},
		},
		{
			connection: "postgres://username:password@:/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "", Port: "", Name: ""},
		},
		{
			connection: "postgres://username:password@:5432",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "", Port: "5432", Name: ""},
		},
		{
			connection: "postgres://username:password@:5432/",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "", Port: "5432", Name: ""},
		},
		{
			connection: "postgres://username:password@:5432/dbname",
			want:       ConnectionDetails{Scheme: "postgres", Username: "username", Password: "password", Host: "", Port: "5432", Name: "dbname"},
		},
	}
	for _, c := range cases {
		got, err := Parse(c.connection)
		if err != nil {
			t.FailNow()
		}
		if got != c.want {
			t.Errorf("Parse(%s) got %+v, wanted %+v", c.connection, got, c.want)
		}
	}
}
