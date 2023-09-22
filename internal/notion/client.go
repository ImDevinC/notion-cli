package notion

import (
	"context"

	"github.com/jomei/notionapi"
)

type Client struct {
	client *notionapi.Client
}

type Config struct {
	Token string
}

func NewClient(cfg Config) *Client {
	client := notionapi.NewClient(notionapi.Token(cfg.Token))
	return &Client{
		client: client,
	}
}

func (c *Client) AddNote(ctx context.Context, blockID string, msg string) error {
	req := notionapi.AppendBlockChildrenRequest{
		Children: []notionapi.Block{
			notionapi.ParagraphBlock{
				BasicBlock: notionapi.BasicBlock{
					Type:   "paragraph",
					Object: "block",
				},
				Paragraph: notionapi.Paragraph{
					RichText: []notionapi.RichText{
						{
							Text: &notionapi.Text{Content: msg},
						},
					},
				},
			},
		},
	}

	_, err := c.client.Block.AppendChildren(ctx, notionapi.BlockID(blockID), &req)
	return err
}
