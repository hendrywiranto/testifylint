// Code generated by testifylint/internal/cmd/testgen. DO NOT EDIT.

package basic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestErrorInsteadOfErrorIs(t *testing.T) {
	var errSentinel = errors.New("user not found")
	var err error

	t.Run("assert", func(t *testing.T) {
		{
			assert.Error(t, err, errSentinel)                        // want "error-is: invalid usage of assert\\.Error, use assert\\.ErrorIs instead"
			assert.Error(t, err, errSentinel, "msg")                 // want "error-is: invalid usage of assert\\.Error, use assert\\.ErrorIs instead"
			assert.Error(t, err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of assert\\.Error, use assert\\.ErrorIs instead"

			assert.NoError(t, err, errSentinel)                        // want "error-is: invalid usage of assert\\.NoError, use assert\\.NotErrorIs instead"
			assert.NoError(t, err, errSentinel, "msg")                 // want "error-is: invalid usage of assert\\.NoError, use assert\\.NotErrorIs instead"
			assert.NoError(t, err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of assert\\.NoError, use assert\\.NotErrorIs instead"
		}

		// Valid asserts.

		{
			assert.Error(t, err)
			assert.Error(t, err, "msg")
			assert.Error(t, err, "msg with arg %d", 42)
			assert.Errorf(t, err, "msg")
			assert.Errorf(t, err, "msg with arg %d", 42)

			assert.ErrorIs(t, err, errSentinel)
			assert.ErrorIs(t, err, errSentinel, "msg")
			assert.ErrorIs(t, err, errSentinel, "msg with arg %d", 42)
			assert.ErrorIsf(t, err, errSentinel, "msg")
			assert.ErrorIsf(t, err, errSentinel, "msg with arg %d", 42)

			assert.NoError(t, err)
			assert.NoError(t, err, "msg")
			assert.NoError(t, err, "msg with arg %d", 42)
			assert.NoErrorf(t, err, "msg")
			assert.NoErrorf(t, err, "msg with arg %d", 42)

			assert.NotErrorIs(t, err, errSentinel)
			assert.NotErrorIs(t, err, errSentinel, "msg")
			assert.NotErrorIs(t, err, errSentinel, "msg with arg %d", 42)
			assert.NotErrorIsf(t, err, errSentinel, "msg")
			assert.NotErrorIsf(t, err, errSentinel, "msg with arg %d", 42)
		}
	})

	t.Run("assertObj", func(t *testing.T) {
		ass := assert.New(t)

		{
			ass.Error(err, errSentinel)                        // want "error-is: invalid usage of ass\\.Error, use ass\\.ErrorIs instead"
			ass.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of ass\\.Error, use ass\\.ErrorIs instead"
			ass.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of ass\\.Error, use ass\\.ErrorIs instead"

			ass.NoError(err, errSentinel)                        // want "error-is: invalid usage of ass\\.NoError, use ass\\.NotErrorIs instead"
			ass.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of ass\\.NoError, use ass\\.NotErrorIs instead"
			ass.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of ass\\.NoError, use ass\\.NotErrorIs instead"
		}

		// Valid asserts.

		{
			ass.Error(err)
			ass.Error(err, "msg")
			ass.Error(err, "msg with arg %d", 42)
			ass.Errorf(err, "msg")
			ass.Errorf(err, "msg with arg %d", 42)

			ass.ErrorIs(err, errSentinel)
			ass.ErrorIs(err, errSentinel, "msg")
			ass.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			ass.ErrorIsf(err, errSentinel, "msg")
			ass.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			ass.NoError(err)
			ass.NoError(err, "msg")
			ass.NoError(err, "msg with arg %d", 42)
			ass.NoErrorf(err, "msg")
			ass.NoErrorf(err, "msg with arg %d", 42)

			ass.NotErrorIs(err, errSentinel)
			ass.NotErrorIs(err, errSentinel, "msg")
			ass.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			ass.NotErrorIsf(err, errSentinel, "msg")
			ass.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	})

	t.Run("require", func(t *testing.T) {
		{
			require.Error(t, err, errSentinel)                        // want "error-is: invalid usage of require\\.Error, use require\\.ErrorIs instead"
			require.Error(t, err, errSentinel, "msg")                 // want "error-is: invalid usage of require\\.Error, use require\\.ErrorIs instead"
			require.Error(t, err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of require\\.Error, use require\\.ErrorIs instead"

			require.NoError(t, err, errSentinel)                        // want "error-is: invalid usage of require\\.NoError, use require\\.NotErrorIs instead"
			require.NoError(t, err, errSentinel, "msg")                 // want "error-is: invalid usage of require\\.NoError, use require\\.NotErrorIs instead"
			require.NoError(t, err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of require\\.NoError, use require\\.NotErrorIs instead"
		}

		// Valid requires.

		{
			require.Error(t, err)
			require.Error(t, err, "msg")
			require.Error(t, err, "msg with arg %d", 42)
			require.Errorf(t, err, "msg")
			require.Errorf(t, err, "msg with arg %d", 42)

			require.ErrorIs(t, err, errSentinel)
			require.ErrorIs(t, err, errSentinel, "msg")
			require.ErrorIs(t, err, errSentinel, "msg with arg %d", 42)
			require.ErrorIsf(t, err, errSentinel, "msg")
			require.ErrorIsf(t, err, errSentinel, "msg with arg %d", 42)

			require.NoError(t, err)
			require.NoError(t, err, "msg")
			require.NoError(t, err, "msg with arg %d", 42)
			require.NoErrorf(t, err, "msg")
			require.NoErrorf(t, err, "msg with arg %d", 42)

			require.NotErrorIs(t, err, errSentinel)
			require.NotErrorIs(t, err, errSentinel, "msg")
			require.NotErrorIs(t, err, errSentinel, "msg with arg %d", 42)
			require.NotErrorIsf(t, err, errSentinel, "msg")
			require.NotErrorIsf(t, err, errSentinel, "msg with arg %d", 42)
		}
	})

	t.Run("requireObj", func(t *testing.T) {
		r := require.New(t)

		{
			r.Error(err, errSentinel)                        // want "error-is: invalid usage of r\\.Error, use r\\.ErrorIs instead"
			r.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of r\\.Error, use r\\.ErrorIs instead"
			r.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of r\\.Error, use r\\.ErrorIs instead"

			r.NoError(err, errSentinel)                        // want "error-is: invalid usage of r\\.NoError, use r\\.NotErrorIs instead"
			r.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of r\\.NoError, use r\\.NotErrorIs instead"
			r.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of r\\.NoError, use r\\.NotErrorIs instead"
		}

		// Valid requires.

		{
			r.Error(err)
			r.Error(err, "msg")
			r.Error(err, "msg with arg %d", 42)
			r.Errorf(err, "msg")
			r.Errorf(err, "msg with arg %d", 42)

			r.ErrorIs(err, errSentinel)
			r.ErrorIs(err, errSentinel, "msg")
			r.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			r.ErrorIsf(err, errSentinel, "msg")
			r.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			r.NoError(err)
			r.NoError(err, "msg")
			r.NoError(err, "msg with arg %d", 42)
			r.NoErrorf(err, "msg")
			r.NoErrorf(err, "msg with arg %d", 42)

			r.NotErrorIs(err, errSentinel)
			r.NotErrorIs(err, errSentinel, "msg")
			r.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			r.NotErrorIsf(err, errSentinel, "msg")
			r.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	})
}

type ErrorInsteadOfErrorIsSuite struct {
	suite.Suite
}

func (s *ErrorInsteadOfErrorIsSuite) TestAssert() {
	var errSentinel = errors.New("user not found")
	var err error

	{
		{
			s.Error(err, errSentinel)                        // want "error-is: invalid usage of s\\.Error, use s\\.ErrorIs instead"
			s.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Error, use s\\.ErrorIs instead"
			s.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Error, use s\\.ErrorIs instead"

			s.NoError(err, errSentinel)                        // want "error-is: invalid usage of s\\.NoError, use s\\.NotErrorIs instead"
			s.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.NoError, use s\\.NotErrorIs instead"
			s.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.NoError, use s\\.NotErrorIs instead"
		}

		// Valid asserts.

		{
			s.Error(err)
			s.Error(err, "msg")
			s.Error(err, "msg with arg %d", 42)
			s.Errorf(err, "msg")
			s.Errorf(err, "msg with arg %d", 42)

			s.ErrorIs(err, errSentinel)
			s.ErrorIs(err, errSentinel, "msg")
			s.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.ErrorIsf(err, errSentinel, "msg")
			s.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			s.NoError(err)
			s.NoError(err, "msg")
			s.NoError(err, "msg with arg %d", 42)
			s.NoErrorf(err, "msg")
			s.NoErrorf(err, "msg with arg %d", 42)

			s.NotErrorIs(err, errSentinel)
			s.NotErrorIs(err, errSentinel, "msg")
			s.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.NotErrorIsf(err, errSentinel, "msg")
			s.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}

	{
		{
			s.Assert().Error(err, errSentinel)                        // want "error-is: invalid usage of s\\.Assert\\(\\)\\.Error, use s\\.Assert\\(\\)\\.ErrorIs instead"
			s.Assert().Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Assert\\(\\)\\.Error, use s\\.Assert\\(\\)\\.ErrorIs instead"
			s.Assert().Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Assert\\(\\)\\.Error, use s\\.Assert\\(\\)\\.ErrorIs instead"

			s.Assert().NoError(err, errSentinel)                        // want "error-is: invalid usage of s\\.Assert\\(\\)\\.NoError, use s\\.Assert\\(\\)\\.NotErrorIs instead"
			s.Assert().NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Assert\\(\\)\\.NoError, use s\\.Assert\\(\\)\\.NotErrorIs instead"
			s.Assert().NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Assert\\(\\)\\.NoError, use s\\.Assert\\(\\)\\.NotErrorIs instead"
		}

		// Valid asserts.

		{
			s.Assert().Error(err)
			s.Assert().Error(err, "msg")
			s.Assert().Error(err, "msg with arg %d", 42)
			s.Assert().Errorf(err, "msg")
			s.Assert().Errorf(err, "msg with arg %d", 42)

			s.Assert().ErrorIs(err, errSentinel)
			s.Assert().ErrorIs(err, errSentinel, "msg")
			s.Assert().ErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.Assert().ErrorIsf(err, errSentinel, "msg")
			s.Assert().ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			s.Assert().NoError(err)
			s.Assert().NoError(err, "msg")
			s.Assert().NoError(err, "msg with arg %d", 42)
			s.Assert().NoErrorf(err, "msg")
			s.Assert().NoErrorf(err, "msg with arg %d", 42)

			s.Assert().NotErrorIs(err, errSentinel)
			s.Assert().NotErrorIs(err, errSentinel, "msg")
			s.Assert().NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.Assert().NotErrorIsf(err, errSentinel, "msg")
			s.Assert().NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}

	{
		ass := s.Assert()

		{
			ass.Error(err, errSentinel)                        // want "error-is: invalid usage of ass\\.Error, use ass\\.ErrorIs instead"
			ass.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of ass\\.Error, use ass\\.ErrorIs instead"
			ass.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of ass\\.Error, use ass\\.ErrorIs instead"

			ass.NoError(err, errSentinel)                        // want "error-is: invalid usage of ass\\.NoError, use ass\\.NotErrorIs instead"
			ass.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of ass\\.NoError, use ass\\.NotErrorIs instead"
			ass.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of ass\\.NoError, use ass\\.NotErrorIs instead"
		}

		// Valid asserts.

		{
			ass.Error(err)
			ass.Error(err, "msg")
			ass.Error(err, "msg with arg %d", 42)
			ass.Errorf(err, "msg")
			ass.Errorf(err, "msg with arg %d", 42)

			ass.ErrorIs(err, errSentinel)
			ass.ErrorIs(err, errSentinel, "msg")
			ass.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			ass.ErrorIsf(err, errSentinel, "msg")
			ass.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			ass.NoError(err)
			ass.NoError(err, "msg")
			ass.NoError(err, "msg with arg %d", 42)
			ass.NoErrorf(err, "msg")
			ass.NoErrorf(err, "msg with arg %d", 42)

			ass.NotErrorIs(err, errSentinel)
			ass.NotErrorIs(err, errSentinel, "msg")
			ass.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			ass.NotErrorIsf(err, errSentinel, "msg")
			ass.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}
}

func (s *ErrorInsteadOfErrorIsSuite) TestRequire() {
	var errSentinel = errors.New("user not found")
	var err error

	{
		{
			s.Require().Error(err, errSentinel)                        // want "error-is: invalid usage of s\\.Require\\(\\)\\.Error, use s\\.Require\\(\\)\\.ErrorIs instead"
			s.Require().Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Require\\(\\)\\.Error, use s\\.Require\\(\\)\\.ErrorIs instead"
			s.Require().Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Require\\(\\)\\.Error, use s\\.Require\\(\\)\\.ErrorIs instead"

			s.Require().NoError(err, errSentinel)                        // want "error-is: invalid usage of s\\.Require\\(\\)\\.NoError, use s\\.Require\\(\\)\\.NotErrorIs instead"
			s.Require().NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of s\\.Require\\(\\)\\.NoError, use s\\.Require\\(\\)\\.NotErrorIs instead"
			s.Require().NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of s\\.Require\\(\\)\\.NoError, use s\\.Require\\(\\)\\.NotErrorIs instead"
		}

		// Valid requires.

		{
			s.Require().Error(err)
			s.Require().Error(err, "msg")
			s.Require().Error(err, "msg with arg %d", 42)
			s.Require().Errorf(err, "msg")
			s.Require().Errorf(err, "msg with arg %d", 42)

			s.Require().ErrorIs(err, errSentinel)
			s.Require().ErrorIs(err, errSentinel, "msg")
			s.Require().ErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.Require().ErrorIsf(err, errSentinel, "msg")
			s.Require().ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			s.Require().NoError(err)
			s.Require().NoError(err, "msg")
			s.Require().NoError(err, "msg with arg %d", 42)
			s.Require().NoErrorf(err, "msg")
			s.Require().NoErrorf(err, "msg with arg %d", 42)

			s.Require().NotErrorIs(err, errSentinel)
			s.Require().NotErrorIs(err, errSentinel, "msg")
			s.Require().NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			s.Require().NotErrorIsf(err, errSentinel, "msg")
			s.Require().NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}

	{
		req := s.Require()

		{
			req.Error(err, errSentinel)                        // want "error-is: invalid usage of req\\.Error, use req\\.ErrorIs instead"
			req.Error(err, errSentinel, "msg")                 // want "error-is: invalid usage of req\\.Error, use req\\.ErrorIs instead"
			req.Error(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of req\\.Error, use req\\.ErrorIs instead"

			req.NoError(err, errSentinel)                        // want "error-is: invalid usage of req\\.NoError, use req\\.NotErrorIs instead"
			req.NoError(err, errSentinel, "msg")                 // want "error-is: invalid usage of req\\.NoError, use req\\.NotErrorIs instead"
			req.NoError(err, errSentinel, "msg with arg %d", 42) // want "error-is: invalid usage of req\\.NoError, use req\\.NotErrorIs instead"
		}

		// Valid requires.

		{
			req.Error(err)
			req.Error(err, "msg")
			req.Error(err, "msg with arg %d", 42)
			req.Errorf(err, "msg")
			req.Errorf(err, "msg with arg %d", 42)

			req.ErrorIs(err, errSentinel)
			req.ErrorIs(err, errSentinel, "msg")
			req.ErrorIs(err, errSentinel, "msg with arg %d", 42)
			req.ErrorIsf(err, errSentinel, "msg")
			req.ErrorIsf(err, errSentinel, "msg with arg %d", 42)

			req.NoError(err)
			req.NoError(err, "msg")
			req.NoError(err, "msg with arg %d", 42)
			req.NoErrorf(err, "msg")
			req.NoErrorf(err, "msg with arg %d", 42)

			req.NotErrorIs(err, errSentinel)
			req.NotErrorIs(err, errSentinel, "msg")
			req.NotErrorIs(err, errSentinel, "msg with arg %d", 42)
			req.NotErrorIsf(err, errSentinel, "msg")
			req.NotErrorIsf(err, errSentinel, "msg with arg %d", 42)
		}
	}
}
