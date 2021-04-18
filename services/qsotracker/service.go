package qsotracker

type(
	Service interface {
		RecordContact(ctx context.Context, contact *pb.Contact) error
	}
)
