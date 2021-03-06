// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/lagzi/EntSymbolBug/ent/picture"
	"github.com/lagzi/EntSymbolBug/ent/predicate"
	"github.com/lagzi/EntSymbolBug/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePicture = "Picture"
	TypeUser    = "User"
)

// PictureMutation represents an operation that mutates the Picture nodes in the graph.
type PictureMutation struct {
	config
	op             Op
	typ            string
	id             *int
	url            *string
	clearedFields  map[string]struct{}
	picture        map[int]struct{}
	removedpicture map[int]struct{}
	clearedpicture bool
	done           bool
	oldValue       func(context.Context) (*Picture, error)
	predicates     []predicate.Picture
}

var _ ent.Mutation = (*PictureMutation)(nil)

// pictureOption allows management of the mutation configuration using functional options.
type pictureOption func(*PictureMutation)

// newPictureMutation creates new mutation for the Picture entity.
func newPictureMutation(c config, op Op, opts ...pictureOption) *PictureMutation {
	m := &PictureMutation{
		config:        c,
		op:            op,
		typ:           TypePicture,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPictureID sets the ID field of the mutation.
func withPictureID(id int) pictureOption {
	return func(m *PictureMutation) {
		var (
			err   error
			once  sync.Once
			value *Picture
		)
		m.oldValue = func(ctx context.Context) (*Picture, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Picture.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPicture sets the old Picture of the mutation.
func withPicture(node *Picture) pictureOption {
	return func(m *PictureMutation) {
		m.oldValue = func(context.Context) (*Picture, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PictureMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PictureMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PictureMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PictureMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Picture.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetURL sets the "url" field.
func (m *PictureMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *PictureMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Picture entity.
// If the Picture object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PictureMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *PictureMutation) ResetURL() {
	m.url = nil
}

// AddPictureIDs adds the "picture" edge to the User entity by ids.
func (m *PictureMutation) AddPictureIDs(ids ...int) {
	if m.picture == nil {
		m.picture = make(map[int]struct{})
	}
	for i := range ids {
		m.picture[ids[i]] = struct{}{}
	}
}

// ClearPicture clears the "picture" edge to the User entity.
func (m *PictureMutation) ClearPicture() {
	m.clearedpicture = true
}

// PictureCleared reports if the "picture" edge to the User entity was cleared.
func (m *PictureMutation) PictureCleared() bool {
	return m.clearedpicture
}

// RemovePictureIDs removes the "picture" edge to the User entity by IDs.
func (m *PictureMutation) RemovePictureIDs(ids ...int) {
	if m.removedpicture == nil {
		m.removedpicture = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.picture, ids[i])
		m.removedpicture[ids[i]] = struct{}{}
	}
}

// RemovedPicture returns the removed IDs of the "picture" edge to the User entity.
func (m *PictureMutation) RemovedPictureIDs() (ids []int) {
	for id := range m.removedpicture {
		ids = append(ids, id)
	}
	return
}

// PictureIDs returns the "picture" edge IDs in the mutation.
func (m *PictureMutation) PictureIDs() (ids []int) {
	for id := range m.picture {
		ids = append(ids, id)
	}
	return
}

// ResetPicture resets all changes to the "picture" edge.
func (m *PictureMutation) ResetPicture() {
	m.picture = nil
	m.clearedpicture = false
	m.removedpicture = nil
}

// Where appends a list predicates to the PictureMutation builder.
func (m *PictureMutation) Where(ps ...predicate.Picture) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PictureMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Picture).
func (m *PictureMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PictureMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.url != nil {
		fields = append(fields, picture.FieldURL)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PictureMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case picture.FieldURL:
		return m.URL()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PictureMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case picture.FieldURL:
		return m.OldURL(ctx)
	}
	return nil, fmt.Errorf("unknown Picture field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PictureMutation) SetField(name string, value ent.Value) error {
	switch name {
	case picture.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	}
	return fmt.Errorf("unknown Picture field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PictureMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PictureMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PictureMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Picture numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PictureMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PictureMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PictureMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Picture nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PictureMutation) ResetField(name string) error {
	switch name {
	case picture.FieldURL:
		m.ResetURL()
		return nil
	}
	return fmt.Errorf("unknown Picture field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PictureMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.picture != nil {
		edges = append(edges, picture.EdgePicture)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PictureMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case picture.EdgePicture:
		ids := make([]ent.Value, 0, len(m.picture))
		for id := range m.picture {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PictureMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedpicture != nil {
		edges = append(edges, picture.EdgePicture)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PictureMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case picture.EdgePicture:
		ids := make([]ent.Value, 0, len(m.removedpicture))
		for id := range m.removedpicture {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PictureMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedpicture {
		edges = append(edges, picture.EdgePicture)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PictureMutation) EdgeCleared(name string) bool {
	switch name {
	case picture.EdgePicture:
		return m.clearedpicture
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PictureMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Picture unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PictureMutation) ResetEdge(name string) error {
	switch name {
	case picture.EdgePicture:
		m.ResetPicture()
		return nil
	}
	return fmt.Errorf("unknown Picture edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                   Op
	typ                  string
	id                   *int
	age                  *int
	addage               *int
	name                 *string
	clearedFields        map[string]struct{}
	usertopicedge        map[int]struct{}
	removedusertopicedge map[int]struct{}
	clearedusertopicedge bool
	done                 bool
	oldValue             func(context.Context) (*User, error)
	predicates           []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddUsertopicedgeIDs adds the "usertopicedge" edge to the Picture entity by ids.
func (m *UserMutation) AddUsertopicedgeIDs(ids ...int) {
	if m.usertopicedge == nil {
		m.usertopicedge = make(map[int]struct{})
	}
	for i := range ids {
		m.usertopicedge[ids[i]] = struct{}{}
	}
}

// ClearUsertopicedge clears the "usertopicedge" edge to the Picture entity.
func (m *UserMutation) ClearUsertopicedge() {
	m.clearedusertopicedge = true
}

// UsertopicedgeCleared reports if the "usertopicedge" edge to the Picture entity was cleared.
func (m *UserMutation) UsertopicedgeCleared() bool {
	return m.clearedusertopicedge
}

// RemoveUsertopicedgeIDs removes the "usertopicedge" edge to the Picture entity by IDs.
func (m *UserMutation) RemoveUsertopicedgeIDs(ids ...int) {
	if m.removedusertopicedge == nil {
		m.removedusertopicedge = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.usertopicedge, ids[i])
		m.removedusertopicedge[ids[i]] = struct{}{}
	}
}

// RemovedUsertopicedge returns the removed IDs of the "usertopicedge" edge to the Picture entity.
func (m *UserMutation) RemovedUsertopicedgeIDs() (ids []int) {
	for id := range m.removedusertopicedge {
		ids = append(ids, id)
	}
	return
}

// UsertopicedgeIDs returns the "usertopicedge" edge IDs in the mutation.
func (m *UserMutation) UsertopicedgeIDs() (ids []int) {
	for id := range m.usertopicedge {
		ids = append(ids, id)
	}
	return
}

// ResetUsertopicedge resets all changes to the "usertopicedge" edge.
func (m *UserMutation) ResetUsertopicedge() {
	m.usertopicedge = nil
	m.clearedusertopicedge = false
	m.removedusertopicedge = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.usertopicedge != nil {
		edges = append(edges, user.EdgeUsertopicedge)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeUsertopicedge:
		ids := make([]ent.Value, 0, len(m.usertopicedge))
		for id := range m.usertopicedge {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedusertopicedge != nil {
		edges = append(edges, user.EdgeUsertopicedge)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeUsertopicedge:
		ids := make([]ent.Value, 0, len(m.removedusertopicedge))
		for id := range m.removedusertopicedge {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedusertopicedge {
		edges = append(edges, user.EdgeUsertopicedge)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeUsertopicedge:
		return m.clearedusertopicedge
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeUsertopicedge:
		m.ResetUsertopicedge()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
