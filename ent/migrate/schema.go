// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlacklistsColumns holds the columns for the "blacklists" table.
	BlacklistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"media", "torrent"}},
		{Name: "value", Type: field.TypeJSON},
		{Name: "notes", Type: field.TypeString, Nullable: true},
	}
	// BlacklistsTable holds the schema information for the "blacklists" table.
	BlacklistsTable = &schema.Table{
		Name:       "blacklists",
		Columns:    BlacklistsColumns,
		PrimaryKey: []*schema.Column{BlacklistsColumns[0]},
	}
	// DownloadClientsColumns holds the columns for the "download_clients" table.
	DownloadClientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "enable", Type: field.TypeBool},
		{Name: "name", Type: field.TypeString},
		{Name: "implementation", Type: field.TypeEnum, Enums: []string{"transmission", "qbittorrent"}},
		{Name: "url", Type: field.TypeString},
		{Name: "user", Type: field.TypeString, Default: ""},
		{Name: "password", Type: field.TypeString, Default: ""},
		{Name: "settings", Type: field.TypeString, Default: ""},
		{Name: "priority1", Type: field.TypeInt, Default: 1},
		{Name: "remove_completed_downloads", Type: field.TypeBool, Default: true},
		{Name: "remove_failed_downloads", Type: field.TypeBool, Default: true},
		{Name: "tags", Type: field.TypeString, Default: ""},
	}
	// DownloadClientsTable holds the schema information for the "download_clients" table.
	DownloadClientsTable = &schema.Table{
		Name:       "download_clients",
		Columns:    DownloadClientsColumns,
		PrimaryKey: []*schema.Column{DownloadClientsColumns[0]},
	}
	// EpisodesColumns holds the columns for the "episodes" table.
	EpisodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "season_number", Type: field.TypeInt},
		{Name: "episode_number", Type: field.TypeInt},
		{Name: "title", Type: field.TypeString},
		{Name: "overview", Type: field.TypeString},
		{Name: "air_date", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"missing", "downloading", "downloaded"}, Default: "missing"},
		{Name: "monitored", Type: field.TypeBool, Default: false},
		{Name: "target_file", Type: field.TypeString, Nullable: true},
		{Name: "media_id", Type: field.TypeInt, Nullable: true},
	}
	// EpisodesTable holds the schema information for the "episodes" table.
	EpisodesTable = &schema.Table{
		Name:       "episodes",
		Columns:    EpisodesColumns,
		PrimaryKey: []*schema.Column{EpisodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "episodes_media_episodes",
				Columns:    []*schema.Column{EpisodesColumns[9]},
				RefColumns: []*schema.Column{MediaColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HistoriesColumns holds the columns for the "histories" table.
	HistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "media_id", Type: field.TypeInt},
		{Name: "episode_id", Type: field.TypeInt, Nullable: true},
		{Name: "source_title", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "target_dir", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt, Default: 0},
		{Name: "download_client_id", Type: field.TypeInt, Nullable: true},
		{Name: "indexer_id", Type: field.TypeInt, Nullable: true},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"running", "success", "fail", "uploading", "seeding"}},
		{Name: "saved", Type: field.TypeString, Nullable: true},
	}
	// HistoriesTable holds the schema information for the "histories" table.
	HistoriesTable = &schema.Table{
		Name:       "histories",
		Columns:    HistoriesColumns,
		PrimaryKey: []*schema.Column{HistoriesColumns[0]},
	}
	// ImportListsColumns holds the columns for the "import_lists" table.
	ImportListsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"plex", "doulist"}},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "qulity", Type: field.TypeString},
		{Name: "storage_id", Type: field.TypeInt},
		{Name: "settings", Type: field.TypeJSON, Nullable: true},
	}
	// ImportListsTable holds the schema information for the "import_lists" table.
	ImportListsTable = &schema.Table{
		Name:       "import_lists",
		Columns:    ImportListsColumns,
		PrimaryKey: []*schema.Column{ImportListsColumns[0]},
	}
	// IndexersColumns holds the columns for the "indexers" table.
	IndexersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "implementation", Type: field.TypeString},
		{Name: "settings", Type: field.TypeString},
		{Name: "enable_rss", Type: field.TypeBool, Default: true},
		{Name: "priority", Type: field.TypeInt, Default: 50},
		{Name: "seed_ratio", Type: field.TypeFloat32, Nullable: true, Default: 0},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// IndexersTable holds the schema information for the "indexers" table.
	IndexersTable = &schema.Table{
		Name:       "indexers",
		Columns:    IndexersColumns,
		PrimaryKey: []*schema.Column{IndexersColumns[0]},
	}
	// MediaColumns holds the columns for the "media" table.
	MediaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tmdb_id", Type: field.TypeInt},
		{Name: "imdb_id", Type: field.TypeString, Nullable: true},
		{Name: "media_type", Type: field.TypeEnum, Enums: []string{"tv", "movie"}},
		{Name: "name_cn", Type: field.TypeString},
		{Name: "name_en", Type: field.TypeString},
		{Name: "original_name", Type: field.TypeString},
		{Name: "overview", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "air_date", Type: field.TypeString, Default: ""},
		{Name: "resolution", Type: field.TypeEnum, Enums: []string{"720p", "1080p", "2160p", "any"}, Default: "1080p"},
		{Name: "storage_id", Type: field.TypeInt, Nullable: true},
		{Name: "target_dir", Type: field.TypeString, Nullable: true},
		{Name: "download_history_episodes", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "limiter", Type: field.TypeJSON, Nullable: true},
		{Name: "extras", Type: field.TypeJSON, Nullable: true},
	}
	// MediaTable holds the schema information for the "media" table.
	MediaTable = &schema.Table{
		Name:       "media",
		Columns:    MediaColumns,
		PrimaryKey: []*schema.Column{MediaColumns[0]},
	}
	// NotificationClientsColumns holds the columns for the "notification_clients" table.
	NotificationClientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "service", Type: field.TypeString},
		{Name: "settings", Type: field.TypeString},
		{Name: "enabled", Type: field.TypeBool, Default: true},
	}
	// NotificationClientsTable holds the schema information for the "notification_clients" table.
	NotificationClientsTable = &schema.Table{
		Name:       "notification_clients",
		Columns:    NotificationClientsColumns,
		PrimaryKey: []*schema.Column{NotificationClientsColumns[0]},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
	}
	// StoragesColumns holds the columns for the "storages" table.
	StoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "implementation", Type: field.TypeEnum, Enums: []string{"webdav", "local"}},
		{Name: "tv_path", Type: field.TypeString, Nullable: true},
		{Name: "movie_path", Type: field.TypeString, Nullable: true},
		{Name: "settings", Type: field.TypeString, Nullable: true},
		{Name: "deleted", Type: field.TypeBool, Default: false},
		{Name: "default", Type: field.TypeBool, Default: false},
	}
	// StoragesTable holds the schema information for the "storages" table.
	StoragesTable = &schema.Table{
		Name:       "storages",
		Columns:    StoragesColumns,
		PrimaryKey: []*schema.Column{StoragesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlacklistsTable,
		DownloadClientsTable,
		EpisodesTable,
		HistoriesTable,
		ImportListsTable,
		IndexersTable,
		MediaTable,
		NotificationClientsTable,
		SettingsTable,
		StoragesTable,
	}
)

func init() {
	EpisodesTable.ForeignKeys[0].RefTable = MediaTable
}
