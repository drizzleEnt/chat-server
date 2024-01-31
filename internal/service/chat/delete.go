package chat

import "context"

func (s *srv) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
