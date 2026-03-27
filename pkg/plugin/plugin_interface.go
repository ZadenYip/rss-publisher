package plugin

import "zadenyip.com/rss-publisher/pkg/common"

type Plugin interface {
	BuildItems() []common.Item
}
