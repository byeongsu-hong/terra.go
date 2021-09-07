package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/airbloc/logger"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/pkg/errors"
	terraapp "github.com/terra-project/core/app"
)

type Client interface {
	Codec() *codec.Codec
	Request(payload RequestPayload) (*http.Response, error)
	RequestJSON(payload RequestPayload, respBody interface{}) error
}

type client struct {
	codec    *codec.Codec
	endpoint string
	logger   logger.Logger
	*http.Client
}

func New(codec *codec.Codec, endpoint string) Client {
	if codec == nil {
		codec = terraapp.MakeCodec()
	}

	transport := logTransport{
		transport: http.DefaultTransport,
		logger:    logger.New("http/transport"),
	}

	return client{
		codec:    codec,
		endpoint: endpoint,
		logger:   logger.New("http/client"),
		Client:   &http.Client{Transport: transport},
	}
}

func (c client) Codec() *codec.Codec { return c.codec }

func (c client) Request(payload RequestPayload) (*http.Response, error) {
	u, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "parse endpoint")
	}
	u.Path = path.Join(u.Path, payload.Path)

	q := u.Query()
	for k, v := range payload.Query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(
		payload.Context,
		payload.Method,
		u.String(),
		payload.Body,
	)
	if err != nil {
		return nil, errors.Wrap(err, "new request with context")
	}
	return c.Client.Do(req)
}

func (c client) RequestJSON(payload RequestPayload, respBody interface{}) error {
	resp, err := c.Request(payload)
	if err != nil {
		return errors.Wrap(err, "request")
	}
	defer resp.Body.Close()

	if strings.HasPrefix(payload.Path, "/wasm/contracts/") {
		// json
		var buf bytes.Buffer
		resp.Body = ioutil.NopCloser(io.TeeReader(resp.Body, &buf))

		if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
			if rawBody, err := ioutil.ReadAll(&buf); err == nil {
				c.logger.Debug("failed to parse response body. rawBody={}", string(rawBody))
			}
			return errors.Wrap(err, "parse response body with json")
		}
	} else {
		// amino
		rawBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "read raw body")
		}

		if err := c.codec.UnmarshalJSON(rawBody, respBody); err != nil {
			c.logger.Debug("failed to parse response body. rawBody={}", string(rawBody))
			return errors.Wrap(err, "parse response body with codec")
		}
	}
	return nil
}
