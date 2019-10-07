// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Episodes", testEpisodes)
	t.Run("Quotes", testQuotes)
}

func TestDelete(t *testing.T) {
	t.Run("Episodes", testEpisodesDelete)
	t.Run("Quotes", testQuotesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Episodes", testEpisodesQueryDeleteAll)
	t.Run("Quotes", testQuotesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Episodes", testEpisodesSliceDeleteAll)
	t.Run("Quotes", testQuotesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Episodes", testEpisodesExists)
	t.Run("Quotes", testQuotesExists)
}

func TestFind(t *testing.T) {
	t.Run("Episodes", testEpisodesFind)
	t.Run("Quotes", testQuotesFind)
}

func TestBind(t *testing.T) {
	t.Run("Episodes", testEpisodesBind)
	t.Run("Quotes", testQuotesBind)
}

func TestOne(t *testing.T) {
	t.Run("Episodes", testEpisodesOne)
	t.Run("Quotes", testQuotesOne)
}

func TestAll(t *testing.T) {
	t.Run("Episodes", testEpisodesAll)
	t.Run("Quotes", testQuotesAll)
}

func TestCount(t *testing.T) {
	t.Run("Episodes", testEpisodesCount)
	t.Run("Quotes", testQuotesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Episodes", testEpisodesHooks)
	t.Run("Quotes", testQuotesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Episodes", testEpisodesInsert)
	t.Run("Episodes", testEpisodesInsertWhitelist)
	t.Run("Quotes", testQuotesInsert)
	t.Run("Quotes", testQuotesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("QuoteToEpisodeUsingEpisode", testQuoteToOneEpisodeUsingEpisode)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("EpisodeToQuotes", testEpisodeToManyQuotes)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("QuoteToEpisodeUsingQuotes", testQuoteToOneSetOpEpisodeUsingEpisode)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("EpisodeToQuotes", testEpisodeToManyAddOpQuotes)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Episodes", testEpisodesReload)
	t.Run("Quotes", testQuotesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Episodes", testEpisodesReloadAll)
	t.Run("Quotes", testQuotesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Episodes", testEpisodesSelect)
	t.Run("Quotes", testQuotesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Episodes", testEpisodesUpdate)
	t.Run("Quotes", testQuotesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Episodes", testEpisodesSliceUpdateAll)
	t.Run("Quotes", testQuotesSliceUpdateAll)
}
