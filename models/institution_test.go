// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testInstitutions(t *testing.T) {
	t.Parallel()

	query := Institutions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testInstitutionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = institution.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInstitutionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Institutions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInstitutionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := InstitutionSlice{institution}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testInstitutionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := InstitutionExists(tx, institution.ID)
	if err != nil {
		t.Errorf("Unable to check if Institution exists: %s", err)
	}
	if !e {
		t.Errorf("Expected InstitutionExistsG to return true, but got false.")
	}
}
func testInstitutionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	institutionFound, err := FindInstitution(tx, institution.ID)
	if err != nil {
		t.Error(err)
	}

	if institutionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testInstitutionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Institutions(tx).Bind(institution); err != nil {
		t.Error(err)
	}
}

func testInstitutionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Institutions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testInstitutionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institutionOne := &Institution{}
	institutionTwo := &Institution{}
	if err = randomize.Struct(seed, institutionOne, institutionDBTypes, false, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}
	if err = randomize.Struct(seed, institutionTwo, institutionDBTypes, false, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institutionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = institutionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Institutions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testInstitutionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	institutionOne := &Institution{}
	institutionTwo := &Institution{}
	if err = randomize.Struct(seed, institutionOne, institutionDBTypes, false, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}
	if err = randomize.Struct(seed, institutionTwo, institutionDBTypes, false, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institutionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = institutionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func institutionBeforeInsertHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func institutionAfterInsertHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func institutionAfterSelectHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func institutionBeforeUpdateHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func institutionAfterUpdateHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func institutionBeforeDeleteHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func institutionAfterDeleteHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func institutionBeforeUpsertHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func institutionAfterUpsertHook(e boil.Executor, o *Institution) error {
	*o = Institution{}
	return nil
}

func testInstitutionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Institution{}
	o := &Institution{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, institutionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Institution object: %s", err)
	}

	AddInstitutionHook(boil.BeforeInsertHook, institutionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	institutionBeforeInsertHooks = []InstitutionHook{}

	AddInstitutionHook(boil.AfterInsertHook, institutionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	institutionAfterInsertHooks = []InstitutionHook{}

	AddInstitutionHook(boil.AfterSelectHook, institutionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	institutionAfterSelectHooks = []InstitutionHook{}

	AddInstitutionHook(boil.BeforeUpdateHook, institutionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	institutionBeforeUpdateHooks = []InstitutionHook{}

	AddInstitutionHook(boil.AfterUpdateHook, institutionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	institutionAfterUpdateHooks = []InstitutionHook{}

	AddInstitutionHook(boil.BeforeDeleteHook, institutionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	institutionBeforeDeleteHooks = []InstitutionHook{}

	AddInstitutionHook(boil.AfterDeleteHook, institutionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	institutionAfterDeleteHooks = []InstitutionHook{}

	AddInstitutionHook(boil.BeforeUpsertHook, institutionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	institutionBeforeUpsertHooks = []InstitutionHook{}

	AddInstitutionHook(boil.AfterUpsertHook, institutionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	institutionAfterUpsertHooks = []InstitutionHook{}
}
func testInstitutionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInstitutionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx, institutionColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInstitutionToManyInstitutionRoles(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Institution
	var b, c InstitutionRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, institutionRoleDBTypes, false, institutionRoleColumnsWithDefault...)
	randomize.Struct(seed, &c, institutionRoleDBTypes, false, institutionRoleColumnsWithDefault...)

	b.InstitutionID = a.ID
	c.InstitutionID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	institutionRole, err := a.InstitutionRoles(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range institutionRole {
		if v.InstitutionID == b.InstitutionID {
			bFound = true
		}
		if v.InstitutionID == c.InstitutionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := InstitutionSlice{&a}
	if err = a.L.LoadInstitutionRoles(tx, false, (*[]*Institution)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.InstitutionRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.InstitutionRoles = nil
	if err = a.L.LoadInstitutionRoles(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.InstitutionRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", institutionRole)
	}
}

func testInstitutionToManyAddOpInstitutionRoles(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Institution
	var b, c, d, e InstitutionRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, institutionDBTypes, false, strmangle.SetComplement(institutionPrimaryKeyColumns, institutionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*InstitutionRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, institutionRoleDBTypes, false, strmangle.SetComplement(institutionRolePrimaryKeyColumns, institutionRoleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*InstitutionRole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddInstitutionRoles(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.InstitutionID {
			t.Error("foreign key was wrong value", a.ID, first.InstitutionID)
		}
		if a.ID != second.InstitutionID {
			t.Error("foreign key was wrong value", a.ID, second.InstitutionID)
		}

		if first.R.Institution != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Institution != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.InstitutionRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.InstitutionRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.InstitutionRoles(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testInstitutionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = institution.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testInstitutionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := InstitutionSlice{institution}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testInstitutionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Institutions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	institutionDBTypes = map[string]string{`Acronym`: `character varying`, `CreatedAt`: `timestamp without time zone`, `CreatedBy`: `character varying`, `Description`: `character varying`, `FromDate`: `timestamp without time zone`, `ID`: `integer`, `Name`: `character varying`, `ThruDate`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `UpdatedBy`: `character varying`}
	_                  = bytes.MinRead
)

func testInstitutionsUpdate(t *testing.T) {
	t.Parallel()

	if len(institutionColumns) == len(institutionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	if err = institution.Update(tx); err != nil {
		t.Error(err)
	}
}

func testInstitutionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(institutionColumns) == len(institutionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	institution := &Institution{}
	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, institution, institutionDBTypes, true, institutionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(institutionColumns, institutionPrimaryKeyColumns) {
		fields = institutionColumns
	} else {
		fields = strmangle.SetComplement(
			institutionColumns,
			institutionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(institution))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := InstitutionSlice{institution}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testInstitutionsUpsert(t *testing.T) {
	t.Parallel()

	if len(institutionColumns) == len(institutionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	institution := Institution{}
	if err = randomize.Struct(seed, &institution, institutionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = institution.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Institution: %s", err)
	}

	count, err := Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &institution, institutionDBTypes, false, institutionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Institution struct: %s", err)
	}

	if err = institution.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Institution: %s", err)
	}

	count, err = Institutions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}