package driven

import "github.com/Ilja-R/library-auth-service/internal/domain"

type MessagePublisher interface {
	PublishMessage(message domain.Message) error
}
