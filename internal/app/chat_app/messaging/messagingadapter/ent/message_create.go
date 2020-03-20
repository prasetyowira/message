// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/prasetyowira/message/internal/app/chat_app/messaging/messagingadapter/ent/message"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	uid        *string
	text       *string
	created_at *time.Time
	updated_at *time.Time
}

// SetUID sets the uid field.
func (mc *MessageCreate) SetUID(s string) *MessageCreate {
	mc.uid = &s
	return mc
}

// SetText sets the text field.
func (mc *MessageCreate) SetText(s string) *MessageCreate {
	mc.text = &s
	return mc
}

// SetCreatedAt sets the created_at field.
func (mc *MessageCreate) SetCreatedAt(t time.Time) *MessageCreate {
	mc.created_at = &t
	return mc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the updated_at field.
func (mc *MessageCreate) SetUpdatedAt(t time.Time) *MessageCreate {
	mc.updated_at = &t
	return mc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	if mc.uid == nil {
		return nil, errors.New("ent: missing required field \"uid\"")
	}
	if err := message.UIDValidator(*mc.uid); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"uid\": %v", err)
	}
	if mc.text == nil {
		return nil, errors.New("ent: missing required field \"text\"")
	}
	if mc.created_at == nil {
		v := message.DefaultCreatedAt()
		mc.created_at = &v
	}
	if mc.updated_at == nil {
		v := message.DefaultUpdatedAt()
		mc.updated_at = &v
	}
	return mc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	var (
		m     = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: message.FieldID,
			},
		}
	)
	if value := mc.uid; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: message.FieldUID,
		})
		m.UID = *value
	}
	if value := mc.text; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: message.FieldText,
		})
		m.Text = *value
	}
	if value := mc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: message.FieldCreatedAt,
		})
		m.CreatedAt = *value
	}
	if value := mc.updated_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: message.FieldUpdatedAt,
		})
		m.UpdatedAt = *value
	}
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}
