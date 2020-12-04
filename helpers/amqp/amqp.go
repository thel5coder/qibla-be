package amqp

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

// IQueue ...
type IQueue interface {
	PushQueue(data map[string]interface{}, types string) error
	PushQueueReconnect(url string, data map[string]interface{}, types, deadLetterKey string) (*amqp.Connection, *amqp.Channel, error)
}

const (
	MailExchange = "qibla.mail.exchange"
	//mail incoming
	MailIncoming = "qibla.mail.incoming.queue"
	//mail deadletter
	MailDeadLetter = "qibla.mail.deadletter.queue"

	// SendNotificationExchange ...
	SendNotificationExchange = "qibla.send_notification.exchange"
	// SendNotification ...
	SendNotification = "qibla.send_notification.incoming.queue"
	// SendNotificationDeadLetter ...
	SendNotificationDeadLetter = "qibla.send_notification.deadletter.queue"

	// DisbursementMutationExchange ...
	DisbursementMutationExchange = "disbursement_mutation.exchange"
	// DisbursementMutation ...
	DisbursementMutation = "disbursement_mutation.incoming.queue"
	// DisbursementMutationDeadLetter ...
	DisbursementMutationDeadLetter = "disbursement_mutation.deadletter.queue"

	// DisbursementRequestExchange ...
	DisbursementRequestExchange = "disbursement_request.exchange"
	// DisbursementRequest ...
	DisbursementRequest = "disbursement_request.incoming.queue"
	// DisbursementRequestDeadLetter ...
	DisbursementRequestDeadLetter = "disbursement_request.deadletter.queue"

	// DisbursementCallbackExchange ...
	DisbursementCallbackExchange = "disbursement_callback.exchange"
	// DisbursementCallback ...
	DisbursementCallback = "disbursement_callback.incoming.queue"
	// DisbursementCallbackDeadLetter ...
	DisbursementCallbackDeadLetter = "disbursement_callback.deadletter.queue"
)

// queue ...
type queue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewQueue ...
func NewQueue(conn *amqp.Connection, channel *amqp.Channel) IQueue {
	return &queue{
		Connection: conn,
		Channel:    channel,
	}
}

// PushQueue ...
func (m queue) PushQueue(data map[string]interface{}, types string) error {
	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	return err
}

// PushQueueReconnect ...
func (m queue) PushQueueReconnect(url string, data map[string]interface{}, types, deadLetterKey string) (*amqp.Connection, *amqp.Channel, error) {
	if m.Connection != nil {
		if m.Connection.IsClosed() {
			c := Connection{
				URL: url,
			}
			newConn, newChannel, err := c.Connect()
			if err != nil {
				return nil, nil, err
			}
			m.Connection = newConn
			m.Channel = newChannel
		}
	} else {
		c := Connection{
			URL: url,
		}
		newConn, newChannel, err := c.Connect()
		if err != nil {
			return nil, nil, err
		}
		m.Connection = newConn
		m.Channel = newChannel
	}

	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": deadLetterKey,
	}
	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, args)
	if err != nil {
		return nil, nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, nil, nil
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	return m.Connection, m.Channel, err
}
