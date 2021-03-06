package barista

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/appadeia/barista/barista-go/commandlib"
	"github.com/appadeia/barista/barista-go/log"
)

type BugzillaInstance struct {
	Name    string
	Matches []string
	Icon    string
	URL     string
	Colour  int
}

var BugzillaInstances []BugzillaInstance = []BugzillaInstance{
	{
		Name:    I18n("Red Hat Bugzilla"),
		Matches: []string{"rh", "rhbz"},
		Icon:    "https://lcp.world/images/trackers/RedHat.png",
		URL:     "https://bugzilla.redhat.com",
		Colour:  0xEE0000,
	},
	{
		Name:    I18n("SUSE Bugzilla"),
		Matches: []string{"bsc", "susebz"},
		Icon:    "https://lcp.world/images/trackers/SUSE.png",
		URL:     "https://bugzilla.suse.com",
		Colour:  0x2D35F,
	},
	{
		Name:    I18n("openSUSE Bugzilla"),
		Matches: []string{"boo"},
		Icon:    "https://lcp.world/images/trackers/openSUSE.png",
		URL:     "https://bugzilla.opensuse.org",
		Colour:  0x73BA25,
	},
	{
		Name:    I18n("Novell Bugzilla"),
		Matches: []string{"bnc"},
		Icon:    "https://lcp.world/images/trackers/Novell.png",
		URL:     "https://bugzilla.novell.com",
		Colour:  0xFF1E1E,
	},
	{
		Name:    I18n("GNOME Bugzilla"),
		Matches: []string{"bgo"},
		Icon:    "https://lcp.world/images/trackers/GNOME.png",
		URL:     "https://bugzilla.gnome.org",
		Colour:  0x4A86CF,
	},
	{
		Name:    I18n("Kernel Bugzilla"),
		Matches: []string{"bko"},
		Icon:    "https://lcp.world/images/trackers/Linux.png",
		URL:     "https://bugzilla.kernel.org",
		Colour:  0xFFD133,
	},
	{
		Name:    I18n("Mozilla Bugzilla"),
		Matches: []string{"bmo"},
		Icon:    "https://lcp.world/images/trackers/Mozilla.png",
		URL:     "https://bugzilla.mozilla.org",
		Colour:  0xFFFFFF,
	},
	{
		Name:    I18n("Samba Bugzilla"),
		Matches: []string{"bso"},
		Icon:    "https://lcp.world/images/trackers/Samba.png",
		URL:     "13172736",
		Colour:  0xC90000,
	},
	{
		Name:    I18n("Xfce Bugzilla"),
		Matches: []string{"bxo"},
		Icon:    "https://lcp.world/images/trackers/Xfce.png",
		URL:     "https://bugzilla.xfce.org",
		Colour:  0x2284F2,
	},
	{
		Name:    I18n("KDE Bugzilla"),
		Matches: []string{"kde"},
		Icon:    "https://lcp.world/images/trackers/KDE.png",
		URL:     "https://bugs.kde.org",
		Colour:  1939955,
	},
	{
		Name:    I18n("Freedesktop Bugzilla"),
		Matches: []string{"fdo"},
		Icon:    "https://lcp.world/images/trackers/Freedesktop.png",
		URL:     "https://bugs.freedesktop.org",
		Colour:  3899566,
	},
	{
		Name:    I18n("GCC Bugzilla"),
		Matches: []string{"gcc"},
		Icon:    "https://lcp.world/images/trackers/GCC.png",
		URL:     "https://gcc.gnu.org/bugzilla/",
		Colour:  16764843,
	},
	{
		Name:    I18n("Mageia Bugzilla"),
		Matches: []string{"mga", "mgabz"},
		Icon:    "https://lcp.world/images/trackers/Mageia.png",
		URL:     "https://bugs.mageia.org/",
		Colour:  2332628,
	},
}

func init() {
	var matches []string
	for _, bz := range BugzillaInstances {
		for _, match := range bz.Matches {
			matches = append(matches, match)
		}
	}
	commandlib.RegisterTag(commandlib.Tag{
		Name:     I18n("Bugzilla"),
		Usage:    I18n("Tag bugs on various Bugzilla instances"),
		Examples: `bnc#1140570, bsc#1140570, boo#1140570, rh#1327846, mga#17400, bko#204371, rh#1327846, mga#17400, bko#204371`,
		Samples: func() []commandlib.TagSample {
			var ret []commandlib.TagSample
			for _, bugzilla := range BugzillaInstances {
				for _, match := range bugzilla.Matches {
					ret = append(ret, commandlib.TagSample{Tag: fmt.Sprintf("%s#1234", match), Desc: bugzilla.Name})
				}
			}
			return ret
		}(),
		ID:     "bz",
		Match:  regexp.MustCompile("(" + strings.Join(matches, "|") + ")"),
		Action: Bugzilla,
	})
}

func Bugzilla(c commandlib.Context) {
	var embeds []commandlib.Embed
	for _, word := range c.Args() {
	InstanceLoop:
		for _, bugzilla := range BugzillaInstances {
			for _, match := range bugzilla.Matches {
				if strings.HasPrefix(word, match+"#") {
					tag := strings.TrimPrefix(word, match+"#")
					bug, err := http.Get(fmt.Sprintf("%s/show_bug.cgi?id=%s&ctype=xml", bugzilla.URL, tag))
					if err != nil {
						log.Error("%+v", err)
						continue InstanceLoop
					}
					body, err := ioutil.ReadAll(bug.Body)
					if err != nil {
						log.Error("%+v", err)
						continue InstanceLoop
					}

					doc, err := xmlquery.Parse(strings.NewReader(string(body)))
					if err != nil {
						log.Error("%+v", err)
						continue InstanceLoop
					}

					bugs, err := xmlquery.QueryAll(doc, "/bugzilla/bug")
					if err != nil {
						log.Error("%+v", err)
						continue InstanceLoop
					}

					if len(bugs) == 0 {
						continue InstanceLoop
					}

					if bugs[0].SelectElement("short_desc") == nil {
						continue InstanceLoop
					}

					embeds = append(embeds, commandlib.Embed{
						Title: commandlib.EmbedHeader{
							Text: bugs[0].SelectElement("short_desc").InnerText(),
							URL:  fmt.Sprintf("%s/show_bug.cgi?id=%s", bugzilla.URL, tag),
						},
						Colour: bugzilla.Colour,
						Footer: commandlib.EmbedHeader{
							Text: fmt.Sprintf("Bug #%s at %s", tag, bugzilla.Name),
							Icon: bugzilla.Icon,
						},
						Header: commandlib.EmbedHeader{
							Text: bugs[0].SelectElement("reporter").SelectAttr("name"),
						},
						Fields: []commandlib.EmbedField{
							{
								Title:  c.I18n("Product"),
								Body:   bugs[0].SelectElement("product").InnerText(),
								Inline: true,
							},
							{
								Title:  c.I18n("Version"),
								Body:   bugs[0].SelectElement("version").InnerText(),
								Inline: true,
							},
							{
								Title:  c.I18n("Component"),
								Body:   bugs[0].SelectElement("component").InnerText(),
								Inline: true,
							},
							{
								Title:  c.I18n("Priority"),
								Body:   bugs[0].SelectElement("priority").InnerText(),
								Inline: true,
							},
							{
								Title:  c.I18n("Severity"),
								Body:   bugs[0].SelectElement("bug_severity").InnerText(),
								Inline: true,
							},
							{
								Title:  c.I18n("Status"),
								Body:   bugs[0].SelectElement("bug_status").InnerText(),
								Inline: true,
							},
						},
					})
				}
			}
		}
	}
	if len(embeds) > 0 {
		c.SendTags("tag-primary", embeds)
	}
}
