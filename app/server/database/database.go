package database

import (
	"fmt"

	"github.com/cosasdepuma/elliot/app/config"
	"github.com/go-redis/redis"
)

// === STRUCTURES ===

type Database struct {
	client *redis.Client
	data   *DBSchema
}

// === CONSTRUCTOR ===

func NewDatabase() *Database {
	return &Database{
		client: redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
			Password: config.DBPass,
			DB:       config.DBIndex,

			MaxRetries: 3,
		}),
		data: new(DBSchema),
	}
}

// === GENERAL ===

func (db Database) StoredData() *DBSchema {
	return db.data
}

func (db *Database) Purge() {
	db.client.FlushDB()
	db.data = new(DBSchema)
}

// === SETTERS ===

// === Domain ===

func (db *Database) SetDomain(value string) {
	db.data.Domain.Value = value
	db.client.Set("domain", value, 0)
}

func (db *Database) SetDomainIPv4(value string) {
	db.data.Domain.IPv4 = value
	db.client.Set("domain:ipv4", value, 0)
}

func (db *Database) SetDomainIPv6(value string) {
	db.data.Domain.IPv6 = value
	db.client.Set("domain:ipv6", value, 0)
}

func (db *Database) SetDomainSubdomains(value []string) {
	db.data.Domain.Subdomains = value
	db.client.RPush("domain:subdomains", value)
}

// === Domain:Whois ===

func (db *Database) SetDomainWhoisTLD(value string) {
	db.data.Domain.Whois.TLD = value
	db.client.Set("domain:whois:tld", value, 0)
}

func (db *Database) SetDomainWhoisStatus(value bool) {
	db.data.Domain.Whois.Status = value
	data := "false"
	if value {
		data = "true"
	}
	db.client.Set("domain:whois:status", data, 0)
}

func (db *Database) SetDomainWhoisCreated(value string) {
	db.data.Domain.Whois.Created = value
	db.client.Set("domain:whois:created", value, 0)
}

func (db *Database) SetDomainWhoisChanged(value string) {
	db.data.Domain.Whois.Changed = value
	db.client.Set("domain:whois:changed", value, 0)
}

func (db *Database) SetDomainWhoisPhones(value []string) {
	db.data.Domain.Whois.Phones = value
	db.client.RPush("domain:whois:phones", value)
}

func (db *Database) SetDomainWhoisEmails(value []string) {
	db.data.Domain.Whois.Emails = value
	db.client.RPush("domain:whois:emails", value)
}

// === Domain:Web ===

func (db *Database) SetDomainWebUrl(value string) {
	db.data.Domain.Web.Url = value
	db.client.Set("domain:web:url", value, 0)
}

func (db *Database) SetDomainWebServer(value string) {
	db.data.Domain.Web.Server = value
	db.client.Set("domain:web:server", value, 0)
}

func (db *Database) SetDomainWebRating(value string) {
	db.data.Domain.Web.Rating = value
	db.client.Set("domain:web:rating", value, 0)
}

func (db *Database) SetDomainWebRedirects(value []string) {
	db.data.Domain.Web.Redirects = value
	db.client.RPush("domain:web:redirects", value)
}

func (db *Database) SetDomainWebLinks(value []string) {
	db.data.Domain.Web.Links = value
	db.client.RPush("domain:web:links", value)
}

func (db *Database) SetDomainWebJS(value []string) {
	db.data.Domain.Web.Js = value
	db.client.RPush("domain:web:js", value)
}

// === GETTERS ===

// === Domain ===

func (db *Database) GetDomain() {
	value, err := db.client.Get("domain").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.Value = value
}

func (db *Database) GetDomainIPv4() {
	value, err := db.client.Get("domain:ipv4").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.IPv4 = value
}

func (db *Database) GetDomainIPv6() {
	value, err := db.client.Get("domain:ipv6").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.IPv6 = value
}

func (db *Database) GetDomainSubdomains() {
	value, err := db.client.LRange("domain:subdomains", 0, -1).Result()
	if err != nil {
		value = make([]string, 0)
	}
	db.data.Domain.Subdomains = value
}

// === Domain:Whois ===

func (db *Database) GetDomainWhoisTLD() {
	value, err := db.client.Get("domain:whois:tld").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.Whois.TLD = value
}

func (db *Database) GetDomainWhoisStatus() {
	value, err := db.client.Get("domain:whois:tld").Result()
	data := false
	if err == nil && value == "true" {
		data = true
	}
	db.data.Domain.Whois.Status = data
}

func (db *Database) GetDomainWhoisCreated() {
	value, err := db.client.Get("domain:whois:created").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.Whois.Created = value
}

func (db *Database) GetDomainWhoisChanged() {
	value, err := db.client.Get("domain:whois:changed").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.Whois.Changed = value
}

func (db *Database) GetDomainWhoisPhones() {
	value, err := db.client.LRange("domain:whois:phones", 0, -1).Result()
	if err != nil {
		value = make([]string, 0)
	}
	db.data.Domain.Whois.Phones = value
}

func (db *Database) GetDomainWhoisEmails() {
	value, err := db.client.LRange("domain:whois:emails", 0, -1).Result()
	if err != nil {
		value = make([]string, 0)
	}
	db.data.Domain.Whois.Emails = value
}

// === Domain:Web ===

func (db *Database) GetDomainWebUrl() {
	value, err := db.client.Get("domain:web:url").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.Web.Url = value
}

func (db *Database) GetDomainWebRating() {
	value, err := db.client.Get("domain:web:rating").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.Web.Rating = value
}

func (db *Database) GetDomainWebServer() {
	value, err := db.client.Get("domain:web:server").Result()
	if err != nil {
		value = ""
	}
	db.data.Domain.Web.Server = value
}

func (db *Database) GetDomainWebRedirects() {
	value, err := db.client.LRange("domain:web:redirects", 0, -1).Result()
	if err != nil {
		value = make([]string, 0)
	}
	db.data.Domain.Web.Redirects = value
}

func (db *Database) GetDomainWebJS() {
	value, err := db.client.LRange("domain:web:js", 0, -1).Result()
	if err != nil {
		value = make([]string, 0)
	}
	db.data.Domain.Web.Js = value
}

func (db *Database) GetDomainWebLinks() {
	value, err := db.client.LRange("domain:web:links", 0, -1).Result()
	if err != nil {
		value = make([]string, 0)
	}
	db.data.Domain.Web.Links = value
}