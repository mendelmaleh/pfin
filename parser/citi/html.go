package citi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

var (
	SelectorTxRow     = cascadia.MustCompile(".transaction-row.mobile")
	SelectorTxDesc    = cascadia.MustCompile("td > div > div.body > div.top > div.description")
	SelectorTxAmount  = cascadia.MustCompile("td > div > div.body > div.top > div.amount")
	SelectorTxDetails = cascadia.MustCompile("td.transaction-details-cell > div.transaction-details-wrapper > div.extended-descriptions >  div.extended-description-row")
)

func TxFromHTML(n *html.Node) (tx Transaction, err error) {
	if m := SelectorTxDesc.MatchFirst(n); m != nil {
		tx.Description = strings.TrimSpace(m.FirstChild.Data)
	} else {
		return tx, errors.New("no transaction description")
	}

	if m := SelectorTxAmount.MatchFirst(n); m != nil {
		s := strings.ReplaceAll(strings.TrimSpace(m.FirstChild.Data), ",", "")

		var negative bool
		if s[0] == '-' {
			s = s[2:]
			negative = true
		} else {
			s = s[1:]
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return tx, fmt.Errorf("couldn't parse %q to float", s)
		}

		if negative {
			f *= -1
		}

		tx.Raw.Amount = f
	} else {
		return tx, errors.New("no transaction amount")
	}

	// additional details
	for _, m := range SelectorTxDetails.MatchAll(n.NextSibling) {
		name := m.FirstChild.FirstChild.Data
		value := m.FirstChild.NextSibling.FirstChild.Data

		if name == "Purchased On" || name == "Posted On" {
			var t time.Time
			if len(strings.Fields(value)) > 3 {
				// TODO: fix timezone parsing
				t, err = time.Parse("Jan 02, 2006 03:04 PM ET", value)
			} else {
				t, err = time.Parse("Jan 02, 2006", value)
			}

			if err != nil {
				return tx, err
			}

			switch name {
			case "Purchased On":
				tx.Purchased = t
			case "Posted On":
				tx.Posted = t
			}

			continue
		}

		switch name {
		case "Cardmember Name":
			tx.Cardmember = value
		case "Purchase Method":
			tx.Method = value
		case "Spend Category":
			tx.Raw.Category = value
		case "Rewards":
			if strings.Fields(value)[0] == "N/A" {
				continue
			}

			i, err := strconv.Atoi(strings.Fields(value)[0])
			if err != nil {
				return tx, err
			}

			tx.Rewards = i
		case "Type":
			tx.Type = value
		default:
			err = fmt.Errorf("unknown name %q (value %q)", name, value)
			return
		}
	}

	return
}
