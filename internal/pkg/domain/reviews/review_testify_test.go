package reviews

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type reviewSuite struct {
	suite.Suite
	ctx context.Context
}

func (t *reviewSuite) SetupTest() {
	t.ctx = context.Background()
}

func TestReview(t *testing.T) {
	suite.Run(t, new(reviewSuite))
}

func (t *reviewSuite) Test_withCorrectParams() {
	r := NewReview(4, "The iphone X looks good")

	err := r.validate()
	t.Nil(err)
}

func (t *reviewSuite) Test_shouldFailWithWrongNumberOfStars() {
	r := NewReview(8, "The iphone looks REALLY good")

	err := r.validate()
	t.NotNil(err)
	t.Equal(err.Error(), "stars must be between 1 - 5")
}
