// Package sqlite3 contains the types for schema ''.
package sqlite3

// GENERATED BY XOXO. DO NOT EDIT.

import (
	"errors"
)

// AuthGroupPermission represents a row from 'auth_group_permissions'.
type AuthGroupPermission struct {
	ID           int `json:"id"`            // id
	GroupID      int `json:"group_id"`      // group_id
	PermissionID int `json:"permission_id"` // permission_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthGroupPermission exists in the database.
func (agp *AuthGroupPermission) Exists() bool {
	return agp._exists
}

// Deleted provides information if the AuthGroupPermission has been deleted from the database.
func (agp *AuthGroupPermission) Deleted() bool {
	return agp._deleted
}

// Insert inserts the AuthGroupPermission to the database.
func (agp *AuthGroupPermission) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if agp._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO auth_group_permissions (` +
		`group_id, permission_id` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, agp.GroupID, agp.PermissionID)
	res, err := db.Exec(sqlstr, agp.GroupID, agp.PermissionID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	agp.ID = int(id)
	agp._exists = true

	return nil
}

// Update updates the AuthGroupPermission in the database.
func (agp *AuthGroupPermission) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !agp._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if agp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE auth_group_permissions SET ` +
		`group_id = ?, permission_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, agp.GroupID, agp.PermissionID, agp.ID)
	_, err = db.Exec(sqlstr, agp.GroupID, agp.PermissionID, agp.ID)
	return err
}

// Save saves the AuthGroupPermission to the database.
func (agp *AuthGroupPermission) Save(db XODB) error {
	if agp.Exists() {
		return agp.Update(db)
	}

	return agp.Insert(db)
}

// Delete deletes the AuthGroupPermission from the database.
func (agp *AuthGroupPermission) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !agp._exists {
		return nil
	}

	// if deleted, bail
	if agp._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM auth_group_permissions WHERE id = ?`

	// run query
	XOLog(sqlstr, agp.ID)
	_, err = db.Exec(sqlstr, agp.ID)
	if err != nil {
		return err
	}

	// set deleted
	agp._deleted = true

	return nil
}

// AuthGroup returns the AuthGroup associated with the AuthGroupPermission's GroupID (group_id).
//
// Generated from foreign key 'auth_group_permissions_group_id_fkey'.
func (agp *AuthGroupPermission) AuthGroup(db XODB) (*AuthGroup, error) {
	return AuthGroupByID(db, agp.GroupID)
}

// AuthPermission returns the AuthPermission associated with the AuthGroupPermission's PermissionID (permission_id).
//
// Generated from foreign key 'auth_group_permissions_permission_id_fkey'.
func (agp *AuthGroupPermission) AuthPermission(db XODB) (*AuthPermission, error) {
	return AuthPermissionByID(db, agp.PermissionID)
}

// AuthGroupPermissionsByGroupID retrieves a row from 'auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_0e939a4f'.
func AuthGroupPermissionsByGroupID(db XODB, groupID int) ([]*AuthGroupPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM auth_group_permissions ` +
		`WHERE group_id = ?`

	// run query
	XOLog(sqlstr, groupID)
	q, err := db.Query(sqlstr, groupID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthGroupPermission{}
	for q.Next() {
		agp := AuthGroupPermission{
			_exists: true,
		}

		// scan
		err = q.Scan(&agp.ID, &agp.GroupID, &agp.PermissionID)
		if err != nil {
			return nil, err
		}

		res = append(res, &agp)
	}

	return res, nil
}

// AuthGroupPermissionsByPermissionID retrieves a row from 'auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_8373b171'.
func AuthGroupPermissionsByPermissionID(db XODB, permissionID int) ([]*AuthGroupPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM auth_group_permissions ` +
		`WHERE permission_id = ?`

	// run query
	XOLog(sqlstr, permissionID)
	q, err := db.Query(sqlstr, permissionID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthGroupPermission{}
	for q.Next() {
		agp := AuthGroupPermission{
			_exists: true,
		}

		// scan
		err = q.Scan(&agp.ID, &agp.GroupID, &agp.PermissionID)
		if err != nil {
			return nil, err
		}

		res = append(res, &agp)
	}

	return res, nil
}

// AuthGroupPermissionByGroupIDPermissionID retrieves a row from 'auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_group_id_0cd325b0_uniq'.
func AuthGroupPermissionByGroupIDPermissionID(db XODB, groupID int, permissionID int) (*AuthGroupPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM auth_group_permissions ` +
		`WHERE group_id = ? AND permission_id = ?`

	// run query
	XOLog(sqlstr, groupID, permissionID)
	agp := AuthGroupPermission{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, groupID, permissionID).Scan(&agp.ID, &agp.GroupID, &agp.PermissionID)
	if err != nil {
		return nil, err
	}

	return &agp, nil
}

// AuthGroupPermissionByID retrieves a row from 'auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_id_pkey'.
func AuthGroupPermissionByID(db XODB, id int) (*AuthGroupPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM auth_group_permissions ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	agp := AuthGroupPermission{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&agp.ID, &agp.GroupID, &agp.PermissionID)
	if err != nil {
		return nil, err
	}

	return &agp, nil
}
