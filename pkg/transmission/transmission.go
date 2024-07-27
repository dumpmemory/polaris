package transmission

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"polaris/log"
	"strings"

	"github.com/hekmon/transmissionrpc/v3"
	"github.com/pkg/errors"
)

func NewClient(c Config) (*Client, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, errors.Wrap(err, "parse url")
	}
	if c.User != "" {
		log.Info("transmission login with user: ", c.User)
		u.User = url.UserPassword(c.User, c.Password)
	}
	u.Path = "/transmission/rpc"

	tbt, err := transmissionrpc.New(u, nil)
	if err != nil {
		return nil, errors.Wrap(err, "connect transmission")
	}
	_, err = tbt.TorrentGetAll(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "transmission cannot connect")
	}
	return &Client{c: tbt, cfg: c}, nil
}

type Config struct {
	URL      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
}
type Client struct {
	c   *transmissionrpc.Client
	cfg Config
}

func (c *Client) GetAll() ([]*Torrent, error) {
	all, err := c.c.TorrentGetAll(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "get all")
	}
	var torrents []*Torrent
	for _, t := range all {
		torrents = append(torrents, &Torrent{
			ID:     *t.ID,
			c:      c.c,
			Config: c.cfg,		
		})
	}
	return torrents, nil
}

func (c *Client) Download(link, dir string) (*Torrent, error) {
	if strings.HasPrefix(link, "http") {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		resp, err:=client.Get(link)
		if err == nil {
			if resp.StatusCode == http.StatusFound {
				loc, err := resp.Location()
				if err == nil {
					link = loc.String()
					log.Warnf("transimision redirect to url: %v", link)
				}
			}
	
		}
	
	} 
	t, err := c.c.TorrentAdd(context.TODO(), transmissionrpc.TorrentAddPayload{
		Filename:    &link,
		DownloadDir: &dir,
	})
	log.Infof("get torrent info: %+v", t)
	if t.ID == nil {
		return nil, fmt.Errorf("download torrent error: %v", link)
	}

	return &Torrent{
		ID:     *t.ID,
		c:      c.c,
		Config: c.cfg,
	}, err
}

type Torrent struct {
	//t *transmissionrpc.Torrent
	c  *transmissionrpc.Client
	ID int64 `json:"id"`
	Config
}

func (t *Torrent) reloadClient() error {
	c, err := NewClient(t.Config)
	if err != nil {
		return err
	}
	t.c = c.c
	return nil
}

func (t *Torrent) getTorrent() transmissionrpc.Torrent {
	r, err := t.c.TorrentGetAllFor(context.TODO(), []int64{t.ID})
	if err != nil {
		log.Errorf("get torrent info for error: %v", err)
	}
	return r[0]
}

func (t *Torrent) Exists() bool {
	r, err := t.c.TorrentGetAllFor(context.TODO(), []int64{t.ID})
	if err != nil {
		log.Errorf("get torrent info for error: %v", err)
	}
	return len(r) > 0
}

func (t *Torrent) Name() string {
	return *t.getTorrent().Name
}

func (t *Torrent) Progress() int {
	if t.getTorrent().IsFinished != nil && *t.getTorrent().IsFinished {
		return 100
	}
	if t.getTorrent().PercentComplete != nil && *t.getTorrent().PercentComplete >= 1 {
		return 100
	}

	if t.getTorrent().PercentComplete != nil {
		p := int(*t.getTorrent().PercentComplete * 100)
		if p == 100 {
			p = 99
		}
		return p
	}
	return 0
}

func (t *Torrent) Stop() error {
	return t.c.TorrentStopIDs(context.TODO(), []int64{t.ID})
}

func (t *Torrent) Start() error {
	return t.c.TorrentStartIDs(context.TODO(), []int64{t.ID})
}

func (t *Torrent) Remove() error {
	return t.c.TorrentRemove(context.TODO(), transmissionrpc.TorrentRemovePayload{
		IDs:             []int64{t.ID},
		DeleteLocalData: true,
	})
}

func (t *Torrent) Size() int {
	return int(t.getTorrent().TotalSize.Byte())
}

func (t *Torrent) Save() string {

	d, _ := json.Marshal(*t)
	return string(d)
}

func ReloadTorrent(s string) (*Torrent, error) {
	var torrent = Torrent{}
	err := json.Unmarshal([]byte(s), &torrent)
	if err != nil {
		return nil, err
	}

	err = torrent.reloadClient()
	if err != nil {
		return nil, errors.Wrap(err, "reload client")
	}
	return &torrent, nil
}
