package service

import (
	"context"
	"errors"
	"testing"

	"github.com/m-bromo/atom-ly/internal/mocks"
	repository "github.com/m-bromo/atom-ly/internal/repository/link"
	"github.com/m-bromo/atom-ly/pkg/hasher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ShortenLink(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name      string
		args      args
		setupMock func(repo *mocks.MockLinkRepository, h *mocks.MockHasher)
		want      string
		wantErr   bool
	}{
		{
			name: "should succeed and return a short code (new link)",
			args: args{
				url: "http://example.com",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				repo.EXPECT().GetByUrl(mock.Anything, mock.Anything).Return(0, repository.ErrLinkNotFound)
				repo.EXPECT().Save(mock.Anything, mock.Anything).Return(1, nil)
				h.EXPECT().Encode(1).Return("abcde", nil)
			},
			want:    "abcde",
			wantErr: false,
		},
		{
			name: "should find an existing url and return it's short code",
			args: args{
				url: "http://example.com",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				repo.EXPECT().GetByUrl(mock.Anything, "http://example.com").Return(1, nil)
				h.EXPECT().Encode(1).Return("abcde", nil)
			},
			want:    "abcde",
			wantErr: false,
		},
		{
			name: "should fail when hasher fails to encode existing url",
			args: args{
				url: "http://example.com",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				repo.EXPECT().GetByUrl(mock.Anything, mock.Anything).Return(1, nil)
				h.EXPECT().Encode(1).Return("", errors.New("hasher error"))
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should fail when hasher fails to encode new url",
			args: args{
				url: "http://example.com",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				repo.EXPECT().GetByUrl(mock.Anything, mock.Anything).Return(0, repository.ErrLinkNotFound)
				repo.EXPECT().Save(mock.Anything, mock.Anything).Return(1, nil)
				h.EXPECT().Encode(1).Return("", hasher.ErrInvalidCode)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should fail when repository fails to find existing url",
			args: args{
				url: "http://example.com",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				repo.EXPECT().GetByUrl(mock.Anything, "http://example.com").Return(0, errors.New("db error"))
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should fail when repository fails to save",
			args: args{
				url: "http://example.com",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				repo.EXPECT().GetByUrl(mock.Anything, mock.Anything).Return(0, repository.ErrLinkNotFound)
				repo.EXPECT().Save(mock.Anything, mock.Anything).Return(0, errors.New("db error"))
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLinkRepo := mocks.NewMockLinkRepository(t)
			mockHasher := mocks.NewMockHasher(t)

			if tt.setupMock != nil {
				tt.setupMock(mockLinkRepo, mockHasher)
			}
			service := NewLinkService(mockLinkRepo, mockHasher)

			got, err := service.ShortenLink(context.Background(), tt.args.url)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_Redirect(t *testing.T) {
	type args struct {
		shortCode string
	}

	tests := []struct {
		name      string
		args      args
		setupMock func(repo *mocks.MockLinkRepository, h *mocks.MockHasher)
		want      string
		wantErr   bool
	}{
		{
			name: "should succed and return the original url",
			args: args{
				shortCode: "abcd123",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				h.EXPECT().Decode("abcd123").Return(1, nil)
				repo.EXPECT().GetByID(context.Background(), 1).Return("abcd123", nil)
			},
			want:    "abcd123",
			wantErr: false,
		},
		{
			name: "should fail when repository fails",
			args: args{
				shortCode: "abcd123",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				h.EXPECT().Decode("abcd123").Return(1, nil)
				repo.EXPECT().GetByID(context.Background(), 1).Return("", errors.New("db error"))
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should fail when hasher fails",
			args: args{
				shortCode: "abcd123",
			},
			setupMock: func(repo *mocks.MockLinkRepository, h *mocks.MockHasher) {
				h.EXPECT().Decode("abcd123").Return(0, errors.New("hasher error"))
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLinkRepo := mocks.NewMockLinkRepository(t)
			mockHasher := mocks.NewMockHasher(t)

			if tt.setupMock != nil {
				tt.setupMock(mockLinkRepo, mockHasher)
			}
			service := NewLinkService(mockLinkRepo, mockHasher)

			got, err := service.Redirect(context.Background(), tt.args.shortCode)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
