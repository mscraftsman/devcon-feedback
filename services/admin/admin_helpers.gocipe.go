// generated by gocipe; DO NOT EDIT

package admin

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gosimple/slug"
	"github.com/mscraftsman/devcon-feedback/models"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type slugKeyType int

var slugKey slugKeyType

// NewBoolFilter returns a new filter for boolean types
func NewBoolFilter(filter *Filter) (*models.Filter, error) {
	value, err := strconv.ParseBool(filter.Value)

	if err != nil {
		return nil, err
	}

	return &models.Filter{Field: filter.Field, Operation: "=", Value: value}, nil
}

// NewStringFilter returns a new filter for text values
func NewStringFilter(filter *Filter) (*models.Filter, error) {
	switch filter.Operation {
	case "=":
		return &models.Filter{Field: filter.Field, FieldFunction: models.FilterFuncLower, Operation: "=", Value: strings.ToLower(filter.Value)}, nil
	case "!=":
		return &models.Filter{Field: filter.Field, FieldFunction: models.FilterFuncLower, Operation: "!=", Value: strings.ToLower(filter.Value)}, nil
	case "~":
		return &models.Filter{Field: filter.Field, FieldFunction: models.FilterFuncLower, Operation: "LIKE", Value: strings.ToLower(filter.Value)}, nil
	}

	return nil, ErrorInvalidOperation
}

// NewDateFilter returns a new filter to handle dates
func NewDateFilter(filter *Filter) (*models.Filter, error) {
	var (
		value time.Time
		err   error
	)

	if value, err = time.Parse(time.RFC3339, filter.Value); err != nil {
		return nil, err
	}

	switch filter.Operation {
	case "=":
		return &models.Filter{Field: filter.Field, Operation: "=", Value: value}, nil
	case ">":
		return &models.Filter{Field: filter.Field, Operation: ">", Value: value}, nil
	case ">=":
		return &models.Filter{Field: filter.Field, Operation: ">=", Value: value}, nil
	case "<":
		return &models.Filter{Field: filter.Field, Operation: "<", Value: value}, nil
	case "<=":
		return &models.Filter{Field: filter.Field, Operation: "<=", Value: value}, nil
	}

	return &models.Filter{Field: filter.Field, Operation: "LIKE", Value: value}, nil
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func setSlug(id string, value string, table string, field string) (string, error) {
	suffix := ""
	max := 10
	// Normalize input
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	value, _, _ = transform.String(t, value)
	for i := 0; i < max; i++ {
		if i >= max-1 {
			now := time.Now()
			suffix = fmt.Sprintf("%d", now.Nanosecond())
		} else if i > 0 {
			suffix = fmt.Sprintf("%02d", i)
		}

		tx, err := models.StartTransaction(context.Background())
		if err != nil {
			return "", err
		}
		defer tx.Rollback()

		slug := slug.Make(value + `-` + suffix)
		rows, err := tx.Query(`SELECT id FROM `+table+` WHERE slug = $1 AND id <> $2 LIMIT 1`, slug, id)
		if err != nil {
			continue
		}
		defer rows.Close()

		var exist bool
		for rows.Next() {
			exist = true
			break
		}

		if exist {
			continue
		}

		tx, err = models.StartTransaction(context.Background())
		if err != nil {
			return "", err
		}

		stmt, err := tx.Prepare(`UPDATE ` + table + ` SET ` + field + ` = $1 WHERE id = $2`)
		defer stmt.Close()
		if err != nil {
			continue
		}

		_, err = stmt.Exec(slug, id)
		if err != nil {
			continue
		}

		err = tx.Commit()
		if err != nil {
			continue
		}

		return slug, nil
	}

	return "", fmt.Errorf("Error setting slug")
}

func setSlugAsync(id string, value string, table string, field string) <-chan string {
	ch := make(chan string, 1)
	go func() {
		slug, err := setSlug(id, value, table, field)
		if err != nil {
			ch <- ""
			return
		}
		ch <- slug
	}()
	return ch
}
